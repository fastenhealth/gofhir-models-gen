#!/usr/bin/env bash

# generateDefinitions "4.0.1" "http://hl7.org/fhir/R4/definitions.json.zip"
curl -L -o definitions.zip https://www.hl7.org/fhir/R4/definitions.json.zip
unzip definitions.zip profiles-types.json valuesets.json -d fhir
rm definitions.zip
curl -L -o fhir/bundle.json http://hl7.org/fhir/R4/bundle.profile.json
curl -L -o fhir/codesystem.json http://hl7.org/fhir/R4/codesystem.profile.json
curl -L -o fhir/structuredefinition.json http://hl7.org/fhir/R4/structuredefinition.profile.json
curl -L -o fhir/valueset.json http://hl7.org/fhir/R4/valueset.profile.json

go generate ./fhir
