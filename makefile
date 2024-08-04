start:
	air &
	templ generate -watch -proxy=http://localhost:3000

build:
	nix run github:a-h/templ generate &
	go build -o ./tmp/main .
