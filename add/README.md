# Add microservice

## Direct usage

* On one terminal page
```shell
go run main.go
```
 Press `Ctrl + C` to exit

 * On a separate terminal page
```shell
curl -XPOST -d '{"a":4, "b":5}' http://localhost:8080/add
```
Where `a` and `b` have any int value

## Usage via docker

* On one terminal page
```shell
docker build -t mu-add . 
# Any other tag will do inside docker build but must be the same used in docker run
docker run --rm -p 4000:8080 mcaci/mu-add
```
* On a separate terminal page
```shell
curl -XPOST -d '{"a":4, "b":5}' http://localhost:8080/add
```
Where `a` and `b` have any int value

For now `Ctrl + C` doesn't work so the `docker run` command must be killed manually until I find a better solution