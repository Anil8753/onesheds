dev:
	mkdir -p ../dbg
	cp ./service.env ../dbg
	cd httpd && go build -o ../../dbg/service
	cd ../dbg && ./service

ware:
	mkdir -p ../bin
	rm -rf ../bin/ware
	mkdir -p ../bin/ware
	cp ./cfg/ware/service.env ../bin/ware/service.env
	cd httpd && go build -o ../../bin/ware

dep:
	mkdir -p ../bin
	rm -rf ../bin/dep
	mkdir -p ../bin/dep
	cp ./cfg/dep/service.env ../bin/dep/service.env
	cd httpd && go build -o ../../bin/dep
