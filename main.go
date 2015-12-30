package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	mongoaddr string

	paramFields string
)

func init() {
	flag.StringVar(&mongoaddr, "mongo", "localhost:3001", "MongoDB address:port, default to localhost:3001 which is where Meteor sets-it up")
	flag.StringVar(&paramFields, "fields", "", "Fields to search")

	flag.Parse()

	if paramFields == "" {
		fmt.Println("You must specify the 'fields' parameter. Ex: --fields=firstName,lastName")
		os.Exit(-1)
	}
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
