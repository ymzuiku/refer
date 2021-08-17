base:
	make bench-copy
copy:
	go test ./... -test.run TestCopy
test:
	go test ./...
bench:
	go test -bench="Ref" ./... -test.benchmem
bench-copy:
	go test -bench="Copy" ./... -test.benchmem