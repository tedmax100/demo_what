.PHONY: bench
bench:
	go test -bench=. -cpuprofile=./cpu.prof 
	echo 'go tool pprof ./cpu.prof >> top 10'