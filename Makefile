build_docker:
	docker build -t nicholasjackson/fake-monolith .

push_docker:
	docker push nicholasjackson/fake-monolith

build_linux:
	CGO_ENABLED=0 go build -ldflags='-w -s' -o ./bin/fake-monolith  .
