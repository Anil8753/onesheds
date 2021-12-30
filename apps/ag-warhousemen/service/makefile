
dep:
	cp ./external/identity/fabric-ca-client ./bin
	cp ./external/identity/config.json ./bin
	cp ./service.env ./bin
	
dev:
	cd httpd && go build -o ../bin/service
	cd bin && ./service

build:
	cp ./external/identity/fabric-ca-client ../bin
	cp ./external/identity/config.json ../bin
	cp ./service.env ../bin
	cd httpd && go build -o ../../bin/service

	
	