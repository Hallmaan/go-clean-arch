package nats_repository_impl

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func NewMultipleNatsClient() (*nats.Conn, error) {
	// natsClients := make([]*nats.Conn, 0)
	// for i := 0; i < 3; i++ {
	// 	natsConn, err := nats.Connect("nats://localhost:4222")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	natsClients = append(natsClients, natsConn)
	// }

	// return natsClients, nil
	natsConn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, err
	}

	return natsConn, nil
}

func NewNatsJetstreamClient() (nats.JetStreamContext, error) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	// if err := createStream(js); err != nil {
	// 	return nil, err
	// }

	return js, nil
}

func NewNatsKeyValue(js nats.JetStreamContext) (nats.KeyValue, error) {
	kv, err := keyValueBacket(js, "bucket-uk", time.Duration(10)*time.Second, 1)
	if err != nil {
		return nil, err
	}

	return kv, nil
}

func keyValueBacket(client nats.JetStreamContext, bucket string, ttl time.Duration, history int) (nats.KeyValue, error) {
	kv, err := client.KeyValue(bucket)
	if err != nil {
		kv, err = client.CreateKeyValue(
			&nats.KeyValueConfig{
				Bucket:  bucket,
				TTL:     ttl,
				History: uint8(history),
			},
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return kv, nil
}
