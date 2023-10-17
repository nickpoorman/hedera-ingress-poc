/*
Hedera Mirror Node REST API

The Mirror Node REST API offers the ability to query cryptocurrency transactions and account information from a Hedera managed mirror node.  Base url: [/api/v1](/api/v1)  OpenAPI Spec: [/api/v1/docs/openapi.yml](/api/v1/docs/openapi.yml)

API version: 0.89.0
Contact: mirrornode@hedera.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mirrorclient

import (
	"encoding/json"
)

// checks if the TokenRelationshipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRelationshipResponse{}

// TokenRelationshipResponse struct for TokenRelationshipResponse
type TokenRelationshipResponse struct {
	Tokens []TokenRelationship `json:"tokens,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewTokenRelationshipResponse instantiates a new TokenRelationshipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRelationshipResponse() *TokenRelationshipResponse {
	this := TokenRelationshipResponse{}
	return &this
}

// NewTokenRelationshipResponseWithDefaults instantiates a new TokenRelationshipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRelationshipResponseWithDefaults() *TokenRelationshipResponse {
	this := TokenRelationshipResponse{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *TokenRelationshipResponse) GetTokens() []TokenRelationship {
	if o == nil || IsNil(o.Tokens) {
		var ret []TokenRelationship
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRelationshipResponse) GetTokensOk() ([]TokenRelationship, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *TokenRelationshipResponse) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []TokenRelationship and assigns it to the Tokens field.
func (o *TokenRelationshipResponse) SetTokens(v []TokenRelationship) {
	o.Tokens = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TokenRelationshipResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRelationshipResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TokenRelationshipResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *TokenRelationshipResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o TokenRelationshipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRelationshipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTokenRelationshipResponse struct {
	value *TokenRelationshipResponse
	isSet bool
}

func (v NullableTokenRelationshipResponse) Get() *TokenRelationshipResponse {
	return v.value
}

func (v *NullableTokenRelationshipResponse) Set(val *TokenRelationshipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRelationshipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRelationshipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRelationshipResponse(val *TokenRelationshipResponse) *NullableTokenRelationshipResponse {
	return &NullableTokenRelationshipResponse{value: val, isSet: true}
}

func (v NullableTokenRelationshipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRelationshipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


