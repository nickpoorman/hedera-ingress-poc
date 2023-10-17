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

// checks if the ContractActionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractActionsResponse{}

// ContractActionsResponse struct for ContractActionsResponse
type ContractActionsResponse struct {
	Actions []ContractAction `json:"actions,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewContractActionsResponse instantiates a new ContractActionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractActionsResponse() *ContractActionsResponse {
	this := ContractActionsResponse{}
	return &this
}

// NewContractActionsResponseWithDefaults instantiates a new ContractActionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractActionsResponseWithDefaults() *ContractActionsResponse {
	this := ContractActionsResponse{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ContractActionsResponse) GetActions() []ContractAction {
	if o == nil || IsNil(o.Actions) {
		var ret []ContractAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractActionsResponse) GetActionsOk() ([]ContractAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ContractActionsResponse) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ContractAction and assigns it to the Actions field.
func (o *ContractActionsResponse) SetActions(v []ContractAction) {
	o.Actions = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ContractActionsResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractActionsResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ContractActionsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ContractActionsResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o ContractActionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractActionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableContractActionsResponse struct {
	value *ContractActionsResponse
	isSet bool
}

func (v NullableContractActionsResponse) Get() *ContractActionsResponse {
	return v.value
}

func (v *NullableContractActionsResponse) Set(val *ContractActionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractActionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractActionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractActionsResponse(val *ContractActionsResponse) *NullableContractActionsResponse {
	return &NullableContractActionsResponse{value: val, isSet: true}
}

func (v NullableContractActionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractActionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

