SRC_DIR := .
DST_DIR := tutorialpb
PROTO := $(wildcard $(SRC_DIR)/*.proto)

all: proto

proto: $(PROTO)
	mkdir -p $(DST_DIR)
	# pyi_out: make IDE happy with the generated python class 
	# https://stackoverflow.com/q/78088003/10003645
	protoc -I $(SRC_DIR) --python_out $(DST_DIR) --pyi_out $(DST_DIR) $(PROTO) 


.PHONY: all proto clean

clean:
	rm -rfv $(DST_DIR)