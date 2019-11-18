# Fibonacci microservice

## Direct usage

* On one terminal page
```shell
go run main.go
```
 Press `Ctrl + C` to exit

 * On a separate terminal page
```shell
curl -XPOST -d '{"n":4}' http://localhost:8080/fibonacci
```
Where `n` has any int value

## Usage via docker

* On one terminal page
```shell
docker build -t mcaci/fibonacci-service . 
# Any other tag will do inside docker build but must be the same used in docker run
docker run --rm -p 4000:8080 mcaci/fibonacci-service
```
* On a separate terminal page
```shell
curl -XPOST -d '{"n":4}' http://localhost:4000/fibonacci
```
Where `n` has any int value

For now `Ctrl + C` doesn't work so the `docker run` command must be killed manually until I find a better solution