all: build-docker-upload build-docker-content

build-docker-upload:
	cd upload-service && $(MAKE) all

build-docker-content:
	cd content-service && $(MAKE) all

clean-all: clean-content clean-upload

clean-content:
	cd content-service && $(MAKE) clean

clean-upload:
	cd upload-service && $(MAKE) clean
