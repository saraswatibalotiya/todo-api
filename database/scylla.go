package database

import (
	"os"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
	cluster := gocql.NewCluster(os.Getenv("SCYLLA_HOST"))
	cluster.Keyspace = os.Getenv("SCYLLA_KEYSPACE")
	cluster.Consistency = gocql.Quorum

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

func Close() {
	Session.Close()
}
