FROM golang

RUN ["git", "config", "--global", "url.\"git@github.com:\".insteadOf", "\"https://github.com/\""]
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN ["go", "get", "-u", "github.com/mcaci/fibonacci-service"]

WORKDIR /go/src/github.com/mcaci/fibonacci-service/

#EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]

