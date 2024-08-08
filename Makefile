test:
	go test -v -bench=. -benchmem ./$(name)/... > ./$(name)/out.txt
	cat ./$(name)/out.txt
