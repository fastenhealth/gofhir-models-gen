// Copyright 2022 - Fasten Health
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// CodeSystemContentMode is documented here http://hl7.org/fhir/ValueSet/codesystem-content-mode
type CodeSystemContentMode int

const (
	CodeSystemContentModeNotPresent CodeSystemContentMode = iota
	CodeSystemContentModeExample
	CodeSystemContentModeFragment
	CodeSystemContentModeComplete
	CodeSystemContentModeSupplement
)

func (code CodeSystemContentMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CodeSystemContentMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-present":
		*code = CodeSystemContentModeNotPresent
	case "example":
		*code = CodeSystemContentModeExample
	case "fragment":
		*code = CodeSystemContentModeFragment
	case "complete":
		*code = CodeSystemContentModeComplete
	case "supplement":
		*code = CodeSystemContentModeSupplement
	default:
		return fmt.Errorf("unknown CodeSystemContentMode code `%s`", s)
	}
	return nil
}
func (code CodeSystemContentMode) String() string {
	return code.Code()
}
func (code CodeSystemContentMode) Code() string {
	switch code {
	case CodeSystemContentModeNotPresent:
		return "not-present"
	case CodeSystemContentModeExample:
		return "example"
	case CodeSystemContentModeFragment:
		return "fragment"
	case CodeSystemContentModeComplete:
		return "complete"
	case CodeSystemContentModeSupplement:
		return "supplement"
	}
	return "<unknown>"
}
func (code CodeSystemContentMode) Display() string {
	switch code {
	case CodeSystemContentModeNotPresent:
		return "Not Present"
	case CodeSystemContentModeExample:
		return "Example"
	case CodeSystemContentModeFragment:
		return "Fragment"
	case CodeSystemContentModeComplete:
		return "Complete"
	case CodeSystemContentModeSupplement:
		return "Supplement"
	}
	return "<unknown>"
}
func (code CodeSystemContentMode) Definition() string {
	switch code {
	case CodeSystemContentModeNotPresent:
		return "None of the concepts defined by the code system are included in the code system resource."
	case CodeSystemContentModeExample:
		return "A few representative concepts are included in the code system resource. There is no useful intent in the subset choice and there's no process to make it workable: it's not intended to be workable."
	case CodeSystemContentModeFragment:
		return "A subset of the code system concepts are included in the code system resource. This is a curated subset released for a specific purpose under the governance of the code system steward, and that the intent, bounds and consequences of the fragmentation are clearly defined in the fragment or the code system documentation. Fragments are also known as partitions."
	case CodeSystemContentModeComplete:
		return "All the concepts defined by the code system are included in the code system resource."
	case CodeSystemContentModeSupplement:
		return "The resource doesn't define any new concepts; it just provides additional designations and properties to another code system."
	}
	return "<unknown>"
}
