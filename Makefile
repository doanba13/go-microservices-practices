check_install:
	which swagger || go get https://github.com/go-swagger/go-swagger/tree/master/cmd/swagger

swagger:check_install
	swagger generate spec -o ./swagger.yaml --scan-models
