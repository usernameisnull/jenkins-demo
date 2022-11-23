image = release-ci.daocloud.io/demo/e2e:latest
binary_name = simple-http-server
build-binary:
	go build -o simple-http-server

build-image:
	docker build . -t release-ci.daocloud.io/demo/e2e:latest

push-image:
	docker push release-ci.daocloud.io/demo/e2e:latest