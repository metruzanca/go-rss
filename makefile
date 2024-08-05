start:
	air &
	tailwindcss -w -i lib/views/assets/input.css -o lib/views/assets/output.css &
	templ generate -watch -proxy=http://localhost:3000
