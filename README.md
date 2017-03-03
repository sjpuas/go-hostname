[![Codefresh build status]( https://g.codefresh.io/api/badges/build?repoOwner=sjpuas&repoName=go-hostname&branch=master&pipelineName=go-hostname&accountName=sjpuas&type=cf-1)]( https://g.codefresh.io/repositories/sjpuas/go-hostname/builds?filter=trigger:build;branch:master;service:589a209421858e01002ff80d~go-hostname)

For Compile use:

  docker run -t -v $(pwd):/go/src/github.com/sjpuas/go-hostname golang:1.7.3 /bin/bash -c 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo github.com/sjpuas/go-hostname'
