# display info about the project
info:
    just -l

generate-kon:
    wsdl2openapi2go --file ../wsdl2openapi/Konnektor-OPB6.json --output ./api --naming naming-kon.json

build-linux:
    GOOS=linux GOARCH=amd64 go build -o ./build/koap-linux-amd64 ./cmd/ti