package gothreemamsgapi

//namespace Threema\MsgApi;

//use Threema\Core\Exception;
//use Threema\Core\Url;
//use Threema\MsgApi\Commands\Capability;
//use Threema\MsgApi\Commands\CommandInterface;
//use Threema\MsgApi\Commands\Credits;
//use Threema\MsgApi\Commands\DownloadFile;
//use Threema\MsgApi\Commands\FetchPublicKey;
//use Threema\MsgApi\Commands\LookupEmail;
//use Threema\MsgApi\Commands\LookupPhone;
//use Threema\MsgApi\Commands\MultiPartCommandInterface;
//use Threema\MsgApi\Commands\Results\CapabilityResult;
//use Threema\MsgApi\Commands\Results\DownloadFileResult;
//use Threema\MsgApi\Commands\Results\FetchPublicKeyResult;
//use Threema\MsgApi\Commands\Results\LookupIdResult;
//use Threema\MsgApi\Commands\Results\Result;
//use Threema\MsgApi\Commands\Results\SendSimpleResult;
//use Threema\MsgApi\Commands\Results\SendE2EResult;
//use Threema\MsgApi\Commands\Results\UploadFileResult;
//use Threema\MsgApi\Commands\SendSimple;
//use Threema\MsgApi\Commands\SendE2E;
//use Threema\MsgApi\Commands\UploadFile;

/**
 * Class Connection
 */
type Connection struct {
	Setting        string

/**
 * @var PublicKeyStore
 */
	PublicKeyStore string

}

(self *Connection) func Init(setting ConnectionSettings, publicKeyStore PublicKeyStore) {
	self.setting = setting
	self.publicKeyStore = publicKeyStore
}

/**
 * @param Receiver $receiver
 * @param $text
 * @return SendSimpleResult
 */
(self *Connection) func SendSimple(receiver Receiver, text string) SendSimpleResult {
	command := PerformSendSimple(receiver, text)
	return self.Post(command)
}

/**
 * @param string $threemaId
 * @param string $nonce
 * @param string $box
 * @return SendE2EResult
 */
(self *Connection) func SendE2E(threemaId, nonce, box string) SendE2EResult {
	command := PerformSendE2E(threemaId, nonce, box)
	return self.Post(command)
}

/**
 * @param $encryptedFileData (binary string)
 * @return UploadFileResult
 */
(self *Connection) func UploadFile(encryptedFileData []byte) UploadFileResult {
	command := PerformUploadFile(encryptedFileData)
	return self.PostMultiPart(command)
}


/**
 * @param $blobId
 * @param callable $progress
 * @return DownloadFileResult
 */
(self *Connection) func DownloadFile(blobId string, progress interface{}) DownloadFileResult { // TODO: this is a coding error, make a progress interface
	command := PerformDownloadFile(blobId)
	return self.Get(command, progress)
}

/**
 * @param $phoneNumber
 * @return LookupIdResult
 */
(self *Connection) func KeyLookupByPhoneNumber(phoneNumber string) LookupIdResult {
	command := PerformLookupPhone(phoneNumber)
	return self.Get(command)
}

/**
 * @param string $email
 * @return LookupIdResult
 */
(self *Connection) func KeyLookupByEmail(email string) LookupIdResult {
	command := PerformLookupEmail(email)
	return self.Get(command)
}

/**
 * @param string $threemaId valid threema id (8 Chars)
 * @return CapabilityResult
 */
(self *Connection) func KeyCapability(threemaId string) CapabilityResult {
	return self.Get(PerformCapability(threemaId))
}


/**
 * @return CreditsResult
 */
(self *Connection) func Credits() CreditsResult {
	return self.Get(PerformCredits())
}

/**
 * @param $threemaId
 * @return FetchPublicKeyResult
 */
(self *Connection) func FetchPublicKey(threemaId) FetchPublicKeyResult {
	publicKey := nil

	if (nil != = self.PublicKeyStore) {
	publicKey = self.PublicKeyStore.getPublicKey(threemaId)
	}

	if (nil == = publicKey) {
	command := FetchPublicKey(threemaId)
	result := self.Get(command)
	if (false === result.isSuccess()) {
	return result
	}
	publicKey = result.getRawResponse()

	if (nil != = self.publicKeyStore) {
	self.publicKeyStore.setPublicKey(threemaId, publicKey)
	}
	}

	//create a key result
	return FetchPublicKeyResult(200, publicKey)
}

/**
 * @param callable $progress
 * @return array
 */
(self *Connection) func CreateDefaultOptions(progress interface{}) {
	options := make(map[string]string)
	options["CURLOPT_RETURNTRANSFER"] = true
}

