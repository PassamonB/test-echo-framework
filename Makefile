debug:
	echo 'Make is OK'
# for dev
run-dev:
	go build
	./test-echo-framework
build-docker:
	docker build -t passamon/test-echo-framework:dev .
run-docker:
	docker run -p 1323:1323 passamon/test-echo-framework:dev

# for CI/CD
build:
	go build
test:
	go test
