#!/usr/bin/env bash

curl -o definitions.zip https://www.hl7.org/fhir/definitions.json.zip
unzip definitions.zip profiles-types.json valuesets.json -d fhir
rm definitions.zip
curl -o fhir/bundle.json http://hl7.org/fhir/bundle.profile.json
curl -o fhir/codesystem.json http://hl7.org/fhir/codesystem.profile.json
curl -o fhir/structuredefinition.json http://hl7.org/fhir/structuredefinition.profile.json
curl -o fhir/valueset.json http://hl7.org/fhir/valueset.profile.json

go generate ./fhir
