package main

import (
	"encoding/json"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// search represents a request for search. The Query is in `q` and the response
// will be responded on `r`. The response is a slice of interface{} which are
// the IDs of all the entries that match the query. A `limit` can be set as to
// how many entries to be returned
type search struct {
	q     string
	limit int
	r     chan []interface{}
}

// _search used to unmarshal from JSON
type _search struct {
	Q     string `json:"q"`
	Limit int    `json:"limit"`
}

var (
	searchChan = make(chan search)
)

func serveSearchRequests() {
	var err error

	responder, _ := zmq.NewSocket(zmq.REP)
	responder.Bind(paramSearchListen)
	time.Sleep(500 * time.Millisecond)
	defer responder.Close()

	log.Println("Accepting Search requests...")

	// Grand Search loop
	for {
		b := mustRecv(responder)

		var r _search
		err = json.Unmarshal(b, &r)
		if err != nil {
			log.Printf("JSON Unmarshal error: %s. Ignoring!", err)

			// Respond with an error
			mustSend(responder, "{\"error\":\"Bad request\"}")
			continue
		}

		if len(r.Q) > 0 {
			// Create a proper request
			req := search{
				q:     r.Q,
				limit: r.Limit,
				r:     make(chan []interface{}),
			}

			// Query the search
			searchChan <- req
			result := <-req.r

			// Convert the result to JSON
			b, err = json.Marshal(result)
			if err != nil {
				// Respond with error
				mustSend(responder, "{\"error\":\"Marshaling failure\"}")
				continue
			}

			// Send back the result
			mustSendBytes(responder, b)
			continue
		}

		// Respond empty
		mustSend(responder, "{\"result\":[]}")
	}
}

func mustRecv(socket *zmq.Socket) []byte {
	b, err := socket.RecvBytes(0)
	if err != nil {
		log.Fatalln("ZMQ RecvBytes failure:", err)
	}
	return b
}

func mustSend(socket *zmq.Socket, data string) {
	_, err := socket.Send(data, 0)
	if err != nil {
		log.Fatalln("Failed ZMQ send:", err)
	}
}

func mustSendBytes(socket *zmq.Socket, data []byte) {
	_, err := socket.SendBytes(data, 0)
	if err != nil {
		log.Fatalln("Failed ZMQ send:", err)
	}
}
