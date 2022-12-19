package registry

type Http struct {
}

type Config struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

type ProtocolEnum string

var (
	PROTOCOL_ENUM_HTTP  ProtocolEnum = "http"
	PROTOCOL_ENUM_HTTPS ProtocolEnum = "https"
)

func (o ProtocolEnum) Validate() (r bool) {
	switch o {
	case PROTOCOL_ENUM_HTTP, PROTOCOL_ENUM_HTTPS:
		r = true
	}
	return
}
