package server

import (
	"redis"
	"log"
	"encoding/json"
	"time"
	"strconv"
	"crypto/md5"
	"io"
)

func getclient() (redis.Client) {
	spec := redis.DefaultSpec().Db(1)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Fatal(e)
	}
	return client
}

func StoreEvent(event map[string]string) {
	client := getclient()
	str_event, e := json.Marshal(event)
	expire, e := strconv.ParseInt(event["date"], 10, 64)
	expire = expire/1000 - time.Now().Unix()
	if expire <= 0 {
		return
	}
	expire += 48 * 60 * 60
	if e != nil {
		log.Fatal(e)
	}
	hash := md5.New()
	io.WriteString(hash, string(str_event))
	event_key := string(hash.Sum(nil))
	client.Set(event_key, str_event) 
	client.Expire(event_key, expire) 
	val, e := client.Get(event_key)
	log.Print(string(event_key))
	log.Print(string(val))
}
