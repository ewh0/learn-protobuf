# where the .proto files are located
SRC_DIR := .
# where the .pb.go files are generated
DST_DIR := .
PROTO = $(wildcard $(SRC_DIR)/*.proto)
GO_MODULE := github.com/ewh0/learn-protobuf/official-tutorial/go

GO_SRC_DIR := cmd
GO_SRC = $(shell find $(GO_SRC_DIR) -type f -name '*.go')

.PHONY: all clean vet build
all: proto build

proto: $(PROTO)
	protoc -I $(SRC_DIR) --go_out $(DST_DIR) --go_opt=module=$(GO_MODULE) $^

fmt: $(GO_SRC)
	go fmt ./...
	goimports -l -w .

vet: fmt
	go vet ./...
	staticcheck ./...

build: vet add_person
	
add_person: $(wildcard cmd/add_person/*.go)
	go build -o $@ cmd/add_person/add_person.go

clean:
	rm -rvf tutorialpb
	rm -v add_person