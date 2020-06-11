docker run --name cassandra-node-3 -d \
-e CASSANDRA_CLUSTER_NAME="Oauth-cluster" \
-e CASSANDRA_CLUSTER_TOKENS="8" \
-e CASSANDRA_CLUSTER_DC="dc1" \
-e CASSANDRA_RACK="rack3" \
-e CASSANDRA_ENPOINT_SNITCH="GossipingPropertyFileSnitch" \
-e CASSANDRA_SEEDS="$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' cassandra-node-1)" \
-e  CASSANDRA_BROADCAST_ADDRESS=127.0.0.1 \
-v /Users/pkoolwijk/go/src/github.com/appletouch/bookstore-oauth-api/cassandraDB/data/node-3:/var/lib/cassandra/data \
cassandra:3.11.6