package storage

type Storage interface {
	Set(key string, value *Campaign) error
	Get(key string) (string, error)
}
