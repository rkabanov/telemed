clean:
	rm -f build/*

createdb:
	docker exec -it pg15 createdb --username=root --owner=root servicedb

tables-create:
	cat ./migrate/servicedb.sql | docker exec -i pg15 psql -U root servicedb

tables-truncate:
	docker exec -i pg15 psql -U root servicedb -c 'truncate table doctors; truncate table patients;'

psql:
	docker exec -i pg15 psql servicedb root

build:
	rm -f build/*
	CGO_ENABLED=0 go build -ldflags '-X "main.buildDate=${shell date +%Y%m%d.%H%M%S}"' -o build/service

run:
	go run .

run-pg:
	rm -f build/*
	CGO_ENABLED=0 go build -ldflags '-X "main.buildDate=${shell date +%Y%m%d.%H%M%S}"' -o build/service
	build/service -store=postgres -pghost=localhost -pgport=5433 -pguser=root -pgpass=secret -pgdb=servicedb \
		-webpath=telemed -webport=8133

upload:
	tar zcvf build/service-build.tgz -C build -- service
	scp build/service-build.tgz deploy1@192.3.120.10:/tmp

deploy:
	tar zcvf build/service-build.tgz -C build -- service
	scp build/service-build.tgz deploy1@192.3.120.10:/tmp
	ssh deploy1@192.3.120.10 -t "/home/deploy1/services/telemed/update-service.sh"

test:
	go test -v --cover ./...

.PHONY: clean createdb tables-create tables-truncate build run-mem run-pg deploy test
