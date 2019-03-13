package pkg

//go:generate sh -c "docker run --rm  -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) -v $(pwd)/../api/cflapi.yaml:/swagger.yaml quay.io/goswagger/swagger:v0.18.0 generate client -f /swagger.yaml"
