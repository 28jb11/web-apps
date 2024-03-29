run:
	@go mod tidy
	@npx tailwindcss -i src/input.css -o src/output.css
	@go run main.go
