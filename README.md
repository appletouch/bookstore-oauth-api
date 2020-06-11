#Bookstore-oauth-api
Project make a oauth api in Go

##Design
This is a Domain Driven Design






##Frameworks

- gin gonic
    Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter.

`go get -u github.com/gin-gonic/gin`

- gocql
    Package gocql implements a fast and robust Cassandra client for the Go programming language.
    
`go get github.com/gocql/gocql`

##Databases (Cassandra)

**Run cassandra start script "createCluster.sh"** ( NEED ENOUGH MEMMORY TO RUN THREE NODES=>2.2GB /node)

- Enter node 1:
`docker exec -it cassandra-node-1 bash`

- check nodes:
`nodetool status`

remark: starting and stoping is done with script to give the cluster some time to start.

**Cassandra ports: **
9042 CQL Native Transport Port
- 7199 JMX monitoring port
- 1024 - 65355 Random port required by JMX. Starting with Java 7u4 a specific port can be specified using the com.sun.management.jmxremote.rmi.port property.
- 7000 Inter-node cluster
- 7001 SSL inter-node cluster
- 9042 CQL Native Transport Port
- 9160 Thrift

remark: to prevent performance problems do NOT allow filtering to run queries on cassandra
allwsys interact with cassandra through primary keys!

##MANUAL INSTALL
- Docker network 
`docker create myNetwork`

- Install:
`docker run --name Oauth-cluster --network myNetwork -d cassandra:3.11.6`

- Connect to cluster
`run -it --network myNetwork --rm cassandra cqlsh Oauth-cluster`

- Create Oauth in CQL shell
`CREATE KEYSPACE IF NOT EXISTS Oauth WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'datacenter1' : 3 };`

    `DESCRIBE keyspaces;`
    
   ` USE Oauth;` 

- Create table

    `CREATE TABLE access_tokens(access_token VARCHAR PRIMARY KEY, user_id BIGINT, client_id BIGINT,expires BIGINT);
`- DML COMMANDS

   ` SELECT * FROM access_tokens;`
   
    `INSERT INTO tokens(access_token, user_id, client_id, expires) VALUES ('xxx', 123, 10, 1293383838);`

- Truncate
`Truncate Oauth.tokens`

- Drop
`Drop Table Oauth.tokens`
