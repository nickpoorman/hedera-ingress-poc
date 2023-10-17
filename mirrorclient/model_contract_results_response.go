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

// checks if the ContractResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractResultsResponse{}

// ContractResultsResponse struct for ContractResultsResponse
type ContractResultsResponse struct {
	Results []ContractResult `json:"results,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewContractResultsResponse instantiates a new ContractResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractResultsResponse() *ContractResultsResponse {
	this := ContractResultsResponse{}
	return &this
}

// NewContractResultsResponseWithDefaults instantiates a new ContractResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractResultsResponseWithDefaults() *ContractResultsResponse {
	this := ContractResultsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ContractResultsResponse) GetResults() []ContractResult {
	if o == nil || IsNil(o.Results) {
		var ret []ContractResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractResultsResponse) GetResultsOk() ([]ContractResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ContractResultsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ContractResult and assigns it to the Results field.
func (o *ContractResultsResponse) SetResults(v []ContractResult) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ContractResultsResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractResultsResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ContractResultsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ContractResultsResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o ContractResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableContractResultsResponse struct {
	value *ContractResultsResponse
	isSet bool
}

func (v NullableContractResultsResponse) Get() *ContractResultsResponse {
	return v.value
}

func (v *NullableContractResultsResponse) Set(val *ContractResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractResultsResponse(val *ContractResultsResponse) *NullableContractResultsResponse {
	return &NullableContractResultsResponse{value: val, isSet: true}
}

func (v NullableContractResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


