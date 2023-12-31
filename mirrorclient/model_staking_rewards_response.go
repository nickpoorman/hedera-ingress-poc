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

// checks if the StakingRewardsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StakingRewardsResponse{}

// StakingRewardsResponse struct for StakingRewardsResponse
type StakingRewardsResponse struct {
	Rewards []StakingReward `json:"rewards,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewStakingRewardsResponse instantiates a new StakingRewardsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakingRewardsResponse() *StakingRewardsResponse {
	this := StakingRewardsResponse{}
	return &this
}

// NewStakingRewardsResponseWithDefaults instantiates a new StakingRewardsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakingRewardsResponseWithDefaults() *StakingRewardsResponse {
	this := StakingRewardsResponse{}
	return &this
}

// GetRewards returns the Rewards field value if set, zero value otherwise.
func (o *StakingRewardsResponse) GetRewards() []StakingReward {
	if o == nil || IsNil(o.Rewards) {
		var ret []StakingReward
		return ret
	}
	return o.Rewards
}

// GetRewardsOk returns a tuple with the Rewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingRewardsResponse) GetRewardsOk() ([]StakingReward, bool) {
	if o == nil || IsNil(o.Rewards) {
		return nil, false
	}
	return o.Rewards, true
}

// HasRewards returns a boolean if a field has been set.
func (o *StakingRewardsResponse) HasRewards() bool {
	if o != nil && !IsNil(o.Rewards) {
		return true
	}

	return false
}

// SetRewards gets a reference to the given []StakingReward and assigns it to the Rewards field.
func (o *StakingRewardsResponse) SetRewards(v []StakingReward) {
	o.Rewards = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *StakingRewardsResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakingRewardsResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StakingRewardsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *StakingRewardsResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o StakingRewardsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StakingRewardsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rewards) {
		toSerialize["rewards"] = o.Rewards
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableStakingRewardsResponse struct {
	value *StakingRewardsResponse
	isSet bool
}

func (v NullableStakingRewardsResponse) Get() *StakingRewardsResponse {
	return v.value
}

func (v *NullableStakingRewardsResponse) Set(val *StakingRewardsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStakingRewardsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStakingRewardsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakingRewardsResponse(val *StakingRewardsResponse) *NullableStakingRewardsResponse {
	return &NullableStakingRewardsResponse{value: val, isSet: true}
}

func (v NullableStakingRewardsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakingRewardsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


