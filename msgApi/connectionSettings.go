package gothreemamsgapi

const TlsOptionForceHttps = "forceHttps"
const TlsOptionVersion = "tlsVersion"
const TlsOptionCipher = "tlsCipher"

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
func (self *ConnectionSettings) NewConnectionSettings(threemaId string, secret string, host string, tlsOptions map[string]int) {
	self.threemaId = threemaId
	self.secret = secret
	if "" == host {
		host = "https://msgapi.threema.ch"
	}
	self.host = host

	// TLS options
	if len(tlsOptions) > 0 {
		self.tlsOptions[TlsOptionForceHttps], _ = tlsOptions[TlsOptionForceHttps]
		self.tlsOptions[TlsOptionVersion], _ = tlsOptions[TlsOptionVersion]
		self.tlsOptions[TlsOptionCipher], _ = tlsOptions[TlsOptionCipher]
	}
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

func (self *ConnectionSettings) GetTlsOptions() string {
	return self.tlsOptions
}

func (self *ConnectionSettings) GetTlsOption(option, defaultValue string) string {
	if _, ok := self.tlsOptions[option]; ok {
		return self.tlsOptions[option]
	} else {
		return defaultValue
	}
}
