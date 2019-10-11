## Download all dependencies
```
dep ensure
```
All the dependencies will be downloaded to vendor

Another way to get the dependencies

go get repoNameWithURL

example:
go get github.com/DATA-DOG/go-sqlmock

## To Compile all the code
```
go build ./...
```

## To run all unit test cases

```
go test ./...
```

Result : All Test are passing successfully
```
nagendra.a.kumar$ go test ./...
ok  	github.com/nagendra547/go-db-loadbalancer/dbadmin	0.097s
ok  	github.com/nagendra547/go-db-loadbalancer/dbquery	0.089s
ok  	github.com/nagendra547/go-db-loadbalancer/health	0.073s
ok  	github.com/nagendra547/go-db-loadbalancer/log	0.079s
```

## Troubleshooting
The whole application has been developed and tested for following version
```
nagendra.a.kumar$ go version
go version go1.11.5 darwin/amd64
```
Please configure GOROOT and GOPATH correctly, in case you face any issues.
