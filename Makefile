.EXPORT_ALL_VARIABLES:

http_proxy=http://127.0.0.1:1087
https_proxy=http://127.0.0.1:1087
ALL_PROXY=socks5://127.0.0.1:1080

clean:
	@echo "clean output directory..."
	@rm -rf output/

build:
	@echo "build project..."
	@mkdir -p output
	@go build -o output/bubblediff
	@cp config.json ./output

run: build
	@cd output && ./bubblediff
