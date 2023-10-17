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

// checks if the NetworkStakeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkStakeResponse{}

// NetworkStakeResponse struct for NetworkStakeResponse
type NetworkStakeResponse struct {
	// The maximum amount of tinybar that can be staked for reward while still achieving the maximum per-hbar reward rate 
	MaxStakeRewarded int64 `json:"max_stake_rewarded"`
	// The maximum reward rate, in tinybars per whole hbar, that any account can receive in a day
	MaxStakingRewardRatePerHbar int64 `json:"max_staking_reward_rate_per_hbar"`
	// The total tinybars to be paid as staking rewards in the ending period, after applying the settings for the 0.0.800 balance threshold and the maximum stake rewarded 
	MaxTotalReward int64 `json:"max_total_reward"`
	// The fraction between zero and one of the network and service fees paid to the node reward account 0.0.801
	NodeRewardFeeFraction float32 `json:"node_reward_fee_fraction"`
	// The amount of the staking reward funds of account 0.0.800 reserved to pay pending rewards that have been earned but not collected 
	ReservedStakingRewards int64 `json:"reserved_staking_rewards"`
	// The unreserved tinybar balance of account 0.0.800 required to achieve the maximum per-hbar reward rate 
	RewardBalanceThreshold int64 `json:"reward_balance_threshold"`
	// The total amount staked to the network in tinybars the start of the current staking period
	StakeTotal int64 `json:"stake_total"`
	StakingPeriod NetworkStakeResponseStakingPeriod `json:"staking_period"`
	// The number of minutes in a staking period
	StakingPeriodDuration int64 `json:"staking_period_duration"`
	// The number of staking periods for which the reward is stored for each node
	StakingPeriodsStored int64 `json:"staking_periods_stored"`
	// The fraction between zero and one of the network and service fees paid to the staking reward account 0.0.800
	StakingRewardFeeFraction float32 `json:"staking_reward_fee_fraction"`
	// The total number of tinybars to be distributed as staking rewards each period
	StakingRewardRate int64 `json:"staking_reward_rate"`
	// The minimum balance of staking reward account 0.0.800 required to active rewards
	StakingStartThreshold int64 `json:"staking_start_threshold"`
	// The unreserved balance of account 0.0.800 at the close of the just-ending period; this value is used to compute the HIP-782 balance ratio 
	UnreservedStakingRewardBalance int64 `json:"unreserved_staking_reward_balance"`
}

// NewNetworkStakeResponse instantiates a new NetworkStakeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkStakeResponse(maxStakeRewarded int64, maxStakingRewardRatePerHbar int64, maxTotalReward int64, nodeRewardFeeFraction float32, reservedStakingRewards int64, rewardBalanceThreshold int64, stakeTotal int64, stakingPeriod NetworkStakeResponseStakingPeriod, stakingPeriodDuration int64, stakingPeriodsStored int64, stakingRewardFeeFraction float32, stakingRewardRate int64, stakingStartThreshold int64, unreservedStakingRewardBalance int64) *NetworkStakeResponse {
	this := NetworkStakeResponse{}
	this.MaxStakeRewarded = maxStakeRewarded
	this.MaxStakingRewardRatePerHbar = maxStakingRewardRatePerHbar
	this.MaxTotalReward = maxTotalReward
	this.NodeRewardFeeFraction = nodeRewardFeeFraction
	this.ReservedStakingRewards = reservedStakingRewards
	this.RewardBalanceThreshold = rewardBalanceThreshold
	this.StakeTotal = stakeTotal
	this.StakingPeriod = stakingPeriod
	this.StakingPeriodDuration = stakingPeriodDuration
	this.StakingPeriodsStored = stakingPeriodsStored
	this.StakingRewardFeeFraction = stakingRewardFeeFraction
	this.StakingRewardRate = stakingRewardRate
	this.StakingStartThreshold = stakingStartThreshold
	this.UnreservedStakingRewardBalance = unreservedStakingRewardBalance
	return &this
}

// NewNetworkStakeResponseWithDefaults instantiates a new NetworkStakeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkStakeResponseWithDefaults() *NetworkStakeResponse {
	this := NetworkStakeResponse{}
	return &this
}

// GetMaxStakeRewarded returns the MaxStakeRewarded field value
func (o *NetworkStakeResponse) GetMaxStakeRewarded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxStakeRewarded
}

// GetMaxStakeRewardedOk returns a tuple with the MaxStakeRewarded field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetMaxStakeRewardedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakeRewarded, true
}

