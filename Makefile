clean:
	rm -f build/*

build:
	rm -f build/*
	CGO_ENABLED=0 go build -ldflags '-X "main.buildDate=${shell date +%Y%m%d.%H%M%S}"' -o build/service

deploy:
	tar zcvf build/service-build.tgz -C build -- service
	scp build/service-build.tgz deploy1@192.3.120.10:/tmp
	ssh deploy1@192.3.120.10 -t "/home/deploy1/services/telemed/update-service.sh"

test:
	go test -v --cover ./...

.PHONY: clean build deploy test


