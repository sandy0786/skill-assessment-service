# skill-assessment-service
Service is used to create a platform to assess candidates.

## Developers start guide:
This topic will cover what you need to start contributing to this project

## Necessary tools
These tools are necessary to build, test and deploy the services:
* golang 1.17
* gotools

## Suggested integrated development environment
* Visual Studio code

## Plugins required for VS Code
* go 

## Steps to run the microservice
* go run .\main\main.go

## Steps to build the microservice
* go build .\main\main.go

## Tools required for source code documentation
* godoc
* Download godoc : go get -u golang.org/x/tools/...
* open cmd and run : godoc -http=:6060
* Open browser and hit : http://localhost:6060/pkg/

## Running test cases
* Run : go test -tags=all  -v -coverpkg=./... -coverprofile=profile.cov ./...
* To run specific test file : go test -tags=<tag-name>  -v -coverpkg=./... -coverprofile=profile.cov ./...

## Swagger documentation
> swagger generate spec -o ./docs/swagger.json --scan-models
> swagger serve -F=swagger swagger.json
> http://localhost:8084/docs
> https://swagger.io/specification/
> https://goswagger.io/use/spec/model.html
> https://zupzup.org/casbin-http-role-auth/
> https://github.com/casbin/casbin
