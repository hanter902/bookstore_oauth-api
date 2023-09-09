package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("cassandra connection successfully created")
	// defer session.Close()

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