//no progress
if (nil != = progress) {
options = ["CURLOPT_NOPROGRESS"] = false
options = ["CURLOPT_PROGRESSFUNCTION"] = progress
}

//tls settings

if (true == = self.setting.GetTlsOption(ConnectionSettings.TlsOptionForceHttps, false)) {
//limit allowed protocols to HTTPS
options["CURLOPT_PROTOCOLS"] = "CURLPROTO_HTTPS"
}
if (tlsVersion = self.setting.GetTlsOption(ConnectionSettings.TlsOptionVersion)) {
if (is_int(tlsVersion)) {
//if number is given use it
options["CURLOPT_SSLVERSION"] = tlsVersion
} else {
//interpret strings as TLS versions
switch (tlsVersion) {
case "1.0":
options["CURLOPT_SSLVERSION"] = "CURL_SSLVERSION_TLSv1_0"
break
case "1.1":
options["CURLOPT_SSLVERSION"] = "CURL_SSLVERSION_TLSv1_1"
break
case "1.2":
options["CURLOPT_SSLVERSION"] = "CURL_SSLVERSION_TLSv1_2"
break
default:
options["CURLOPT_SSLVERSION"] = "CURL_SSLVERSION_DEFAULT"
break
}
}
}
if (tlsCipher = self.setting.GetTlsOption(ConnectionSettings.TlsOptionCipher, nil)) {
if (true == = is_string(tlsCipher)) {
options["CURLOPT_SSL_CIPHER_LIST"] = tlsCipher
}
}
return options
}

/**
 * @param array $params
 * @return array
 */
(self *Connection) func processRequestParams(params array) {
	if (nil == params) {
		params = array() //TODO
	}

	params["from"] = self.setting.getThreemaId()
	params["secret"] = self.setting.getSecret()

	return params
}

/**
 * @param CommandInterface $command
 * @param callable $progress
 * @return Result
 */
(self *Connection) func Get(command CommandInterface, progress intefrace {}) Result {
params := self.ProcessRequestParams(command.GetParams())
return self.Call(command.GetPath(),
self.CreateDefaultOptions(progress),
params,
func (httpCode, response) use /* TODO use */
(command) {
return command.ParseResult(httpCode, response)
})
}

/**
 * @param CommandInterface $command
 * @return Result
 */
(self *Connection) func Post(command CommandInterface) Result {
	options := self.CreateDefaultOptions()
	params := self.processRequestParams(command.GetParams())

	options["CURLOPT_POST"] = true
	options["CURLOPT_POSTFIELDS"] = http_build_query(params)
	options["CURLOPT_HTTPHEADER"] = array(
		"Content-Type: application/x-www-form-urlencoded")

	return self.Call(command.GetPath(), options, nil, func(httpCode, response) use /* TODO use*/
	(command) {
		return command.ParseResult(httpCode, response)
	})
}

/**
 * @param MultiPartCommandInterface $command
 * @return Result
 */
(self *Connection) funct PostMultiPart( command MultiPartCommandInterface ) {
options := self.createDefaultOptions()
params := self.processRequestParams(command.GetParams())

options["CURLOPT_POST"] = true
options["CURLOPT_HTTPHEADER"] = array("Content-Type: multipart/form-data")
options["CURLOPT_SAFE_UPLOAD"] = true
options["CURLOPT_POSTFIELDS"] = array(
"blob" = command.GetData()
)

return self.Call(command.GetPath(), options, params, func (httpCode, response) use /* TODO */
(command) {
return command.ParseResult(httpCode, response)
})
}

/**
 * @param string $path
 * @param array $curlOptions
 * @param array $parameters
 * @param callable $result
 * @return mixed
 * @throws \Threema\Core\Exception
 */
(self *Connection) func Call(path, curlOptions array, parameters array, result interface{}) {
	fullPath := Url("", self.setting.GetHost())
	fullPath.AddPath(path)

	if (nil != parameters && count(parameters)) {
		for key, value := range parameters {
			fullPath.setValue(key, value)
		}
	}
	session := curl_init(fullPath.GetFullPath())
	curl_setopt_array(session, curlOptions)

	response := curl_exec(session)
	if (false == response) {
		/* TODO throw new Exception($path . ' ' . curl_error($session))*/
	}

	httpCode := curl_getinfo(session, "CURLINFO_HTTP_CODE")
	if (nil == result && httpCode != 200) {
		/* TODO throw new Exception(httpCode) */
	}

	if (nil != result) {
		return result.Invoke(httpCode, response)
	} else {
		return response
	}
}
}
