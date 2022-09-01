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

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
// A container for a collection of resources.
type Bundle struct {
	Id            *string       `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string       `bson:"language,omitempty" json:"language,omitempty"`
	Identifier    *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type          BundleType    `bson:"type" json:"type"`
	Timestamp     *string       `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total         *int          `bson:"total,omitempty" json:"total,omitempty"`
	Link          []BundleLink  `bson:"link,omitempty" json:"link,omitempty"`
	Entry         []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature     *Signature    `bson:"signature,omitempty" json:"signature,omitempty"`
}

// A series of links that provide context to this bundle.
/*
Both Bundle.link and Bundle.entry.link are defined to support providing additional context when Bundles are used (e.g. [HATEOAS](http://en.wikipedia.org/wiki/HATEOAS)).

Bundle.entry.link corresponds to links found in the HTTP header if the resource in the entry was [read](http.html#read) directly.

This specification defines some specific uses of Bundle.link for [searching](search.html#conformance) and [paging](http.html#paging), but no specific uses for Bundle.entry.link, and no defined function in a transaction - the meaning is implementation specific.
*/
type BundleLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relation          string      `bson:"relation" json:"relation"`
	Url               string      `bson:"url" json:"url"`
}

// An entry in a bundle resource - will either contain a resource or information about a resource (transactions and history only).
type BundleEntry struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []BundleLink         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl           *string              `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource          json.RawMessage      `bson:"resource,omitempty" json:"resource,omitempty"`
	Search            *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request           *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response          *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}

// Information about the search process that lead to the creation of this entry.
type BundleEntrySearch struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              *SearchEntryMode `bson:"mode,omitempty" json:"mode,omitempty"`
	Score             *string          `bson:"score,omitempty" json:"score,omitempty"`
}

// Additional information about how this entry should be processed as part of a transaction or batch.  For history, it shows how the entry was processed to create the version contained in the entry.
type BundleEntryRequest struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method            HTTPVerb    `bson:"method" json:"method"`
	Url               string      `bson:"url" json:"url"`
	IfNoneMatch       *string     `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist       *string     `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}

// Indicates the results of processing the corresponding 'request' entry in the batch or transaction being responded to or what the results of an operation where when returning history.
type BundleEntryResponse struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            string          `bson:"status" json:"status"`
	Location          *string         `bson:"location,omitempty" json:"location,omitempty"`
	Etag              *string         `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified      *string         `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Outcome           json.RawMessage `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

// This function returns resource reference information
func (r Bundle) ResourceRef() (string, *string) {
	return "Bundle", r.Id
}

type OtherBundle Bundle

// MarshalJSON marshals the given Bundle as JSON into a byte slice
func (r Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBundle
		ResourceType string `json:"resourceType"`
	}{
		OtherBundle:  OtherBundle(r),
		ResourceType: "Bundle",
	})
}

// UnmarshalBundle unmarshals a Bundle.
func UnmarshalBundle(b []byte) (Bundle, error) {
	var bundle Bundle
	if err := json.Unmarshal(b, &bundle); err != nil {
		return bundle, err
	}
	return bundle, nil
}
