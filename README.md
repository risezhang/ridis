# ridis
This is a cli tool for redis writing in go. It helps to copy / backup / restore data in redis.

It is just started... 

## TODO

1. copy any key from one redis to another
    * set
    * sorted set - done
    * hash
    * string
2. dump any key to local file
3. restore local file to key in redis

## Usage

Copy key from one redis to another. In this example it shows how to copy the key in one redis instance with another key.

```
ridis --from-addr localhost:6379 -from-db=2 -to-addr=localhost:6379 -to-db=2 -from-key=a-z-set -to-key=b-z-set
```

## Build

```
go build -o ridis cmd/main.go
```