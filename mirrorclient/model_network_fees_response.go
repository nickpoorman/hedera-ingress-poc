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

// checks if the NetworkFeesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkFeesResponse{}

// NetworkFeesResponse struct for NetworkFeesResponse
type NetworkFeesResponse struct {
	Fees []NetworkFee `json:"fees,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewNetworkFeesResponse instantiates a new NetworkFeesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeesResponse() *NetworkFeesResponse {
	this := NetworkFeesResponse{}
	return &this
}

// NewNetworkFeesResponseWithDefaults instantiates a new NetworkFeesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeesResponseWithDefaults() *NetworkFeesResponse {
	this := NetworkFeesResponse{}
	return &this
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *NetworkFeesResponse) GetFees() []NetworkFee {
	if o == nil || IsNil(o.Fees) {
		var ret []NetworkFee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeesResponse) GetFeesOk() ([]NetworkFee, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *NetworkFeesResponse) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []NetworkFee and assigns it to the Fees field.
func (o *NetworkFeesResponse) SetFees(v []NetworkFee) {
	o.Fees = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *NetworkFeesResponse) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeesResponse) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *NetworkFeesResponse) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *NetworkFeesResponse) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o NetworkFeesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkFeesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableNetworkFeesResponse struct {
	value *NetworkFeesResponse
	isSet bool
}

func (v NullableNetworkFeesResponse) Get() *NetworkFeesResponse {
	return v.value
}

func (v *NullableNetworkFeesResponse) Set(val *NetworkFeesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeesResponse(val *NetworkFeesResponse) *NullableNetworkFeesResponse {
	return &NullableNetworkFeesResponse{value: val, isSet: true}
}

func (v NullableNetworkFeesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


