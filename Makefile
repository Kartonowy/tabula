watchTempl:
	@templ generate --watch --proxy="http://localhost:1323" --open-browser=false -v
watchTailwind:
	@tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify --watch
watchAir:
	@air 
watch:
	make -j4 watchTempl watchTailwind watchAir