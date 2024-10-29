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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// CodeSystem is documented here http://hl7.org/fhir/StructureDefinition/CodeSystem
// The CodeSystem resource is used to declare the existence of and describe a code system or code system supplement and its key properties, and optionally define a part or all of its content.
type CodeSystem struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                     `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                     `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                     `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                     `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus           `bson:"status" json:"status"`
	Experimental      *bool                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                     `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                     `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail             `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext              `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept           `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                     `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                     `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CaseSensitive     *bool                       `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	ValueSet          *string                     `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	HierarchyMeaning  *CodeSystemHierarchyMeaning `bson:"hierarchyMeaning,omitempty" json:"hierarchyMeaning,omitempty"`
	Compositional     *bool                       `bson:"compositional,omitempty" json:"compositional,omitempty"`
	VersionNeeded     *bool                       `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty"`
	Content           CodeSystemContentMode       `bson:"content" json:"content"`
	Supplements       *string                     `bson:"supplements,omitempty" json:"supplements,omitempty"`
	Count             *int                        `bson:"count,omitempty" json:"count,omitempty"`
	Filter            []CodeSystemFilter          `bson:"filter,omitempty" json:"filter,omitempty"`
	Property          []CodeSystemProperty        `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []CodeSystemConcept         `bson:"concept,omitempty" json:"concept,omitempty"`
}

// A filter that can be used in a value set compose statement when selecting concepts using a filter.
// Note that filters defined in code systems usually require custom code on the part of any terminology engine that will make them available for use in value set filters. For this reason, they are generally only seen in high value published terminologies.
type CodeSystemFilter struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string           `bson:"code" json:"code"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Operator          []FilterOperator `bson:"operator" json:"operator"`
	Value             string           `bson:"value" json:"value"`
}

// A property defines an additional slot through which additional information can be provided about a concept.
type CodeSystemProperty struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string       `bson:"code" json:"code"`
	Uri               *string      `bson:"uri,omitempty" json:"uri,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
	Type              PropertyType `bson:"type" json:"type"`
}

// Concepts that are in the code system. The concept definitions are inherently hierarchical, but the definitions must be consulted to determine what the meanings of the hierarchical relationships are.
// If this is empty, it means that the code system resource does not represent the content of the code system.
type CodeSystemConcept struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                         `bson:"code" json:"code"`
	Display           *string                        `bson:"display,omitempty" json:"display,omitempty"`
	Definition        *string                        `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []CodeSystemConcept            `bson:"concept,omitempty" json:"concept,omitempty"`
}

// Additional representations for the concept - other languages, aliases, specialized purposes, used for particular purposes, etc.
// Concepts have both a ```display``` and an array of ```designation```. The display is equivalent to a special designation with an implied ```designation.use``` of "primary code" and a language equal to the [Resource Language](resource.html#language).
type CodeSystemConceptDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
	Value             string      `bson:"value" json:"value"`
}

// A property value for this concept.
type CodeSystemConceptProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	ValueCode         string      `bson:"valueCode" json:"valueCode"`
	ValueCoding       Coding      `bson:"valueCoding" json:"valueCoding"`
	ValueString       string      `bson:"valueString" json:"valueString"`
	ValueInteger      int         `bson:"valueInteger" json:"valueInteger"`
	ValueBoolean      bool        `bson:"valueBoolean" json:"valueBoolean"`
	ValueDateTime     string      `bson:"valueDateTime" json:"valueDateTime"`
	ValueDecimal      json.Number `bson:"valueDecimal" json:"valueDecimal"`
}

// This function returns resource reference information
func (r CodeSystem) ResourceRef() (string, *string) {
	return "CodeSystem", r.Id
}

// This function returns resource reference information
func (r CodeSystem) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCodeSystem CodeSystem

// MarshalJSON marshals the given CodeSystem as JSON into a byte slice
func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

// UnmarshalCodeSystem unmarshals a CodeSystem.
func UnmarshalCodeSystem(b []byte) (CodeSystem, error) {
	var codeSystem CodeSystem
	if err := json.Unmarshal(b, &codeSystem); err != nil {
		return codeSystem, err
	}
	return codeSystem, nil
}
