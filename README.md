# Golang FHIR Models

This repository contains a tool that can be used to generate FHIR R4 Go Models. 

# Usage

```
# git clone repository

# download structure files from HL7 (zip file containing json files)
./gen-resources.sh

# install the generator
go mod vendor
go install
gofhir-models-gen --help
gofhir-models-gen --version


# generate models (in gofhir-models directory)
cd ../gofhir-models
./gen-resources.sh 
```

