build:
	@go build -tags prod -o ./bin/app .

run: build
	@./bin/app

templ:
	@templ generate --watch --proxy=http://localhost:8080

css:
	npx tailwindcss -i ./views/css/app.css -o ./public/app.css --watch

# Just works on windows
clean:
	@del /F /Q -r .\\bin
	@del /F /Q -r .\\tmp