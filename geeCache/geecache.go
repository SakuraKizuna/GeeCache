package geeCache

// A Getter loads data for a key
type Getter interface {
	Get(Key string) ([]byte, error)
}

// A GetterFunc implements(工具) Getter with a function
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
