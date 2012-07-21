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
	expire_d := expire/1000 - time.Now().Unix()
	if expire_d <= 0 {
		return
	}
	expire_d += 48 * 60 * 60
	if e != nil {
		log.Fatal(e)
	}
	hash := md5.New()
	io.WriteString(hash, string(str_event))
	event_key := string(hash.Sum(nil))

	client.Set(event_key, str_event) 
	client.Expire(event_key, expire) 

	date_key := time.Unix(expire/1000, 0).Format("01/02/2006")

	val, e := client.Get(event_key)
	log.Print(string(event_key))
	log.Print(string(val))
	
	client.Rpush(date_key, []byte(event_key))
	client.Expire(date_key, expire)

	dlen, e := client.Llen(date_key)
	vald, e := client.Lrange(date_key, 0, dlen) 
	log.Print(date_key)
	log.Print(len(vald))
}

func getEvents(event_time int64) ([]string) {
	
	client := getclient()
	event_time = event_time/1000
	date_key := time.Unix(event_time, 0).Format("01/02/2006") 
	dlen, e := client.Llen(date_key)
	vald, e := client.Lrange(date_key, 0, dlen)
	if e != nil {
		log.Fatal(e)
	}	
	events := make([]string, len(vald))
	for i, key:= range vald {
		event, e := client.Get(string(key))
		if e != nil {
			log.Fatal(e)
		}
		events[i] = string(event)
	}	
	return events
}
