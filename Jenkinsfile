node {
  stage("Building Go Binaries") {
    sh "docker run -t -v $(pwd):/go/src/github.com/sjpuas/go-hostname golang:1.7.3 /bin/bash -c 'CGO_ENABLED=0 GOOS=linux go build -a -o /go/src/github.com/sjpuas/go-hostname/go-hostname -installsuffix cgo github.com/sjpuas/go-hostname'"
  }

  stage("Building Docker Image") {
    sh "docker build -t sjpuas/go-hostname ."
  }
}
