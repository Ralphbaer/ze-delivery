binary-file=pkg-build

default: build

build:
	go build -o $(binary-file)

test:
	gotestsum --format short-verbose --junitfile junit-report.xml \
	./... \
	-coverprofile=coverage.out

cover-html:
	go tool cover -html=coverage.out -o coverage.html

run: 
	go run main.go

clean:
	rm -f $(binary-file)
	rm -f coverage.out
	rm -f coverage.html
	rm -f junit-report.xml