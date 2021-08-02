build:
	./make.sh "build"

test:
	./make.sh "test"

clean:
	./make.sh "clean"

lint:
	./make.sh "lint"

checkEnvs:
	./make.sh "checkEnvs"

doc:
	./make.sh "doc"

gen:
	./make.sh "gen"

setup-env:
	cp -R ./github/hooks/* .git/hooks/
