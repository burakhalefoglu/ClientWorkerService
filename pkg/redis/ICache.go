package cache

type ICache interface {
	Get(key string)(map[string]string, error)
	Add(key string, value map[string]interface{}) (success bool, err error)
	Delete(key string, fields ...string) (success bool, err error)
	DeleteAll(key string) (success bool, err error)
}
