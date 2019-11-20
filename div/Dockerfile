FROM golang

RUN ["git", "config", "--global", "url.\"git@github.com:\".insteadOf", "\"https://github.com/\""]
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN ["go", "get", "-u", "github.com/mcaci/mu-calc/div"]

ENTRYPOINT ["go", "run", "github.com/mcaci/mu-calc/div"]

EXPOSE 8080
