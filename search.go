package main

// search represents a request for search. The Query is in `q` and the response
// will be responded on `r`. The response is a slice of interface{} which are
// the IDs of all the entries that match the query. A `limit` can be set as to
// how many entries to be returned
type search struct {
	q     string
	limit int
	r     chan []interface{}
}

var (
	searchChan = make(chan search)
)

// func dumb() {
// 	time.Sleep(1 * time.Second)
//
// 	req := search{
// 		q:     "zzzz",
// 		limit: 3,
// 		r:     make(chan []interface{}),
// 	}
//
// 	searchChan <- req
// 	result := <-req.r
// 	if result != nil {
// 		log.Printf("Request completed: %v", result)
// 	} else {
// 		log.Println("Nothing matched!")
// 	}
// }
