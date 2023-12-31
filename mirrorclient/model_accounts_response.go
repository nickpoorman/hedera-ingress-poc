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

// checks if the AccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsResponse{}

// AccountsResponse struct for AccountsResponse
type AccountsResponse struct {
	Accounts []AccountInfo `json:"accounts"`
	Links Links `json:"links"`
}

// NewAccountsResponse instantiates a new AccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsResponse(accounts []AccountInfo, links Links) *AccountsResponse {
	this := AccountsResponse{}
	this.Accounts = accounts
	this.Links = links
	return &this
}

// NewAccountsResponseWithDefaults instantiates a new AccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsResponseWithDefaults() *AccountsResponse {
	this := AccountsResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *AccountsResponse) GetAccounts() []AccountInfo {
	if o == nil {
		var ret []AccountInfo
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AccountsResponse) GetAccountsOk() ([]AccountInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *AccountsResponse) SetAccounts(v []AccountInfo) {
	o.Accounts = v
}

// GetLinks returns the Links field value
func (o *AccountsResponse) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AccountsResponse) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AccountsResponse) SetLinks(v Links) {
	o.Links = v
}

func (o AccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAccountsResponse struct {
	value *AccountsResponse
	isSet bool
}

func (v NullableAccountsResponse) Get() *AccountsResponse {
	return v.value
}

func (v *NullableAccountsResponse) Set(val *AccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsResponse(val *AccountsResponse) *NullableAccountsResponse {
	return &NullableAccountsResponse{value: val, isSet: true}
}

func (v NullableAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


