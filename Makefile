FILES=$(shell find . -not -path "./vendor/*" -type f -name "*.proto")
WORK_PATH=$(shell pwd)

# Install Dependencies
lint:
	protolint lint -fix .


# --go_opt=paths=source_relative
# --go-grpc_opt=paths=source_relative
# --validate_out="lang=go,paths=source_relative:."
# https://github.com/envoyproxy/protoc-gen-validate/issues/498
.PHONY: protoc
protoc:
	@if [[ -e tmp ]]; then rm -rf tmp; fi;
	@mkdir -p tmp/docs
	@for file in $(FILES); do \
		protoc \
			-I . \
			-I ./proto \
			--go_out=./tmp \
			--go-grpc_out=./tmp \
			--validate_out="lang=go,paths=:./tmp" \
			--js_out=import_style=commonjs,binary:./tmp \
			--ts_out=service=grpc-web:./tmp \
			--doc_out=./tmp/docs --doc_opt=html,index.html,source_relative \
			$$file; \
	done
	@if [[ -e gens ]]; then rm -rf gens; fi;
	@mkdir -p gens/go
	@mkdir -p gens/js
	@if [[ -e docs ]]; then rm -rf docs; fi;
	@mv ./tmp/github.com/freenetx/protorepo/gens/go ./gens/
	@rm -rf ./tmp/github.com
	@mv ./tmp/docs/proto ./docs
	@mv ./tmp/proto/* ./gens/js/
	@rm -rf tmp

build: protoc

lint_in_docker:
	/usr/local/bin/docker run --volume "$(WORK_PATH):/workspace" --workdir /workspace yoheimuta/protolint lint ./proto


compile_in_docker: lint_in_docker
	@if [[ -e tmp ]]; then rm -rf tmp; fi;
	@mkdir -p tmp/docs
	@if [[ -e gens ]]; then rm -rf gens; fi;
	@mkdir -p gens/go
	@mkdir -p gens/js
	@if [[ -e docs ]]; then rm -rf docs; fi;
	/usr/local/bin/docker run --rm -v $(WORK_PATH):/work -w /work rvolosatovs/protoc \
		--proto_path=./proto \
		--go_out=./tmp \
		--go-grpc_out=./tmp  \
		--validate_out="lang=go,paths=:./tmp" \
		--js_out=import_style=commonjs,binary:./tmp \
		--ts_out=service=grpc-web:./tmp \
        --doc_out=./tmp/docs --doc_opt=html,index.html,source_relative \
		./proto/share/share.proto $(shell find ./proto ! -name 'share.proto' -name '*.proto')

build_in_docker: compile_in_docker
	@mv ./tmp/github.com/freenetx/protorepo/gens/go ./gens/
	@rm -rf ./tmp/github.com
	@mv ./tmp/docs ./docs
	@mv ./tmp/* ./gens/js/
	@rm -rf tmp
#	$(shell cp -r message go && cp -r service go && find go -depth -name '*.proto' | xargs rm)
#	$(shell find service -depth -name '*.go' | xargs rm && find message -depth -name '*.go' | xargs rm)
