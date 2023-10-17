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

// checks if the RoyaltyFeeAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoyaltyFeeAmount{}

// RoyaltyFeeAmount struct for RoyaltyFeeAmount
type RoyaltyFeeAmount struct {
	Numerator *int64 `json:"numerator,omitempty"`
	Denominator *int64 `json:"denominator,omitempty"`
}

// NewRoyaltyFeeAmount instantiates a new RoyaltyFeeAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoyaltyFeeAmount() *RoyaltyFeeAmount {
	this := RoyaltyFeeAmount{}
	return &this
}

// NewRoyaltyFeeAmountWithDefaults instantiates a new RoyaltyFeeAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoyaltyFeeAmountWithDefaults() *RoyaltyFeeAmount {
	this := RoyaltyFeeAmount{}
	return &this
}

// GetNumerator returns the Numerator field value if set, zero value otherwise.
func (o *RoyaltyFeeAmount) GetNumerator() int64 {
	if o == nil || IsNil(o.Numerator) {
		var ret int64
		return ret
	}
	return *o.Numerator
}

// GetNumeratorOk returns a tuple with the Numerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoyaltyFeeAmount) GetNumeratorOk() (*int64, bool) {
	if o == nil || IsNil(o.Numerator) {
		return nil, false
	}
	return o.Numerator, true
}

// HasNumerator returns a boolean if a field has been set.
func (o *RoyaltyFeeAmount) HasNumerator() bool {
	if o != nil && !IsNil(o.Numerator) {
		return true
	}

	return false
}

// SetNumerator gets a reference to the given int64 and assigns it to the Numerator field.
func (o *RoyaltyFeeAmount) SetNumerator(v int64) {
	o.Numerator = &v
}

// GetDenominator returns the Denominator field value if set, zero value otherwise.
func (o *RoyaltyFeeAmount) GetDenominator() int64 {
	if o == nil || IsNil(o.Denominator) {
		var ret int64
		return ret
	}
	return *o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoyaltyFeeAmount) GetDenominatorOk() (*int64, bool) {
	if o == nil || IsNil(o.Denominator) {
		return nil, false
	}
	return o.Denominator, true
}

// HasDenominator returns a boolean if a field has been set.
func (o *RoyaltyFeeAmount) HasDenominator() bool {
	if o != nil && !IsNil(o.Denominator) {
		return true
	}

	return false
}

// SetDenominator gets a reference to the given int64 and assigns it to the Denominator field.
func (o *RoyaltyFeeAmount) SetDenominator(v int64) {
	o.Denominator = &v
}

func (o RoyaltyFeeAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoyaltyFeeAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numerator) {
		toSerialize["numerator"] = o.Numerator
	}
	if !IsNil(o.Denominator) {
		toSerialize["denominator"] = o.Denominator
	}
	return toSerialize, nil
}

type NullableRoyaltyFeeAmount struct {
	value *RoyaltyFeeAmount
	isSet bool
}

func (v NullableRoyaltyFeeAmount) Get() *RoyaltyFeeAmount {
	return v.value
}

func (v *NullableRoyaltyFeeAmount) Set(val *RoyaltyFeeAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableRoyaltyFeeAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableRoyaltyFeeAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoyaltyFeeAmount(val *RoyaltyFeeAmount) *NullableRoyaltyFeeAmount {
	return &NullableRoyaltyFeeAmount{value: val, isSet: true}
}

func (v NullableRoyaltyFeeAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoyaltyFeeAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


