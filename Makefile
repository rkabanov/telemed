clean:
	rm -f build/*

createdb:
	docker exec -it pg15 createdb --username=root --owner=root telemeddb

tables-create:
	cat ./migrate/telemeddb.sql | docker exec -i pg15 psql -U root telemeddb

tables-truncate:
	docker exec -i pg15 psql -U root telemeddb -c 'truncate table doctors; truncate table patients;'

psql:
	docker exec -i pg15 psql telemeddb root

build:
	rm -f build/*
	CGO_ENABLED=0 go build -ldflags '-X "main.buildDate=${shell date +%Y%m%d.%H%M%S}"' -o build/telemed

run:
	go run .

run-pg:
	rm -f build/*
	CGO_ENABLED=0 go build -ldflags '-X "main.buildDate=${shell date +%Y%m%d.%H%M%S}"' -o build/telemed
	build/telemed -store=postgres -pghost=localhost -pgport=5433 -pguser=root -pgpass=secret -pgdb=telemeddb \
		-webpath=telemed -webport=8133

upload:
	tar zcvf build/telemed-build.tgz -C build -- telemed
	scp build/telemed-build.tgz deploy1@192.3.120.10:/tmp

deploy:
	tar zcvf build/telemed-build.tgz -C build -- telemed
	scp build/telemed-build.tgz deploy1@192.3.120.10:/tmp
	ssh deploy1@192.3.120.10 -t "/home/deploy1/services/telemed/update-service.sh"

test:
	go test -v --cover ./...

.PHONY: clean createdb tables-create tables-truncate build run-mem run-pg deploy test
