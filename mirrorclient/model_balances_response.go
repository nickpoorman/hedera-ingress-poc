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

// checks if the BalancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalancesResponse{}

// BalancesResponse struct for BalancesResponse
type BalancesResponse struct {
	Timestamp NullableString `json:"timestamp,omitempty"`
	Balances []AccountBalance `json:"balances,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewBalancesResponse instantiates a new BalancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancesResponse() *BalancesResponse {
	this := BalancesResponse{}
	return &this
}

// NewBalancesResponseWithDefaults instantiates a new BalancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancesResponseWithDefaults() *BalancesResponse {
	this := BalancesResponse{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalancesResponse) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret string
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancesResponse) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BalancesResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableString and assigns it to the Timestamp field.
func (o *BalancesResponse) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *BalancesResponse) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *BalancesResponse) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *BalancesResponse) GetBalances() []AccountBalance {
	if o == nil || IsNil(o.Balances) {
		var ret []AccountBalance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancesResponse) GetBalancesOk() ([]AccountBalance, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *BalancesResponse) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []AccountBalance and assigns it to the Balances field.
func (o *BalancesResponse) SetBalances(v []AccountBalance) {
	o.Balances = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BalancesResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancesResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BalancesResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *BalancesResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o BalancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBalancesResponse struct {
	value *BalancesResponse
	isSet bool
}

func (v NullableBalancesResponse) Get() *BalancesResponse {
	return v.value
}

func (v *NullableBalancesResponse) Set(val *BalancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancesResponse(val *BalancesResponse) *NullableBalancesResponse {
	return &NullableBalancesResponse{value: val, isSet: true}
}

func (v NullableBalancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


