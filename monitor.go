package main

import (
	"fmt"

	"github.com/rwynn/gtm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getBeginTimestamp(session *mgo.Session, options *gtm.Options) bson.MongoTimestamp {
	return 0
}

func monitor(session *mgo.Session) {
	ops, errs := gtm.Tail(session, &gtm.Options{
		After:               getBeginTimestamp, // if nil defaults to LastOpTimestamp
		Filter:              filterOps,         // filter
		OpLogDatabaseName:   nil,               // if nil defaults to "local"
		OpLogCollectionName: nil,               // if nil a defaults to a collection prefixed "oplog."
		CursorTimeout:       nil,               // if nil defaults to 100s
		ChannelSize:         32,                // if less than 1 defaults to 20
	})

	for {
		// loop forever receiving events
		select {
		case err := <-errs:
			// handle errors
			fmt.Println(err)
		case op := <-ops:
			// op will be an insert, delete or update to mongo
			// you can check which by calling op.IsInsert(), op.IsDelete(), or op.IsUpdate()
			// op.Data will get you the full document for inserts and updates
			msg := fmt.Sprintf(`Got op <%v> for object <%v>
          in database <%v>
          and collection <%v>
          and data <%v>
          and timestamp <%v>`,
				op.Operation, op.Id, op.GetDatabase(),
				op.GetCollection(), op.Data, op.Timestamp)
			fmt.Println(msg) // or do something more interesting
		}
	}
}
