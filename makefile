build:
	go build -buildmode=plugin main.go parse_epoch_file.go

test-analyzer:
	go test -run ^TestRun$
