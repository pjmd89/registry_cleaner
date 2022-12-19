package main

import (
	"flag"

	"github.com/pjmd89/goutils/systemutils/debugmode"
	"github.com/pjmd89/registry_cleaner/lib/registry"
)

var (
	user     string = ""
	password        = ""
	host            = ""
	port            = ""
	config          = "/etc/registry_cleaner/config.json"
	delete          = ""
	protocol        = ""
	get      bool   = false
)
var (
	rc    *registry.Config
	rcErr error
)

func init() {
	flag.StringVar(&user, "user", "", "Nombre de usuario del registry")
	flag.StringVar(&password, "password", "", "password de usuario del registry")
	flag.StringVar(&host, "host", "", "url de conexion del registry")
	flag.StringVar(&port, "port", "", "puerto del registry")
	flag.StringVar(&config, "config", config, "Archivo de configuracion. Por defecto /etc/registry_cleaner/config.json")
	flag.StringVar(&delete, "delete", "", "Archivo json que contiene los registry a eliminar.")
	flag.StringVar(&protocol, "protocol", "", "protocolo del registry (http, https).")
	flag.BoolVar(&get, "get", get, "Nombre de usuario del registry")
	flag.Parse()

	if debugmode.Enabled {
		config = "etc/config.json"
	}

	rc, rcErr = registry.ConfigInit(config, user, password, host, port, protocol)
}
