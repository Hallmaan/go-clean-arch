package nats_ports

type RepositoryNats interface {
	KVKeys(bucket string) ([]string, error)
	KVGet(bucket, key string) ([]byte, error)
	KVPut(bucket, key string, payload any) error
}
