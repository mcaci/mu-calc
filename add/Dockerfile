FROM golang

RUN ["git", "config", "--global", "url.\"git@github.com:\".insteadOf", "\"https://github.com/\""]
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN ["go", "get", "-u", "github.com/mcaci/calc/add"]

ENTRYPOINT ["go", "run", "github.com/mcaci/calc/add"]

EXPOSE 8080
