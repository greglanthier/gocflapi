package pkg

//go:generate sh -c "docker run --rm -e GO_POST_PROCESS_FILE='/usr/local/bin/gofmt -w' -v $(pwd):/local -v $(pwd)/../api/cflapi.yaml:/swagger.yaml -v $(pwd)/../api/config.json:/config.json openapitools/openapi-generator-cli:v3.3.4 generate -i /swagger.yaml -c /config.json -g go -o /local/client"
//go:generate go run ../tools/rename/rename.go
