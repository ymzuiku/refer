test:
	clear
	# go test ./...
	go test ./... -test.run TestCopy
bench:
	go test -bench . -test.benchmem