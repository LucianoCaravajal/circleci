build:
	docker build -t nikoe14/go-test-app .

run: build
	docker run --rm -p 8080:8080 go-test-app

push: build
	docker push nikoe14/go-test-app

deploy:
	kustomize build ${PWD}/.kustomize | kubectl apply -f -