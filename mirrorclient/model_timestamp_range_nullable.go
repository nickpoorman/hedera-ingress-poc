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

// checks if the TimestampRangeNullable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimestampRangeNullable{}

// TimestampRangeNullable A timestamp range an entity is valid for
type TimestampRangeNullable struct {
	From *string `json:"from,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewTimestampRangeNullable instantiates a new TimestampRangeNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestampRangeNullable() *TimestampRangeNullable {
	this := TimestampRangeNullable{}
	return &this
}

// NewTimestampRangeNullableWithDefaults instantiates a new TimestampRangeNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampRangeNullableWithDefaults() *TimestampRangeNullable {
	this := TimestampRangeNullable{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *TimestampRangeNullable) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampRangeNullable) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *TimestampRangeNullable) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *TimestampRangeNullable) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *TimestampRangeNullable) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimestampRangeNullable) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *TimestampRangeNullable) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *TimestampRangeNullable) SetTo(v string) {
	o.To = &v
}

func (o TimestampRangeNullable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimestampRangeNullable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableTimestampRangeNullable struct {
	value *TimestampRangeNullable
	isSet bool
}

func (v NullableTimestampRangeNullable) Get() *TimestampRangeNullable {
	return v.value
}

func (v *NullableTimestampRangeNullable) Set(val *TimestampRangeNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestampRangeNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestampRangeNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestampRangeNullable(val *TimestampRangeNullable) *NullableTimestampRangeNullable {
	return &NullableTimestampRangeNullable{value: val, isSet: true}
}

func (v NullableTimestampRangeNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestampRangeNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


