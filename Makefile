compile:
	go build -o "out"

test: compile
	./out

clean:
	rm out