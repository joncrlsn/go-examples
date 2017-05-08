package main

//
// Persistent queue
//

import (
	"github.com/beeker1121/goque"
	"log"
)

const (
	dirName = "pqueue"
)

type MyObject struct {
	Name string
	Age  int64
}

func main() {
	add(MyObject{"Jon", 30})
	add(MyObject{"Deb", 31})
	add(MyObject{"Fluffy", 8})
	add(MyObject{"Fido", 2})
	getAndLog()
	getAndLog()
	getAndLog()
	getAndLog()
}

// add adds an object to the given prefix
func add(obj MyObject) {
	pq, err := goque.OpenQueue(dirName)
	if err != nil {
		log.Panicln("Error opening queue", err)
	}
	defer pq.Close()

	item, err := pq.EnqueueObject(obj)
	log.Println("item added: ", obj, item.ID)
}

// getAndLog gets an object from the queue and logs it
func getAndLog() MyObject {
	pq, err := goque.OpenQueue(dirName)
	if err != nil {
		log.Panicln("Error opening queue", err)
	}
	defer pq.Close()

	item, err := pq.Peek()
	if err != nil {
		log.Panicln("Error peeking in queue", err)
	}

	item, err = pq.Dequeue()
	if err != nil {
		log.Panicln("Error peeking in queue", err)
	}

	var obj MyObject
	err = item.ToObject(&obj)
	log.Println("got item id:", item.ID, "Key with prefix:", item.Key, "Object:", obj)
	return obj
}
