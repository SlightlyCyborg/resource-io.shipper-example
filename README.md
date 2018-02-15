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



