package gothreemamsgapi

import "strings"

/**
 * @author Threema GmbH
 * @copyright Copyright (c) 2015-2016 Threema GmbH
 */

const TYPE_ID = "to"
const TYPE_PHONE = "phone"
const TYPE_EMAIL = "email"

type Receiver struct {

	/**
	 * @var string
	 */
	r_type string

	/**
	 * @var string
	 */
	value string
}

/**
 * @param string $value
 * @param string $type
 */
func NewReceiver(value string, r_type string) *Receiver {
	if nil == r_type {
		r_type = TYPE_ID
	}
	if "" == r_type {
		r_type = TYPE_ID
	}
	return &Receiver{value, r_type}
}

/**
 * @param string $threemaId
 * @return $this
 */
func (self *Receiver) SetToThreemaId(threemaId string) *Receiver {
	return self.setValue(threemaId, TYPE_ID)
}

/**
 * @param string $phoneNo
 * @return $this
 */
func (self *Receiver) SetToPhoneNo(phoneNo string) *Receiver {
	return self.setValue(phoneNo, TYPE_PHONE)
}

/**
 * @param string $emailAddress
 * @return $this
 */
func (self *Receiver) SetToEmail(emailAddress string) *Receiver {
	return self.setValue(emailAddress, TYPE_EMAIL)
}

/**
 * @param string $value
 * @param string $type
 * @return $this
 */
func (self *Receiver) setValue(value string, r_type string) *Receiver {
	self.value = value
	self.r_type = r_type
	return *self
}

/**
 * @return array
 * @throws \InvalidArgumentException
 */
func (self *Receiver) GetParams() map[string]string {
	var to string
	switch self.r_type {
	case TYPE_ID:
		to = self.r_type
		self.value = strings.ToUpper(strings.TrimSpace(self.value))
		break
	case TYPE_EMAIL:
	case TYPE_PHONE:
		to = self.r_type
		break
	default:
		//throw new \InvalidArgumentException();
	}

	m := make(map[string]string)
	m[to] = self.value

	return m
}
