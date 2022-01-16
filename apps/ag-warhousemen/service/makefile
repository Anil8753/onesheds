dev:
	mkdir -p ../dbg
	cp ./service.env ../dbg
	cd httpd && go build -o ../../dbg/service
	cd ../dbg && ./service

build:
	rm -rf ../bin
	mkdir ../bin
	cp ./service.env ../bin
	cd httpd && go build -o ../../bin/service

	
	