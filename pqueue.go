package main

//
// Persistent queue
//

import (
	"github.com/beeker1121/goque"
	"log"
)

const (
	dirName = "prefix-queue"
)

type MyObject struct {
	Name string
	Age  int64
}

func main() {
	add("friends", MyObject{"Jon", 30})
	add("friends", MyObject{"Deb", 31})
	getAndLog("friends")
	getAndLog("friends")
}

// add adds an object to the given prefix
func add(queueName string, obj MyObject) {
	pq, err := goque.OpenPrefixQueue(dirName)
	if err != nil {
		log.Panicln("Error opening queue", err)
	}
	defer pq.Close()

	item, err := pq.EnqueueObject([]byte(queueName), obj)
	log.Println("item added: ", obj, item.ID)
}

// getAndLog gets an object from the queue and logs it
func getAndLog(queueName string) MyObject {
	pq, err := goque.OpenPrefixQueue(dirName)
	if err != nil {
		log.Panicln("Error opening queue", err)
	}
	defer pq.Close()
	item, err := pq.Dequeue([]byte(queueName))

	var obj MyObject
	err = item.ToObject(&obj)
	log.Println("got item id:", item.ID, "Key with prefix:", item.Key, "Object:", obj)
	return obj
}
