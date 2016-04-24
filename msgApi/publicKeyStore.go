package gothreemamsgapi

/**
 * @author Threema GmbH
 * @copyright Copyright (c) 2015-2016 Threema GmbH
 */

/**
 * Interface PublicKeyStore
 * Store the fetched Public Keys
 *
 * @package Threema\MsgApi
 */
type PublicKeyStore struct {

	/**
	 * threemaId => publicKey cache
	 * @var array
	 */
	cache map[string]string
}

/**
 * return null if the public key not found in the store
 * @param string $threemaId
 * @return string|null
 */
func (self *PublicKeyStore) GetPublicKey(threemaId string) {
	if self.cache[threemaId] != "" {
		return self.cache[threemaId]
	}

	publicKey := self.FindPublicKey(threemaId)
	if nil != publicKey {
		self.cache[threemaId] = publicKey
	}
	return publicKey
}

/**
 * return null if the public key not found in the store
 * @param string $threemaId
 * @return string|null
 */
func (self *PublicKeyStore) FindPublicKey(threemaId string) string {
	return nil
}

/**
 * set and save a public key
 * @param string $threemaId
 * @param string $publicKey
 * @return bool
 */
func (self *PublicKeyStore) SetPublicKey(threemaId string, publicKey string) bool {
	self.cache[threemaId] = publicKey
	return self.SavePublicKey(threemaId, publicKey)
}

/**
 * save a public key
 * @param string $threemaId
 * @param string $publicKey
 * @return bool
 */
func (self *PublicKeyStore) SavePublicKey(threemaId string, publicKey string) bool {
	return false
}
