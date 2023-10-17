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

// checks if the TransactionId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionId{}

// TransactionId struct for TransactionId
type TransactionId struct {
	// Network entity ID in the format of `shard.realm.num`
	AccountId NullableString `json:"account_id,omitempty"`
	Nonce NullableInt32 `json:"nonce,omitempty"`
	Scheduled NullableBool `json:"scheduled,omitempty"`
	TransactionValidStart *string `json:"transaction_valid_start,omitempty"`
}

// NewTransactionId instantiates a new TransactionId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionId() *TransactionId {
	this := TransactionId{}
	return &this
}

// NewTransactionIdWithDefaults instantiates a new TransactionId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionIdWithDefaults() *TransactionId {
	this := TransactionId{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionId) GetAccountId() string {
	if o == nil || IsNil(o.AccountId.Get()) {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionId) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransactionId) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *TransactionId) SetAccountId(v string) {
	o.AccountId.Set(&v)
}
// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *TransactionId) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *TransactionId) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetNonce returns the Nonce field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionId) GetNonce() int32 {
	if o == nil || IsNil(o.Nonce.Get()) {
		var ret int32
		return ret
	}
	return *o.Nonce.Get()
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionId) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nonce.Get(), o.Nonce.IsSet()
}

// HasNonce returns a boolean if a field has been set.
func (o *TransactionId) HasNonce() bool {
	if o != nil && o.Nonce.IsSet() {
		return true
	}

	return false
}

// SetNonce gets a reference to the given NullableInt32 and assigns it to the Nonce field.
func (o *TransactionId) SetNonce(v int32) {
	o.Nonce.Set(&v)
}
// SetNonceNil sets the value for Nonce to be an explicit nil
func (o *TransactionId) SetNonceNil() {
	o.Nonce.Set(nil)
}

// UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
func (o *TransactionId) UnsetNonce() {
	o.Nonce.Unset()
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionId) GetScheduled() bool {
	if o == nil || IsNil(o.Scheduled.Get()) {
		var ret bool
		return ret
	}
	return *o.Scheduled.Get()
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionId) GetScheduledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheduled.Get(), o.Scheduled.IsSet()
}

// HasScheduled returns a boolean if a field has been set.
func (o *TransactionId) HasScheduled() bool {
	if o != nil && o.Scheduled.IsSet() {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given NullableBool and assigns it to the Scheduled field.
func (o *TransactionId) SetScheduled(v bool) {
	o.Scheduled.Set(&v)
}
// SetScheduledNil sets the value for Scheduled to be an explicit nil
func (o *TransactionId) SetScheduledNil() {
	o.Scheduled.Set(nil)
}

// UnsetScheduled ensures that no value is present for Scheduled, not even an explicit nil
func (o *TransactionId) UnsetScheduled() {
	o.Scheduled.Unset()
}

// GetTransactionValidStart returns the TransactionValidStart field value if set, zero value otherwise.
func (o *TransactionId) GetTransactionValidStart() string {
	if o == nil || IsNil(o.TransactionValidStart) {
		var ret string
		return ret
	}
	return *o.TransactionValidStart
}

// GetTransactionValidStartOk returns a tuple with the TransactionValidStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionId) GetTransactionValidStartOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionValidStart) {
		return nil, false
	}
	return o.TransactionValidStart, true
}

// HasTransactionValidStart returns a boolean if a field has been set.
func (o *TransactionId) HasTransactionValidStart() bool {
	if o != nil && !IsNil(o.TransactionValidStart) {
		return true
	}

	return false
}

// SetTransactionValidStart gets a reference to the given string and assigns it to the TransactionValidStart field.
func (o *TransactionId) SetTransactionValidStart(v string) {
	o.TransactionValidStart = &v
}

func (o TransactionId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.Nonce.IsSet() {
		toSerialize["nonce"] = o.Nonce.Get()
	}
	if o.Scheduled.IsSet() {
		toSerialize["scheduled"] = o.Scheduled.Get()
	}
	if !IsNil(o.TransactionValidStart) {
		toSerialize["transaction_valid_start"] = o.TransactionValidStart
	}
	return toSerialize, nil
}

type NullableTransactionId struct {
	value *TransactionId
	isSet bool
}

func (v NullableTransactionId) Get() *TransactionId {
	return v.value
}

func (v *NullableTransactionId) Set(val *TransactionId) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionId) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionId(val *TransactionId) *NullableTransactionId {
	return &NullableTransactionId{value: val, isSet: true}
}

func (v NullableTransactionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


