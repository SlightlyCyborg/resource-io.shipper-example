# Resource IO microservice recipe

## Dependencies

Make sure go version is >= 1.6
```
go version
```

### Docker

Follow the [official install instructions](https://docs.docker.com/install/)

### gRPC & protobuf

Install go [gRPC](https://grpc.io/docs/quickstart/go.html)

```
go get -u google.golang.org/grpc
```

Install proto binary from github
```
mkdir -p ~/3rd-party-repos 
cd ~/3rd-party-repos

wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc-3

mkdir -p ~/bin 
cp protoc-3/bin/protoc ~/bin
```

Make sure protoc is now on PATH
```
protoc --version
```


If `command not found` make sure `~/bin` is on path
```
printf 'export PATH="$PATH":~/bin' >> ~/.bash_profile
```

Install protoc go plugin
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

Make sure protoc-gen-go is on PATH
```
which protoc-gen-go
```
...which should return the path to `protoc-gen-go`. If it doesn't then export Go path

```
printf 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
```

### go-micro

Install the go-micro protobuf plugins

```
go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
```

go-micro itself will be installed dynamicall when `go get` looks at the source files

We still want to install the micro command-line tools though

```
go get -u github.com/micro/micro
```

You also need to install the go-micro docker image

```
docker pull microhq/micro
```



