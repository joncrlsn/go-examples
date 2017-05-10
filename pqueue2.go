package main

//
// Persistent queue for storing an implementatino of an interface
//

import (
	"github.com/beeker1121/goque"
	"log"
	"time"
)

const (
	dirName = "pqueue"
)

type AnObject interface {
	Sleep()
}

type MyObject struct {
	Name string
	Age  int64
}

func (mo *MyObject) Sleep() {
	time.Sleep(time.Duration(mo.Age))
}

func main() {
	add(&MyObject{"Jon", 30})
	add(&MyObject{"Deb", 31})
	add(&MyObject{"Fluffy", 8})
	add(&MyObject{"Fido", 2})

	log.Println("got Object:", get())
	log.Println("got Object:", get())
	log.Println("got Object:", get())
	log.Println("got Object:", get())
}

// add adds an object to the given prefix
func add(obj AnObject) {
	pq, err := goque.OpenQueue(dirName)
	if err != nil {
		log.Panicln("Error opening queue", err)
	}
	defer pq.Close()

	item, err := pq.EnqueueObject(obj)
	log.Println("item added: ", obj, item.ID)
}

// getAndLog gets an object from the queue and logs it
func get() AnObject {
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
	return &obj
}
