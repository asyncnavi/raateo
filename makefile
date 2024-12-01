tlwd:
	npm run tailwind:watch
tmplengine:
	templ generate --watch --proxy=http://localhost:3001 

migrate:
	go build -o ./tmp/migrate ./cmd/migrate && ./tmp/migrate
