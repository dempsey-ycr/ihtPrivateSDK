all:build
	
build:
	export GOOS=linux
	export GOARCH=amd64
	go build /Users/xiaodongxu/Workspace/go/src/ihtPrivateSDK/iht/ipfs/ipfs.go -output ./bin/iht_ipfs
	# go build -osarch="linux/amd64" -output ./bin/iht_ipfs
clean:
	@rm -rf bin
	 
test:
	go test ./go/... -race

out:
	rm ipfs -rf
	mkdir -p ipfs/conf.d
	cp -a conf.d/ipfs_dev.xml ipfs/conf.d/
	cp -a bin/iht_ipfs ipfs/
	tar -zcf ipfs.tar.gz ipfs/
