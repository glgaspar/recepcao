run:
	npm install
	npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css
	go run .

build:
	npm install
	npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css
	CGO_ENABLED=0 GOOS=linux go build -o /recepcao