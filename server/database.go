package server

import (
	"redis"
	"log"
	"encoding/json"
)

var client int

func getclient() (redis.Client) {
	spec := redis.DefaultSpec().Db(1)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Fatal(e)
	}
	return client
}

func StoreEvent(event map[string]) {
	client := getclient()
	client.Set('test', json.Marshall(event));
	client.Expire('test', 0); 
}
