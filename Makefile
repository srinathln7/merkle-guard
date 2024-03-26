# START: begin
GITHUB_BACKUP_PATH=${HOME}/go/src/github.com/srinathLN7/gitbackup/merkle-gaurd
DOC_PATH=${PWD}/docs
DOC_URL=http://localhost:6060/pkg/github.com/srinathLN7/merkle-gaurd/?m=all
CLIENT_URL=http://localhost:6060/pkg/github.com/srinathLN7/merkle-gaurd/internal/client/?m=all
SERVER_URL=http://localhost:6060/pkg/github.com/srinathLN7/merkle-gaurd/internal/server/?m=all
ZKP_URL=http://localhost:6060/pkg/github.com/srinathLN7/merkle-gaurd/internal/cpzkp/?m=all 

.PHONY: compile
compile:
	protoc api/v2/proto/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.


.PHONY: test
test:
	go test -race -parallel=1 -count=1 ./...


# START: build docs
.PHONY: docs
docs:
	rm -rf ${DOC_PATH}
	mkdir -p ${DOC_PATH}

# build the docs
	godoc -url ${DOC_URL} > ${DOC_PATH}/index.html
	godoc -url ${CLIENT_URL} > ${DOC_PATH}/client.html
	godoc -url ${SERVER_URL} > ${DOC_PATH}/server.html
	godoc -url ${ZKP_URL} > ${DOC_PATH}/zkp.html
		
# END: build docs



.PHONY: gitbackup
gitbackup:
	sudo cp -rf ./.git  ${GITHUB_BACKUP_PATH}

# END: begin