// SetMaxStakeRewarded sets field value
func (o *NetworkStakeResponse) SetMaxStakeRewarded(v int64) {
	o.MaxStakeRewarded = v
}

// GetMaxStakingRewardRatePerHbar returns the MaxStakingRewardRatePerHbar field value
func (o *NetworkStakeResponse) GetMaxStakingRewardRatePerHbar() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxStakingRewardRatePerHbar
}

// GetMaxStakingRewardRatePerHbarOk returns a tuple with the MaxStakingRewardRatePerHbar field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetMaxStakingRewardRatePerHbarOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxStakingRewardRatePerHbar, true
}

// SetMaxStakingRewardRatePerHbar sets field value
func (o *NetworkStakeResponse) SetMaxStakingRewardRatePerHbar(v int64) {
	o.MaxStakingRewardRatePerHbar = v
}

// GetMaxTotalReward returns the MaxTotalReward field value
func (o *NetworkStakeResponse) GetMaxTotalReward() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxTotalReward
}

// GetMaxTotalRewardOk returns a tuple with the MaxTotalReward field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetMaxTotalRewardOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTotalReward, true
}

// SetMaxTotalReward sets field value
func (o *NetworkStakeResponse) SetMaxTotalReward(v int64) {
	o.MaxTotalReward = v
}

// GetNodeRewardFeeFraction returns the NodeRewardFeeFraction field value
func (o *NetworkStakeResponse) GetNodeRewardFeeFraction() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NodeRewardFeeFraction
}

// GetNodeRewardFeeFractionOk returns a tuple with the NodeRewardFeeFraction field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetNodeRewardFeeFractionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeRewardFeeFraction, true
}

// SetNodeRewardFeeFraction sets field value
func (o *NetworkStakeResponse) SetNodeRewardFeeFraction(v float32) {
	o.NodeRewardFeeFraction = v
}

// GetReservedStakingRewards returns the ReservedStakingRewards field value
func (o *NetworkStakeResponse) GetReservedStakingRewards() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReservedStakingRewards
}

// GetReservedStakingRewardsOk returns a tuple with the ReservedStakingRewards field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetReservedStakingRewardsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservedStakingRewards, true
}

// SetReservedStakingRewards sets field value
func (o *NetworkStakeResponse) SetReservedStakingRewards(v int64) {
	o.ReservedStakingRewards = v
}

// GetRewardBalanceThreshold returns the RewardBalanceThreshold field value
func (o *NetworkStakeResponse) GetRewardBalanceThreshold() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RewardBalanceThreshold
}

// GetRewardBalanceThresholdOk returns a tuple with the RewardBalanceThreshold field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetRewardBalanceThresholdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardBalanceThreshold, true
}

// SetRewardBalanceThreshold sets field value
func (o *NetworkStakeResponse) SetRewardBalanceThreshold(v int64) {
	o.RewardBalanceThreshold = v
}

// GetStakeTotal returns the StakeTotal field value
func (o *NetworkStakeResponse) GetStakeTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StakeTotal
}

// GetStakeTotalOk returns a tuple with the StakeTotal field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakeTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakeTotal, true
}

// SetStakeTotal sets field value
func (o *NetworkStakeResponse) SetStakeTotal(v int64) {
	o.StakeTotal = v
}

// GetStakingPeriod returns the StakingPeriod field value
func (o *NetworkStakeResponse) GetStakingPeriod() NetworkStakeResponseStakingPeriod {
	if o == nil {
		var ret NetworkStakeResponseStakingPeriod
		return ret
	}

	return o.StakingPeriod
}

// GetStakingPeriodOk returns a tuple with the StakingPeriod field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakingPeriodOk() (*NetworkStakeResponseStakingPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingPeriod, true
}

// SetStakingPeriod sets field value
func (o *NetworkStakeResponse) SetStakingPeriod(v NetworkStakeResponseStakingPeriod) {
	o.StakingPeriod = v
}

// GetStakingPeriodDuration returns the StakingPeriodDuration field value
func (o *NetworkStakeResponse) GetStakingPeriodDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StakingPeriodDuration
}

// GetStakingPeriodDurationOk returns a tuple with the StakingPeriodDuration field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakingPeriodDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingPeriodDuration, true
}

// SetStakingPeriodDuration sets field value
func (o *NetworkStakeResponse) SetStakingPeriodDuration(v int64) {
	o.StakingPeriodDuration = v
}

