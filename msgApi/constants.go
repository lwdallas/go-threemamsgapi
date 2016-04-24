package gothreemamsgapi

const PUBLIC_KEY_PREFIX = "public:"
const PRIVATE_KEY_PREFIX = "private:"

type Constants struct {
	PUBLIC_KEY_PREFIX  string //= "public:";
	PRIVATE_KEY_PREFIX string //= "private:";
}

/**
 * create instance disabled
 */
func NewConstants() *Constants {
	return &Constants{PUBLIC_KEY_PREFIX, PRIVATE_KEY_PREFIX}
}
