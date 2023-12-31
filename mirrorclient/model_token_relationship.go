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

// checks if the TokenRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRelationship{}

// TokenRelationship struct for TokenRelationship
type TokenRelationship struct {
	// Specifies if the relationship is implicitly/explicitly associated.
	AutomaticAssociation bool `json:"automatic_association"`
	// For FUNGIBLE_COMMON, the balance that the account holds in the smallest denomination. For NON_FUNGIBLE_UNIQUE, the number of NFTs held by the account.
	Balance int64 `json:"balance"`
	CreatedTimestamp string `json:"created_timestamp"`
	// The Freeze status of the account.
	FreezeStatus string `json:"freeze_status"`
	// The KYC status of the account.
	KycStatus string `json:"kyc_status"`
	// Network entity ID in the format of `shard.realm.num`
	TokenId NullableString `json:"token_id"`
}

// NewTokenRelationship instantiates a new TokenRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRelationship(automaticAssociation bool, balance int64, createdTimestamp string, freezeStatus string, kycStatus string, tokenId NullableString) *TokenRelationship {
	this := TokenRelationship{}
	this.AutomaticAssociation = automaticAssociation
	this.Balance = balance
	this.CreatedTimestamp = createdTimestamp
	this.FreezeStatus = freezeStatus
	this.KycStatus = kycStatus
	this.TokenId = tokenId
	return &this
}

// NewTokenRelationshipWithDefaults instantiates a new TokenRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRelationshipWithDefaults() *TokenRelationship {
	this := TokenRelationship{}
	return &this
}

// GetAutomaticAssociation returns the AutomaticAssociation field value
func (o *TokenRelationship) GetAutomaticAssociation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutomaticAssociation
}

// GetAutomaticAssociationOk returns a tuple with the AutomaticAssociation field value
// and a boolean to check if the value has been set.
func (o *TokenRelationship) GetAutomaticAssociationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutomaticAssociation, true
}

// SetAutomaticAssociation sets field value
func (o *TokenRelationship) SetAutomaticAssociation(v bool) {
	o.AutomaticAssociation = v
}

// GetBalance returns the Balance field value
func (o *TokenRelationship) GetBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *TokenRelationship) GetBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *TokenRelationship) SetBalance(v int64) {
	o.Balance = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *TokenRelationship) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *TokenRelationship) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *TokenRelationship) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetFreezeStatus returns the FreezeStatus field value
func (o *TokenRelationship) GetFreezeStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FreezeStatus
}

// GetFreezeStatusOk returns a tuple with the FreezeStatus field value
// and a boolean to check if the value has been set.
func (o *TokenRelationship) GetFreezeStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreezeStatus, true
}

// SetFreezeStatus sets field value
func (o *TokenRelationship) SetFreezeStatus(v string) {
	o.FreezeStatus = v
}

// GetKycStatus returns the KycStatus field value
func (o *TokenRelationship) GetKycStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value
// and a boolean to check if the value has been set.
func (o *TokenRelationship) GetKycStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KycStatus, true
}

// SetKycStatus sets field value
func (o *TokenRelationship) SetKycStatus(v string) {
	o.KycStatus = v
}

// GetTokenId returns the TokenId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TokenRelationship) GetTokenId() string {
	if o == nil || o.TokenId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenRelationship) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// SetTokenId sets field value
func (o *TokenRelationship) SetTokenId(v string) {
	o.TokenId.Set(&v)
}

func (o TokenRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["automatic_association"] = o.AutomaticAssociation
	toSerialize["balance"] = o.Balance
	toSerialize["created_timestamp"] = o.CreatedTimestamp
	toSerialize["freeze_status"] = o.FreezeStatus
	toSerialize["kyc_status"] = o.KycStatus
	toSerialize["token_id"] = o.TokenId.Get()
	return toSerialize, nil
}

type NullableTokenRelationship struct {
	value *TokenRelationship
	isSet bool
}

func (v NullableTokenRelationship) Get() *TokenRelationship {
	return v.value
}

func (v *NullableTokenRelationship) Set(val *TokenRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRelationship(val *TokenRelationship) *NullableTokenRelationship {
	return &NullableTokenRelationship{value: val, isSet: true}
}

func (v NullableTokenRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


