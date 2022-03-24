clean:
	@echo "clean output directory..."
	@rm -rf output/

build:
	@echo "build project..."
	@mkdir -p output
	@go build -o output/bubblediff
	@cp conf/config.json ./output

run: build
	@cd output && ./bubblediff

proxy-run: build
	@export http_proxy=http://127.0.0.1:1087;\
	export https_proxy=http://127.0.0.1:1087;\
	export ALL_PROXY=socks5://127.0.0.1:1080;\
    cd output && ./bubblediff