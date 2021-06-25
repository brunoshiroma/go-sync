clean:
	go clean -cache -testcache \
	&& rm -rf webapp/.next \
	&& rm -rf webapp/out \
	&& rm -rf embed_html/_next \
	&& rm go_sync
build:
	cd webapp && yarn build && cd .. && go build cmd/go_sync.go