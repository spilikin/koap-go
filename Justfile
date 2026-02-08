# display info about the project
info:
    just -l

generate-kon:
    wsdl2openapi2go --file ../wsdl2openapi/Konnektor-OPB6.json --output ./api --naming naming-kon.json
    