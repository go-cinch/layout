# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create order service
```
# Create a template project, u can enter a custom project name replace `order`
kratos new order -r https://github.com/go-cinch/layout.git -b dev

# Enter project dir
cd order

# u can change config at configs/config.yaml
# u can change Name/EnvPrefix at cmd/order/main.go

# Run it
cd cmd/order
go run .
```

## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8080:8080 -p 8180:8180 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

