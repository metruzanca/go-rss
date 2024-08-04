start:
	air &
	templ generate -watch -proxy=http://localhost:3000

build:
	templ generate &
	nix run github:a-h/templ &
	go build -o ./tmp/main .
