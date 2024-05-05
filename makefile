build:
	go build -buildmode=plugin main.go

test-analyzer:
	go test -run ^TestRun$
