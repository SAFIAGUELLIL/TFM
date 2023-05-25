BINARY_NAME=benchmark

build:
	GOARCH=amd64 GOOS=darwin go build -o ./compiled/${BINARY_NAME} benchmark.go

run: build
	multitime ./compiled/$(BINARY_NAME) 2>&1 | tail -n +3 > ./compiled/go.time 


clean:
	go clean
	rm ./compiled/${BINARY_NAME}
	rm ./compiled/go.time
	
