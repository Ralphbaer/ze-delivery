default: build

package-name=partner
service-name=partner-service
artifacts-dir=./.artifacts
bin-dir=./.bin

.PHONY: gen
gen:
	go generate ./...

clean_build:
	echo "Cleaning build directory..." && \
	rm -rf $(bin-dir) && \
    mkdir -p $(bin-dir)

build: clean_build
	go build -o $(bin-dir)/$(service-name)

build_linux: clean_build
	GOOS=linux GOARCH=amd64 go build -o $(bin-dir)/$(service-name)

run:
	go run main.go
	
run-air:
	air main.go

test:
	mkdir -p $(artifacts-dir) && \
	gotestsum --format short-verbose --junitfile $(artifacts-dir)/junit-report.xml \
	./... github.com/Ralphbaer/ze-delivery/common... \
	-coverprofile=$(artifacts-dir)/coverage.out
	
cover-html:
	go tool cover -html=$(artifacts-dir)/coverage.out -o $(artifacts-dir)/coverage.html

cover: test cover-html
	open $(artifacts-dir)/coverage.html

clean:
	rm -rf $(bin-dir)
	rm -rf $(artifacts-dir)

image: build_linux
	docker build -t ze-delivery/$(service-name):v1 .

container:
	docker run --env SERVER_ADDRESS=:3000 -it -p 127.0.0.1:3000:3000 ze-delivery/$(service-name):v1 bash

deploy-image:
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 742261599957.dkr.ecr.us-east-2.amazonaws.com/ze-delivery/$(service-name) | bash	
	docker tag ze-delivery/$(service-name):v1 742261599957.dkr.ecr.us-east-2.amazonaws.com/ze-delivery/$(service-name)
	docker push 742261599957.dkr.ecr.us-east-2.amazonaws.com/ze-delivery/$(service-name):latest

localenv:
	sudo rm -rf .localenv/.data && \
	docker-compose -f .localenv/docker-compose.yml up --force-recreate

doc:
	swagger generate spec -o ./gen/swagger.yaml --scan-models \
	-c "github.com/Ralphbaer/ze-delivery/$(package-name)" \
	-c "github.com/Ralphbaer/ze-delivery/common" \
