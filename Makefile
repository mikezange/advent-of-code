build:
	go build -o ./day$(day)/ ./day$(day)/main.go

run: build
	cd ./day$(day) && ./main && rm -rf ./main

new:
	mkdir ./day$(day)