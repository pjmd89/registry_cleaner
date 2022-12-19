package registry

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/pjmd89/goutils/jsonutils"
)

func ConfigInit(path, user, password, host, port, protocol string) (r *Config, err error) {
	r = &Config{}
	user = strings.Trim(user, " ")
	password = strings.Trim(user, " ")
	host = strings.Trim(host, " ")
	port = strings.Trim(port, " ")
	protocol = strings.Trim(protocol, " ")
	if _, err = os.Stat(path); err != nil {
		log.Printf("el archivo %s no existe.", path)
	} else {
		jsonutils.GetJson(path, r)
	}
	if user != "" {
		r.User = user
	}
	if password != "" {
		r.Password = password
	}
	if host != "" {
		r.Host = host
	}
	if port != "" {
		r.Port = port
	}
	if protocol != "" {
		r.Protocol = protocol
	}
	rValue := reflect.ValueOf(r)

	for i := 0; i < rValue.Elem().NumField(); i++ {
		value := rValue.Elem().Field(i).Interface().(string)
		newValue := reflect.ValueOf(strings.Trim(value, " "))
		rValue.Elem().Field(i).Set(newValue)
	}
	if r.Port == "" {
		r.Port = "5000"
	}
	if r.Host == "" {
		r = nil
		err = fmt.Errorf("El host no esta configurado")
		return
	}
	if r.Protocol == "" {
		r = nil
		err = fmt.Errorf("El protocolo no esta configurado")
		return
	}
	protocolValue := ProtocolEnum(r.Protocol)
	if !protocolValue.Validate() {
		r = nil
		err = fmt.Errorf("El protocolo no es valido")
		return
	}
	return
}
func (o *Config) GetBaseURL() (r string) {
	r = ""
	credentials := o.User + ":" + o.Password + "@"
	url := o.Host + ":" + o.Port
	if o.Password == "" {
		credentials = o.User + "@"
	}
	if o.User == "" {
		credentials = ""
	}
	if o.Port == "" {
		url = o.Host
	}
	if o.Protocol != "" && url != "" {
		r = o.Protocol + "://" + credentials + url + "/"
	}

	return
}
