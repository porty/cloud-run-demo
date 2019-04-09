
.PHONY: push-to-docker
push-to-docker: cloud-run-demo-linux
	docker build -t gcr.io/august-impact-91906/cloud-run-demo .
	docker push gcr.io/august-impact-91906/cloud-run-demo

cloud-run-demo-linux: main.go
	GOOS=linux GOARCH=amd64 go build -o cloud-run-demo-linux .
