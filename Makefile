# note: call scripts from /scripts
check_swagger: 
	which swag || go get -u github.com/swaggo/swag/cmd/swag

swagger: check_swagger
	swag init
	