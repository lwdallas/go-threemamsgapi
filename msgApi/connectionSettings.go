package gothreemamsgapi

const TlsOptionForceHttps = "TlsOptionForceHttps"
const TlsOptionVersion = "TlsOptionVersion"
const TlsOptionCipher = "TlsOptionCipher"

type ConnectionSettings struct {
	threemaId string

	secret string

	host string

	tlsOptions map[string]int
}

/**
 * @param string $threemaId valid threema id (8chars)
 * @param string $secret secret
 * @param string|null $host server url
 * @param array|null $tlsOptions advanced TLS options
 */
func NewConnectionSettings(threemaId string, secret string, host string, tlsOptions map[string]int) *ConnectionSettings {
	if "" == host {
		host = "https://msgapi.threema.ch"
	}

	// TLS options
	if len(tlsOptions) > 0 {
		tlsOptions[TlsOptionForceHttps], _ = tlsOptions[TlsOptionForceHttps]
		tlsOptions[TlsOptionVersion], _ = tlsOptions[TlsOptionVersion]
		tlsOptions[TlsOptionCipher], _ = tlsOptions[TlsOptionCipher]
	}
	return &ConnectionSettings{threemaId: threemaId, secret: secret, host: host, tlsOptions: tlsOptions}
}

func (self *ConnectionSettings) GetThreemaId() string {
	return self.threemaId
}

func (self *ConnectionSettings) GetSecret() string {
	return self.secret
}

func (self *ConnectionSettings) GetHost() string {
	return self.host
}

func (self *ConnectionSettings) GetTlsOptions() map[string]int {
	return self.tlsOptions
}

func (self *ConnectionSettings) GetTlsOption(option string, defaultValue int) int {
	if _, ok := self.tlsOptions[option]; ok {
		return self.tlsOptions[option]
	} else {
		return defaultValue
	}
}
