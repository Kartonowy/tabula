watchTempl:
	@templ generate --watch
watchTailwind:
	@tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
watchAir:
	@air 
watch:
	make -j4 watchTempl watchTailwind watchAir