all:build
	
build:
	export GOOS=linux
	export GOARCH=amd64
	go build /Users/xiaodongxu/Workspace/go/src/ihtPrivateSDK/iht/ipfs/fabric.go -output ./bin/iht_fabric
	# go build -osarch="linux/amd64" -output ./bin/iht_ipfs
clean:
	@rm -rf bin
	 
test:
	go test ./go/... -race

out:
	rm fabric -rf
	mkdir -p fabric/conf.d
	cp -a conf.d/fabric_dev.xml fabric/conf.d/
	cp -a bin/fabric_ipfs fabric/
	tar -zcf fabric.tar.gz fabric/