// GetStakingPeriodsStored returns the StakingPeriodsStored field value
func (o *NetworkStakeResponse) GetStakingPeriodsStored() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StakingPeriodsStored
}

// GetStakingPeriodsStoredOk returns a tuple with the StakingPeriodsStored field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakingPeriodsStoredOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingPeriodsStored, true
}

// SetStakingPeriodsStored sets field value
func (o *NetworkStakeResponse) SetStakingPeriodsStored(v int64) {
	o.StakingPeriodsStored = v
}

// GetStakingRewardFeeFraction returns the StakingRewardFeeFraction field value
func (o *NetworkStakeResponse) GetStakingRewardFeeFraction() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StakingRewardFeeFraction
}

// GetStakingRewardFeeFractionOk returns a tuple with the StakingRewardFeeFraction field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakingRewardFeeFractionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingRewardFeeFraction, true
}

// SetStakingRewardFeeFraction sets field value
func (o *NetworkStakeResponse) SetStakingRewardFeeFraction(v float32) {
	o.StakingRewardFeeFraction = v
}

// GetStakingRewardRate returns the StakingRewardRate field value
func (o *NetworkStakeResponse) GetStakingRewardRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StakingRewardRate
}

// GetStakingRewardRateOk returns a tuple with the StakingRewardRate field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakingRewardRateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingRewardRate, true
}

// SetStakingRewardRate sets field value
func (o *NetworkStakeResponse) SetStakingRewardRate(v int64) {
	o.StakingRewardRate = v
}

// GetStakingStartThreshold returns the StakingStartThreshold field value
func (o *NetworkStakeResponse) GetStakingStartThreshold() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StakingStartThreshold
}

// GetStakingStartThresholdOk returns a tuple with the StakingStartThreshold field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetStakingStartThresholdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingStartThreshold, true
}

// SetStakingStartThreshold sets field value
func (o *NetworkStakeResponse) SetStakingStartThreshold(v int64) {
	o.StakingStartThreshold = v
}

// GetUnreservedStakingRewardBalance returns the UnreservedStakingRewardBalance field value
func (o *NetworkStakeResponse) GetUnreservedStakingRewardBalance() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnreservedStakingRewardBalance
}

// GetUnreservedStakingRewardBalanceOk returns a tuple with the UnreservedStakingRewardBalance field value
// and a boolean to check if the value has been set.
func (o *NetworkStakeResponse) GetUnreservedStakingRewardBalanceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnreservedStakingRewardBalance, true
}

// SetUnreservedStakingRewardBalance sets field value
func (o *NetworkStakeResponse) SetUnreservedStakingRewardBalance(v int64) {
	o.UnreservedStakingRewardBalance = v
}

func (o NetworkStakeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkStakeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_stake_rewarded"] = o.MaxStakeRewarded
	toSerialize["max_staking_reward_rate_per_hbar"] = o.MaxStakingRewardRatePerHbar
	toSerialize["max_total_reward"] = o.MaxTotalReward
	toSerialize["node_reward_fee_fraction"] = o.NodeRewardFeeFraction
	toSerialize["reserved_staking_rewards"] = o.ReservedStakingRewards
	toSerialize["reward_balance_threshold"] = o.RewardBalanceThreshold
	toSerialize["stake_total"] = o.StakeTotal
	toSerialize["staking_period"] = o.StakingPeriod
	toSerialize["staking_period_duration"] = o.StakingPeriodDuration
	toSerialize["staking_periods_stored"] = o.StakingPeriodsStored
	toSerialize["staking_reward_fee_fraction"] = o.StakingRewardFeeFraction
	toSerialize["staking_reward_rate"] = o.StakingRewardRate
	toSerialize["staking_start_threshold"] = o.StakingStartThreshold
	toSerialize["unreserved_staking_reward_balance"] = o.UnreservedStakingRewardBalance
	return toSerialize, nil
}

type NullableNetworkStakeResponse struct {
	value *NetworkStakeResponse
	isSet bool
}

func (v NullableNetworkStakeResponse) Get() *NetworkStakeResponse {
	return v.value
}

func (v *NullableNetworkStakeResponse) Set(val *NetworkStakeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkStakeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkStakeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkStakeResponse(val *NetworkStakeResponse) *NullableNetworkStakeResponse {
	return &NullableNetworkStakeResponse{value: val, isSet: true}
}

func (v NullableNetworkStakeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkStakeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

