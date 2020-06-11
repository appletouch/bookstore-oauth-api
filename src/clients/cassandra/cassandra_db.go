package cassandra

import (
	"github.com/appletouch/bookstore-oauth-api/src/utils/profiling"
	"github.com/gocql/gocql"
	"time"
)

var (
	session *gocql.Session
)

func init() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	/*
		docker/gocql performance problems. Basically when it does the host lookup docker is returning a 172..... ip range
		which the gocql client in the application layer cannot see. As a result you get 3 network timeouts.
		DisableInitialHostLookup=true prevents this for docker. (result of setting: 9,9sec ---> 22ms)
	*/
	cluster.DisableInitialHostLookup = true

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

}

func GetSession() *gocql.Session {

	defer profiling.TimeTrack(time.Now(), "GetSession")
	return session

}
