FROM golang

RUN ["git", "config", "--global", "url.\"git@github.com:\".insteadOf", "\"https://github.com/\""]
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN ["go", "get", "-u", "github.com/mcaci/mu-calc/sub"]

ENTRYPOINT ["go", "run", "github.com/mcaci/mu-calc/sub"]

EXPOSE 8080
