package tools

/**
 * @author Threema GmbH
 * @copyright Copyright (c) 2015-2016 Threema GmbH
 */

//use Threema\Core\Exception;
//use Threema\Core\KeyPair;

/**
 * Contains static methods to do various Threema cryptography related tasks.
 * Support libsoidum >= 0.2.0 (Namespaces)
 *
 * @package Threema\Core
 */
type CryptToolSodium struct {
	*CryptTool
}

/**
 * @param string $data
 * @param string $nonce
 * @param string $senderPrivateKey
 * @param string $recipientPublicKey
 * @return string encrypted box
 */
func (self *CryptToolSodium) MakeBox(data string, nonce string, senderPrivateKey string, recipientPublicKey string) string {
	kp := Sodium.Crypto_box_keypair_from_secretkey_and_publickey(senderPrivateKey, recipientPublicKey)

	return Sodium.Crypto_box(data, nonce, kp)
}

/**
 * make a secret box
 *
 * @param $data
 * @param $nonce
 * @param $key
 * @return mixed
 */
func (self *CryptToolSodium) MakeSecretBox(data string, nonce string, key string) {
	return Sodium.Crypto_secretbox(data, nonce, key)
}

/**
 * @param string $box
 * @param string $recipientPrivateKey
 * @param string $senderPublicKey
 * @param string $nonce
 * @return null|string
 */
func (self *CryptToolSodium) OpenBox(box string, recipientPrivateKey string, senderPublicKey string, nonce string) string {
	kp := Sodium.Crypto_box_keypair_from_secretkey_and_publickey(recipientPrivateKey, senderPublicKey)
	return Sodium.Crypto_box_open(box, nonce, kp)
}

/**
 * decrypt a secret box
 *
 * @param string $box as binary
 * @param string $nonce as binary
 * @param string $key as binary
 * @return string as binary
 */
func (self *CryptToolSodium) OpenSecretBox(box string, nonce string, key string) string {
	return Sodium.Crypto_secretbox_open(box, nonce, key)
}

/**
 * Generate a new key pair.
 *
 * @return KeyPair the new key pair
 */
func (self *CryptToolSodium) GenerateKeyPair() map[string]string {
	kp := Sodium.Crypto_box_keypair()
	m := make(map[string]string)
	m[Sodium.Crypto_box_secretkey(kp)] = Sodium.Crypto_box_publickey(kp)
	return m
}

/**
 * @param int $size
 * @return string
 */
func (self *CryptToolSodium) CreateRandom(size int) string {
	return Sodium.randombytes_buf(size)
}

/**
 * Derive the public key
 *
 * @param string $privateKey in binary
 * @return string public key as binary
 */
func (self *CryptToolSodium) DerivePublicKey(privateKey string) string {
	return Sodium.Crypto_box_publickey_from_secretkey(privateKey)
}

/**
 * Check if implementation supported
 * @return bool
 */
func (self *CryptToolSodium) IsSupported() bool {
	// TODO return true == extension_loaded("libsodium") && false == method_exists("Sodium", "sodium_version_string");
	return false
}

/**
 * Validate crypt tool
 *
 * @return bool
 * @throws Exception
 */
func (self *CryptToolSodium) Validate() bool {
	if false == self.IsSupported() {
		panic("Sodium implementation not supported")
	}
	return true
}

/**
 * @return string
 */
func (self *CryptToolSodium) GetName() string {
	return "sodium"
}

/**
 * Description of the CryptTool
 * @return string
 */
func (self *CryptToolSodium) GetDescription() string {
	return "Sodium implementation ".Sodium.Version_string()
}
