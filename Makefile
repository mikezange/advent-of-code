build:
	go build -o ./day$(day)/ ./day$(day)/main.go

run: build
	cd ./day$(day) && ./main && rm -rf ./main

new:
	mkdir ./day$(day) && touch ./day$(day)/main.go && touch ./day$(day)/input.txt