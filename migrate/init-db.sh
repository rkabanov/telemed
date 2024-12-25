#!/bin/sh

# initialise DB
docker exec -it pg15 createdb --username=root --owner=root telemeddb
psql -h localhost -p 1234 -U root -d telemeddb -f telemeddb.sql
