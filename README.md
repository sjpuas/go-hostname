For Compile use:

  docker run -t -v `pwd`:/go golang:1.7.3 /bin/bash -c 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo github.com/sjpuas/go-hostname'
