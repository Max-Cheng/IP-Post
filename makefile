all:
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm,darwin/arm64 \
					-t co1a/ip-post:latest -o type=registry .
