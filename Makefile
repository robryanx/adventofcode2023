run:
	./scripts/run.sh $(DAY)
.PHONY: run

buildall:
	go build -o bin ./days/...
.PHONY: buildall

testall: buildall
	go test ./
.PHONY: testall