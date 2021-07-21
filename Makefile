bench:
	go test -bench=. -benchmem -count=3 -benchtime=1000x .
