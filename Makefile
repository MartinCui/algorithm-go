test:
	go test -v -count=1 ./$(folder)

bench:
	go test -bench=. -count=1 ./$(folder)
