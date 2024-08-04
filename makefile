start:
	air &
	templ generate -watch -proxy=http://localhost:3000

build:
	templ generate &
	go build -o ./tmp/main .
