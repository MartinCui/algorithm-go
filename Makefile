test:
	go test -v -count=1 ./search

benchmark:
	go test -bench=. -count=1 ./search
