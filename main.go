package main

import (
	"flag"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	mongoaddr string
)

func init() {
	flag.StringVar(&mongoaddr, "mongo", "localhost:3001", "MongoDB address:port, default to localhost:3001 which is where Meteor sets-it up")
}

func main() {
	msession, err := mgo.Dial(mongoaddr)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB '%s'", mongoaddr)
	}
	defer msession.Close()
	log.Printf("Connected to MongoDB '%s' ...", mongoaddr)

	msession.SetMode(mgo.Monotonic, true)
	monitor(msession)
}
