
.PHONY: build push run

binary:
# build a local binary
	go build -o=./bin/api
# build a binary for windows/amd64
	GOOS=windows GORARCH=amd64 go build -o=./bin/windows_amd64/api

image:
	podman build -f Dockerfile -t quay.io/${QUAY}/matrices

run: build
	./bin/api
	
publish:
	podman push quay.io/${QUAY}/matrices

deploy-docker:
	${OCP_LOGIN}
	oc new-app . --name=matrices --strategy=docker
	sleep 3
	oc start-build matrices --from-dir=./ --follow=true --wait=true
	oc expose service matrices
	echo "Access the application through the link below"
	oc get routes.route.openshift.io | awk '{ print $2 }'

deploy-source:
	${OCP_LOGIN}
	oc new-app . --name=matrices --strategy=source
	sleep 3
	oc start-build matrices --from-dir=./ --follow=true --wait=true
	oc expose service matrices
	echo "Access the application through the link below"
	oc get routes.route.openshift.io | awk '{ print $2 }'

test:
	go test ./... -v

help:
	@echo "Make - Matrix Calculation:"
	@echo "Available targets:"
	@echo "  binary    - Build a binary the application"
	@echo "  Image	- Create an image of the application"
	@echo "  run	- Run the application by executing the binary"
	@echo "  publish 	- Publish the image to quay.io"
	@echo "  deploy-docker	- Deploy the application to OCP using docker build strategy"
	@echo "  deploy-source	- Deploy the applications to OCP using S2I(Source-to-Image) strategy"
	@echo "  test	- Test the application"
	@echo "  help     - Help"
