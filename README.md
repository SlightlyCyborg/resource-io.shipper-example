# Resource IO microservice recipe

This is a repository that implements the ideas expressed in Ewan Valentine's Go microservice blog series: [part1](https://ewanvalentine.io/microservices-in-golang-part-1/), [part2](https://ewanvalentine.io/microservices-in-golang-part-2/), [part3](https://ewanvalentine.io/microservices-in-golang-part-3/), [part4](https://ewanvalentine.io/microservices-in-golang-part-4/), [part5](https://ewanvalentine.io/microservices-in-golang-part-5/), [part6](https://ewanvalentine.io/microservices-in-golang-part-6/), [part7](https://ewanvalentine.io/microservices-in-golang-part-7/)

## Dependencies

Make sure go version is >= 1.6
```
go version
```

### Docker

Follow the [official install instructions](https://docs.docker.com/install/)

### gRPC & protobuf

Install go [gRPC](https://grpc.io/docs/quickstart/go.html)

```bash
go get -u google.golang.org/grpc
```

Install proto binary from github
```bash
mkdir -p ~/3rd-party-repos 
cd ~/3rd-party-repos

wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc-3

mkdir -p ~/bin 
cp protoc-3/bin/protoc ~/bin
```

Make sure protoc is now on PATH
```bash
protoc --version
```


If `command not found` make sure `~/bin` is on path
```bash
printf 'export PATH="$PATH":~/bin' >> ~/.bash_profile
```

Install protoc go plugin
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

Make sure protoc-gen-go is on PATH
```bash
which protoc-gen-go
```
...which should return the path to `protoc-gen-go`. If it doesn't then export Go path

```bash
printf 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
```

### go-micro

Install the go-micro protobuf plugins

```bash
go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
```

go-micro itself will be installed dynamicall when `go get` looks at the source files

We still want to install the micro command-line tools though

```bash
go get -u github.com/micro/micro
```

You also need to install the go-micro docker image

```
docker pull microhq/micro
```

## Changes made to Ewan's code 

### Repository interface

There are only a few changes I made to Ewan's code. The biggest change is that I write an `InMemRepo` implementation of `interface Repository` to persist the `User`, `Vessel`, and `Consignment` data

Here is a snippet of the `User` implementation of the `InMemRepo`

```go
type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

//For now, I am only implementing an in memory version of these calls. This is to help me get practice!

type InMemRepo struct {
	users []*pb.User
}

func (repo *InMemRepo) GetAll() ([]*pb.User, error) {
	return repo.users, nil
}
```

### Imports

Most of the imports refer to Ewan's github repo. These need to be changed to refer to our own code

Here is a snippet of the changes (notice the lib aliased as `pb`): 

```golang
import (
	"log"
	"os"

	pb "github.com/SlightlyCyborg/resource-io.shipper-example/user-service/auth"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)
```

### Build/Run

I did not use Ewans `Makefiles`, instead I made a `build` and a `run` script in each micro service or cli.

### vessel-cli
I made a vessel-cli, since the vessel-service had no testing written in the Ewan blog.

## Usage


```bash
cd ~/go/src
mkdir -p resource-io
cd resource-io
git clone https://github.com/SlightlyCyborg/resource-io.shipper-example.git shipper
cd shipper
```

##### vessel-service
```bash
cd vessel-service
./build
./run
cd ..
```

This starts the vessel-service in a docker container

##### vessel-cli

```bash
cd vessel-client
./build
./run
```

This simply asks the vessel service if a there is a vessel capable of carying the load you need shipped. Of course, The Dagney Taggart is a worthy vessel and will be returned since it still has the capacity to ship more goods.

##### user-service
```bash
cd user-service
./build
./run
cd ..
```

This starts the service in a docker container

##### user-cli

```bash
cd user-cli
./build
./run
cd ..
```

This will interact with the user-service. It creates a user, then logs in and responds with a JWS token.

Copy the user token and paste it into the end of the file `shipper/consignment-cli/run`. You should see a JWT token already in that file, so delete it and put this one in its place.

##### consignment-service

```bash
cd consignment-service
./build
./run
cd ..
```

##### consignment-cli

```bash
cd consignment-cli
./build
./run
cd ..
```

The consignment cli reads the token you saved into the run script and sends it to the consignment-service, along with a request to create your consignment. The consignment-service sends a request to the Auth service, ensures that the token is valid (it has a valid user in it) and then creates the consignment. Creating a consignment involves requesting a vessel to cary it, so the consignment-service also does this in `consignment-service/handler.go`


##### JSON API

Go micro has a JSON api plugin. It runs seperately in its own docker container. The image is provided by [microhq](https://micro.mu) and if you followed the dependency instructions above, you should already have pulled it from the docker repository.

Just run the following
```bash
./start_json_rpc_api
```

Then you can make curl commands:

```
 curl -XPOST -H 'Content-Type: application/json' \ 
    -d '{ "service": "go.micro.srv.auth", "method": "Auth.Auth", "request":  { "email": "your@email.com", "password": "SomePass" } }' \
    http://localhost:8080/rpc
```

