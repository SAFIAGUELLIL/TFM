BINARY_NAME=benchmark

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} benchmark.go
	# GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux benchmark.go
	# GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows benchmark.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
	# rm ${BINARY_NAME}-darwin
	# rm ${BINARY_NAME}-linux
	# rm ${BINARY_NAME}-windows
