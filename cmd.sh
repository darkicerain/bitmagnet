#!/bin/sh
go build -o main bitmagnet-io/bitmagnet \
&& ./main worker run --keys=dht_crawler --keys=queue_server --keys=http_server
