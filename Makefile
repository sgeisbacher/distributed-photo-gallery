.PHONY: build buildinternal

buildinternal:
	go build -o bin/master.$(GOOS)-$(GOARCH) cmd/master/main.go
	go build -o bin/stats-server.$(GOOS)-$(GOARCH) cmd/statsserver/main.go
	go build -o bin/events-eavesdropper.$(GOOS)-$(GOARCH) cmd/eventseavesdropper/main.go
	go build -o bin/importer.$(GOOS)-$(GOARCH) cmd/importer/main.go
	go build -o bin/apiserver.$(GOOS)-$(GOARCH) cmd/apiserver/main.go

build: 
	env GOOS=linux GOARCH=amd64 make buildinternal

release: 
	env GOOS=linux GOARCH=amd64 make buildinternal
	env GOOS=darwin GOARCH=amd64 make buildinternal