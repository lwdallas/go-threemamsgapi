package tools

/**
 * @author Threema GmbH
 * @copyright Copyright (c) 2015-2016 Threema GmbH
 */

/**
 * Result of a Data Encryption
 *
 * @package Threema\MsgApi\Tool
 */
type EncryptResult struct {

	/**
	 * @var string as binary
	 */
	data string

	/**
	 * @var string as binary
	 */
	key string

	/**
	 * @var string as binary
	 */
	nonce string

	/**
	 * @var int
	 */
	size int
}

/**
 * @param string $data (binary)
 * @param string $key (binary)
 * @param string $nonce (binary)
 * @param int $size
 */
func NewEncryptResult(data string, key string, nonce string, size int) *EncryptResult {
	return &EncryptResult{data, key, nonce, size}
}

/**
 * @return int
 */
func (self *EncryptResult) GetSize() int {
	return self.size
}

/**
 * @return string (binary)
 */
func (self *EncryptResult) GetKey() string {
	return self.key
}

/**
 * @return string (binary)
 */
func (self *EncryptResult) GetNonce() string {
	return self.nonce
}

/**
 * @return string (binary)
 */
func (self *EncryptResult) GetData() string {
	return self.data
}
