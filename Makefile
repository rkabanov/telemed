clean:
	rm -f build/*

build:
	rm -f build/*
	go build -o build/service

deploy:
	tar zcvf build/service-build.tgz build/service
	scp build/service-build.tgz deploy1@192.3.120.10:/tmp


test:
	go test -v --cover ./...

.PHONY: clean build deploy test


