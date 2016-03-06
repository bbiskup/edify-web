#!/bin/bash

ls -l
go vet -x ./...

go test -cpu=1,2,3,4,5,6,7,8 ./...

# avoid excessive output
go test -bench . ./... 2> bench.log
go test -cover ./...
