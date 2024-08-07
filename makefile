tailwind:
	tailwindcss -w -i lib/views/assets/input.css -o lib/views/assets/output.css

templ:
	templ generate -watch -proxy=http://localhost:3000

dev:
	air &
	@make tailwind &
	@make templ

# Railway & Air build command
build:
	go build -o ./tmp/main ./cmd/

tools:
	@go install github.com/air-verse/air@latest &
	@go install github.com/a-h/templ/cmd/templ@latest
