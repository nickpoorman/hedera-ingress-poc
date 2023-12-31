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

// checks if the TokenBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenBalance{}

// TokenBalance struct for TokenBalance
type TokenBalance struct {
	// Network entity ID in the format of `shard.realm.num`
	TokenId NullableString `json:"token_id"`
	Balance int64 `json:"balance"`
}

// NewTokenBalance instantiates a new TokenBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenBalance(tokenId NullableString, balance int64) *TokenBalance {
	this := TokenBalance{}
	this.TokenId = tokenId
	this.Balance = balance
	return &this
}

// NewTokenBalanceWithDefaults instantiates a new TokenBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenBalanceWithDefaults() *TokenBalance {
	this := TokenBalance{}
	return &this
}

// GetTokenId returns the TokenId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokenBalance) GetTokenId() string {
	if o == nil || o.TokenId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenBalance) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// SetTokenId sets field value
func (o *TokenBalance) SetTokenId(v string) {
	o.TokenId.Set(&v)
}

// GetBalance returns the Balance field value
func (o *TokenBalance) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *TokenBalance) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *TokenBalance) SetBalance(v int64) {
	o.Balance = v
}

func (o TokenBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_id"] = o.TokenId.Get()
	toSerialize["balance"] = o.Balance
	return toSerialize, nil
}

type NullableTokenBalance struct {
	value *TokenBalance
	isSet bool
}

func (v NullableTokenBalance) Get() *TokenBalance {
	return v.value
}

func (v *NullableTokenBalance) Set(val *TokenBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenBalance(val *TokenBalance) *NullableTokenBalance {
	return &NullableTokenBalance{value: val, isSet: true}
}

func (v NullableTokenBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


