[![Go Report Card](https://goreportcard.com/badge/github.com/arxdsilva/hc)](https://goreportcard.com/report/github.com/arxdsilva/hc)


smaller path: straight line
smaller ammout of ds: (abs(smaller between x and y) devided by 3) + 1 (if there's remainder)
for a 10x10 grid, It'd take 4 drones

install/usage:
```shell 
$ go get github.com/arxdsilva/hc
$ cd $GOPATH/src/github.com/arxdsilva/hc
$ go install
$ hc 10x10
```

libs: uses only check.v1 as personal choice to test.

testing:
```shell 
$ go test -v .
```