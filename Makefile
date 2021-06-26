.PHONY: check-env

clean:
	go clean -cache -testcache \
	&& rm -rf webapp/.next \
	&& rm -rf webapp/out \
	&& rm -rf embed_html/_next \
	&& rm go_sync*
dep-webapp:
	cd webapp && yarn install
build-webapp: dep-webapp
	cd webapp && yarn build
build-go: check-env
	go build -o ${EXECUTABLE} cmd/go_sync.go
build: build-webapp build-go
start-webapp:
	cd webapp && yarn dev

check-env:
ifndef EXECUTABLE
	$(error EXECUTABLE envvar is undefined)
endif