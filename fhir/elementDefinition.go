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

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// ElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ElementDefinition
// Base StructureDefinition for ElementDefinition Type: Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition struct {
	Id                  *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path                string                        `bson:"path" json:"path"`
	Representation      []PropertyRepresentation      `bson:"representation,omitempty" json:"representation,omitempty"`
	SliceName           *string                       `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	SliceIsConstraining *bool                         `bson:"sliceIsConstraining,omitempty" json:"sliceIsConstraining,omitempty"`
	Label               *string                       `bson:"label,omitempty" json:"label,omitempty"`
	Code                []Coding                      `bson:"code,omitempty" json:"code,omitempty"`
	Slicing             *ElementDefinitionSlicing     `bson:"slicing,omitempty" json:"slicing,omitempty"`
	Short               *string                       `bson:"short,omitempty" json:"short,omitempty"`
	Definition          *string                       `bson:"definition,omitempty" json:"definition,omitempty"`
	Comment             *string                       `bson:"comment,omitempty" json:"comment,omitempty"`
	Requirements        *string                       `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Alias               []string                      `bson:"alias,omitempty" json:"alias,omitempty"`
	Min                 *int                          `bson:"min,omitempty" json:"min,omitempty"`
	Max                 *string                       `bson:"max,omitempty" json:"max,omitempty"`
	Base                *ElementDefinitionBase        `bson:"base,omitempty" json:"base,omitempty"`
	ContentReference    *string                       `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
	Type                []ElementDefinitionType       `bson:"type,omitempty" json:"type,omitempty"`
	MeaningWhenMissing  *string                       `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	OrderMeaning        *string                       `bson:"orderMeaning,omitempty" json:"orderMeaning,omitempty"`
	Example             []ElementDefinitionExample    `bson:"example,omitempty" json:"example,omitempty"`
	MaxLength           *int                          `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Condition           []string                      `bson:"condition,omitempty" json:"condition,omitempty"`
	Constraint          []ElementDefinitionConstraint `bson:"constraint,omitempty" json:"constraint,omitempty"`
	MustSupport         *bool                         `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	IsModifier          *bool                         `bson:"isModifier,omitempty" json:"isModifier,omitempty"`
	IsModifierReason    *string                       `bson:"isModifierReason,omitempty" json:"isModifierReason,omitempty"`
	IsSummary           *bool                         `bson:"isSummary,omitempty" json:"isSummary,omitempty"`
	Binding             *ElementDefinitionBinding     `bson:"binding,omitempty" json:"binding,omitempty"`
	Mapping             []ElementDefinitionMapping    `bson:"mapping,omitempty" json:"mapping,omitempty"`
}

// Indicates that the element is sliced into a set of alternative definitions (i.e. in a structure definition, there are multiple different constraints on a single element in the base resource). Slicing can be used in any resource that has cardinality ..* on the base resource, or any resource with a choice of types. The set of slices is any elements that come after this in the element sequence that have the same path, until a shorter path occurs (the shorter path terminates the set).
// The first element in the sequence, the one that carries the slicing, is the definition that applies to all the slices. This is based on the unconstrained element, but can apply any constraints as appropriate. This may include the common constraints on the children of the element.
type ElementDefinitionSlicing struct {
	Id            *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	Discriminator []ElementDefinitionSlicingDiscriminator `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	Description   *string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Ordered       *bool                                   `bson:"ordered,omitempty" json:"ordered,omitempty"`
	Rules         SlicingRules                            `bson:"rules" json:"rules"`
}

// Designates which child elements are used to discriminate between the slices when processing an instance. If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
// If there is no discriminator, the content is hard to process, so this should be avoided.
type ElementDefinitionSlicingDiscriminator struct {
	Id        *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      DiscriminatorType `bson:"type" json:"type"`
	Path      string            `bson:"path" json:"path"`
}

// Information about the base definition of the element, provided to make it unnecessary for tools to trace the deviation of the element through the derived and related profiles. When the element definition is not the original definition of an element - i.g. either in a constraint on another type, or for elements from a super type in a snap shot - then the information in provided in the element definition may be different to the base definition. On the original definition of the element, it will be same.
// The base information does not carry any information that could not be determined from the path and related profiles, but making this determination requires both that the related profiles are available, and that the algorithm to determine them be available. For tooling simplicity, the base information must always be populated in element definitions in snap shots, even if it is the same.
type ElementDefinitionBase struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string      `bson:"path" json:"path"`
	Min       int         `bson:"min" json:"min"`
	Max       string      `bson:"max" json:"max"`
}

// The data type or resource that the value of this element is permitted to be.
// The Type of the element can be left blank in a differential constraint, in which case the type is inherited from the resource. Abstract types are not permitted to appear as a type when multiple types are listed.  (I.e. Abstract types cannot be part of a choice).
type ElementDefinitionType struct {
	Id            *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	Code          string                 `bson:"code" json:"code"`
	Profile       []string               `bson:"profile,omitempty" json:"profile,omitempty"`
	TargetProfile []string               `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	Aggregation   []AggregationMode      `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
	Versioning    *ReferenceVersionRules `bson:"versioning,omitempty" json:"versioning,omitempty"`
}

// A sample value for this element demonstrating the type of information that would typically be found in the element.
// Examples will most commonly be present for data where it's not implicitly obvious from either the data type or value set what the values might be.  (I.e. Example values for dates or quantities would generally be unnecessary.)  If the example value is fully populated, the publication tool can generate an instance automatically.
type ElementDefinitionExample struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Label     string      `bson:"label" json:"label"`
}

// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated within the context of the instance.
// Constraints should be declared on the "context" element - the lowest element in the hierarchy that is common to all nodes referenced by the constraint.
type ElementDefinitionConstraint struct {
	Id           *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	Key          string             `bson:"key" json:"key"`
	Requirements *string            `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Severity     ConstraintSeverity `bson:"severity" json:"severity"`
	Human        string             `bson:"human" json:"human"`
	Expression   *string            `bson:"expression,omitempty" json:"expression,omitempty"`
	Xpath        *string            `bson:"xpath,omitempty" json:"xpath,omitempty"`
	Source       *string            `bson:"source,omitempty" json:"source,omitempty"`
}

// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri).
// For a CodeableConcept, when no codes are allowed - only text, use a binding of strength "required" with a description explaining that no coded values are allowed and what sort of information to put in the "text" element.
type ElementDefinitionBinding struct {
	Id          *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Strength    BindingStrength `bson:"strength" json:"strength"`
	Description *string         `bson:"description,omitempty" json:"description,omitempty"`
	ValueSet    *string         `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}

// Identifies a concept from an external specification that roughly corresponds to this element.
// Mappings are not necessarily specific enough for safe translation.
type ElementDefinitionMapping struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Identity  string      `bson:"identity" json:"identity"`
	Language  *string     `bson:"language,omitempty" json:"language,omitempty"`
	Map       string      `bson:"map" json:"map"`
	Comment   *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
