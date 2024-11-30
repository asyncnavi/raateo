tlwd:
	tailwind -i ./views/static/tlwd.css -o public/tlwd.css --watch
tmplengine:
	templ generate --watch --proxy=http://localhost:8080