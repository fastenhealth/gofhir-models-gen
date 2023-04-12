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

// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
// Base StructureDefinition for DataRequirement Type: Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement struct {
	Id                     *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	Type                   string                      `bson:"type" json:"type"`
	Profile                []string                    `bson:"profile,omitempty" json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept            `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                  `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	MustSupport            []string                    `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter             []DataRequirementCodeFilter `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter             []DataRequirementDateFilter `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
	Limit                  *int                        `bson:"limit,omitempty" json:"limit,omitempty"`
	Sort                   []DataRequirementSort       `bson:"sort,omitempty" json:"sort,omitempty"`
}

// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data. Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
type DataRequirementCodeFilter struct {
	Id          *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path        *string     `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam *string     `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueSet    *string     `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Code        []Coding    `bson:"code,omitempty" json:"code,omitempty"`
}

// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements. Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
type DataRequirementDateFilter struct {
	Id            *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path          *string     `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam   *string     `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueDateTime *string     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   *Period     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueDuration *Duration   `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
}

// Specifies the order of the results to be returned.
// This element can be used in combination with the sort element to specify quota requirements such as "the most recent 5" or "the highest 5". When multiple sorts are specified, they are applied in the order they appear in the resource.
type DataRequirementSort struct {
	Id        *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string        `bson:"path" json:"path"`
	Direction SortDirection `bson:"direction" json:"direction"`
}
