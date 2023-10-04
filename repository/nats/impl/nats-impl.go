package nats_repository_impl

import (
	nats_ports "clean_arch_ws/repository/nats/ports"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type repositoryNats struct {
	clientJS nats.JetStreamContext
	kvBucket map[string]nats.KeyValue
	clients  *nats.Conn
}

func NewRepositoryNats(clientJS nats.JetStreamContext, clients *nats.Conn) nats_ports.RepositoryNats {

	natsKvUsersKey, err := NewNatsKeyValue(clientJS)
	if err != nil {
	}

	kvBucket := make(map[string]nats.KeyValue)

	kvBucket["Transaction"] = natsKvUsersKey

	return &repositoryNats{
		clientJS: clientJS,
		kvBucket: kvBucket,
		clients:  clients,
	}
}

func (r *repositoryNats) KVKeys(bucket string) ([]string, error) {
	return r.kvBucket[bucket].Keys()
}

func (r *repositoryNats) KVGet(bucket, key string) ([]byte, error) {
	result, err := r.kvBucket[bucket].Get(key)
	if err != nil {
		return nil, err
	}

	return result.Value(), nil
}

func (r *repositoryNats) KVHistory(bucket, key string) ([]nats.KeyValueEntry, error) {
	results, err := r.kvBucket[bucket].History(key)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (r *repositoryNats) KVPut(bucket, key string, payload any) error {
	// marshaledPayload, err := proto.Marshal(payload)
	// if err != nil {
	// 	return err
	// }

	marshaledPayload, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = r.kvBucket[bucket].Put(key, marshaledPayload)
	if err != nil {
		fmt.Println(err, "ini error")
		return err
	}

	return nil
}
