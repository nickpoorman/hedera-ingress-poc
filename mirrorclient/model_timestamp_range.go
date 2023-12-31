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

// checks if the TimestampRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimestampRange{}

// TimestampRange A timestamp range an entity is valid for
type TimestampRange struct {
	From *string `json:"from,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewTimestampRange instantiates a new TimestampRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestampRange() *TimestampRange {
	this := TimestampRange{}
	return &this
}

// NewTimestampRangeWithDefaults instantiates a new TimestampRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampRangeWithDefaults() *TimestampRange {
	this := TimestampRange{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *TimestampRange) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampRange) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *TimestampRange) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *TimestampRange) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *TimestampRange) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampRange) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *TimestampRange) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *TimestampRange) SetTo(v string) {
	o.To = &v
}

func (o TimestampRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimestampRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableTimestampRange struct {
	value *TimestampRange
	isSet bool
}

func (v NullableTimestampRange) Get() *TimestampRange {
	return v.value
}

func (v *NullableTimestampRange) Set(val *TimestampRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestampRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestampRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestampRange(val *TimestampRange) *NullableTimestampRange {
	return &NullableTimestampRange{value: val, isSet: true}
}

func (v NullableTimestampRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestampRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


