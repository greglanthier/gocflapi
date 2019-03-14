package pkg

//go:generate sh -c "docker run --rm -v ${PWD}:/local -v $(pwd)/../api/cflapi.yaml:/swagger.yaml -v $(pwd)/../api/config.json:/config.json swaggerapi/swagger-codegen-cli generate -i /swagger.yaml -c /config.json -l go -o /local/pkg/client"
