package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/eliezedeck/fuzzysearch/fuzzy"
	"github.com/rwynn/gtm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	_searchIDBase   = make([]interface{}, 0, 128)
	_searchDataBase = make([]string, 0, 128)
)

func getBeginTimestamp(session *mgo.Session, options *gtm.Options) bson.MongoTimestamp {
	return 0
}

func monitor(session *mgo.Session) {
	// Parse the "fields" parameter
	paramFields = strings.Replace(paramFields, ",", " ", -1)
	fields := strings.Fields(paramFields)

	ops, errs := gtm.Tail(session, &gtm.Options{
		After:               getBeginTimestamp, // if nil defaults to LastOpTimestamp
		Filter:              filterOps,         // filter
		OpLogDatabaseName:   nil,               // if nil defaults to "local"
		OpLogCollectionName: nil,               // if nil a defaults to a collection prefixed "oplog."
		CursorTimeout:       nil,               // if nil defaults to 100s
		ChannelSize:         32,                // if less than 1 defaults to 20
	})

	ticker := time.NewTicker(time.Second * 10)

	for {
		select {
		case err := <-errs:
			fmt.Println(err)

		case op := <-ops:
			if op.IsInsert() {
				line := ""
				for _, field := range fields {
					if data, ok := op.Data[field]; ok {
						line += fmt.Sprintf(" %s", data)
					}
				}

				if len(line) > 0 {
					_searchDataBase = append(_searchDataBase, line[1:])
					_searchIDBase = append(_searchIDBase, op.Id)
					log.Printf("Added: %s/ %s", op.Id, line[1:])
				}
				continue
			}

			// TODO: Implement the rest of the ops

		case <-ticker.C:
			log.Println("Number of searchable entries:", len(_searchIDBase))

		case r := <-searchChan:
			var result []interface{}
			matches := fuzzy.FindFoldIdx(r.q, _searchDataBase, r.limit)
			for _, idx := range matches {
				result = append(result, _searchIDBase[idx])
			}
			r.r <- result
		}
	}
}
