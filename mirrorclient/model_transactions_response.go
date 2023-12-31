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

// checks if the TransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionsResponse{}

// TransactionsResponse struct for TransactionsResponse
type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewTransactionsResponse instantiates a new TransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsResponse() *TransactionsResponse {
	this := TransactionsResponse{}
	return &this
}

// NewTransactionsResponseWithDefaults instantiates a new TransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsResponseWithDefaults() *TransactionsResponse {
	this := TransactionsResponse{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *TransactionsResponse) GetTransactions() []Transaction {
	if o == nil || IsNil(o.Transactions) {
		var ret []Transaction
		return ret
	}
	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsResponse) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *TransactionsResponse) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given []Transaction and assigns it to the Transactions field.
func (o *TransactionsResponse) SetTransactions(v []Transaction) {
	o.Transactions = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TransactionsResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TransactionsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *TransactionsResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o TransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTransactionsResponse struct {
	value *TransactionsResponse
	isSet bool
}

func (v NullableTransactionsResponse) Get() *TransactionsResponse {
	return v.value
}

func (v *NullableTransactionsResponse) Set(val *TransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsResponse(val *TransactionsResponse) *NullableTransactionsResponse {
	return &NullableTransactionsResponse{value: val, isSet: true}
}

func (v NullableTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


