dev:
	air &
	tailwindcss -w -i lib/views/assets/input.css -o lib/views/assets/output.css &
	templ generate -watch -proxy=http://localhost:3000

# Railway & Air build command
build:
	go build -o ./tmp/main ./cmd/

# Railway Deployment start command
start:
	./tmp/main