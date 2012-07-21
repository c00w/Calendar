package server

import (
	"redis"
	"log"
)

var client int

func init() {
	spec := redis.DefaultSpec().Db(1)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Fatal(e)
	}
	client.Set("hi", []byte("Hi"))
}
