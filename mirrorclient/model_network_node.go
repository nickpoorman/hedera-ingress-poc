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

// checks if the NetworkNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkNode{}

// NetworkNode struct for NetworkNode
type NetworkNode struct {
	// a memo associated with the address book
	Description NullableString `json:"description"`
	// Network entity ID in the format of `shard.realm.num`
	FileId NullableString `json:"file_id"`
	// The maximum stake (rewarded or not rewarded) this node can have as consensus weight
	MaxStake NullableInt64 `json:"max_stake"`
	// memo
	Memo NullableString `json:"memo"`
	// The minimum stake (rewarded or not rewarded) this node must reach before having non-zero consensus weight 
	MinStake NullableInt64 `json:"min_stake"`
	// Network entity ID in the format of `shard.realm.num`
	NodeAccountId NullableString `json:"node_account_id"`
	// An identifier for the node
	NodeId int64 `json:"node_id"`
	// hex encoded hash of the node's TLS certificate
	NodeCertHash NullableString `json:"node_cert_hash"`
	// hex encoded X509 RSA public key used to verify stream file signature
	PublicKey NullableString `json:"public_key"`
	// The total tinybars earned by this node per whole hbar in the last staking period
	RewardRateStart NullableInt64 `json:"reward_rate_start"`
	ServiceEndpoints []ServiceEndpoint `json:"service_endpoints"`
	// The node consensus weight at the beginning of the staking period
	Stake NullableInt64 `json:"stake"`
	// The sum (balance + stakedToMe) for all accounts staked to this node with declineReward=true at the beginning of the staking period 
	StakeNotRewarded NullableInt64 `json:"stake_not_rewarded"`
	// The sum (balance + staked) for all accounts staked to the node that are not declining rewards at the beginning of the staking period 
	StakeRewarded NullableInt64 `json:"stake_rewarded"`
	StakingPeriod NetworkNodeStakingPeriod `json:"staking_period"`
	Timestamp TimestampRange `json:"timestamp"`
}

// NewNetworkNode instantiates a new NetworkNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkNode(description NullableString, fileId NullableString, maxStake NullableInt64, memo NullableString, minStake NullableInt64, nodeAccountId NullableString, nodeId int64, nodeCertHash NullableString, publicKey NullableString, rewardRateStart NullableInt64, serviceEndpoints []ServiceEndpoint, stake NullableInt64, stakeNotRewarded NullableInt64, stakeRewarded NullableInt64, stakingPeriod NetworkNodeStakingPeriod, timestamp TimestampRange) *NetworkNode {
	this := NetworkNode{}
	this.Description = description
	this.FileId = fileId
	this.MaxStake = maxStake
	this.Memo = memo
	this.MinStake = minStake
	this.NodeAccountId = nodeAccountId
	this.NodeId = nodeId
	this.NodeCertHash = nodeCertHash
	this.PublicKey = publicKey
	this.RewardRateStart = rewardRateStart
	this.ServiceEndpoints = serviceEndpoints
	this.Stake = stake
	this.StakeNotRewarded = stakeNotRewarded
	this.StakeRewarded = stakeRewarded
	this.StakingPeriod = stakingPeriod
	this.Timestamp = timestamp
	return &this
}

// NewNetworkNodeWithDefaults instantiates a new NetworkNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkNodeWithDefaults() *NetworkNode {
	this := NetworkNode{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkNode) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *NetworkNode) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetFileId returns the FileId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkNode) GetFileId() string {
	if o == nil || o.FileId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileId.Get()
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileId.Get(), o.FileId.IsSet()
}

// SetFileId sets field value
func (o *NetworkNode) SetFileId(v string) {
	o.FileId.Set(&v)
}

// GetMaxStake returns the MaxStake field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NetworkNode) GetMaxStake() int64 {
	if o == nil || o.MaxStake.Get() == nil {
		var ret int64
		return ret
	}

	return *o.MaxStake.Get()
}

// GetMaxStakeOk returns a tuple with the MaxStake field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetMaxStakeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStake.Get(), o.MaxStake.IsSet()
}

// SetMaxStake sets field value
func (o *NetworkNode) SetMaxStake(v int64) {
	o.MaxStake.Set(&v)
}

// GetMemo returns the Memo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkNode) GetMemo() string {
	if o == nil || o.Memo.Get() == nil {
		var ret string
		return ret
	}

	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// SetMemo sets field value
func (o *NetworkNode) SetMemo(v string) {
	o.Memo.Set(&v)
}

// GetMinStake returns the MinStake field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NetworkNode) GetMinStake() int64 {
	if o == nil || o.MinStake.Get() == nil {
		var ret int64
		return ret
	}

	return *o.MinStake.Get()
}

// GetMinStakeOk returns a tuple with the MinStake field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetMinStakeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinStake.Get(), o.MinStake.IsSet()
}

// SetMinStake sets field value
func (o *NetworkNode) SetMinStake(v int64) {
	o.MinStake.Set(&v)
}

// GetNodeAccountId returns the NodeAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkNode) GetNodeAccountId() string {
	if o == nil || o.NodeAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.NodeAccountId.Get()
}

// GetNodeAccountIdOk returns a tuple with the NodeAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetNodeAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeAccountId.Get(), o.NodeAccountId.IsSet()
}

// SetNodeAccountId sets field value
func (o *NetworkNode) SetNodeAccountId(v string) {
	o.NodeAccountId.Set(&v)
}

// GetNodeId returns the NodeId field value
func (o *NetworkNode) GetNodeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *NetworkNode) GetNodeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *NetworkNode) SetNodeId(v int64) {
	o.NodeId = v
}

// GetNodeCertHash returns the NodeCertHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkNode) GetNodeCertHash() string {
	if o == nil || o.NodeCertHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.NodeCertHash.Get()
}

// GetNodeCertHashOk returns a tuple with the NodeCertHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetNodeCertHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeCertHash.Get(), o.NodeCertHash.IsSet()
}

// SetNodeCertHash sets field value
func (o *NetworkNode) SetNodeCertHash(v string) {
	o.NodeCertHash.Set(&v)
}

// GetPublicKey returns the PublicKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NetworkNode) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// SetPublicKey sets field value
func (o *NetworkNode) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}

// GetRewardRateStart returns the RewardRateStart field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NetworkNode) GetRewardRateStart() int64 {
	if o == nil || o.RewardRateStart.Get() == nil {
		var ret int64
		return ret
	}

	return *o.RewardRateStart.Get()
}

// GetRewardRateStartOk returns a tuple with the RewardRateStart field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetRewardRateStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardRateStart.Get(), o.RewardRateStart.IsSet()
}

// SetRewardRateStart sets field value
func (o *NetworkNode) SetRewardRateStart(v int64) {
	o.RewardRateStart.Set(&v)
}

// GetServiceEndpoints returns the ServiceEndpoints field value
func (o *NetworkNode) GetServiceEndpoints() []ServiceEndpoint {
	if o == nil {
		var ret []ServiceEndpoint
		return ret
	}

	return o.ServiceEndpoints
}

// GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field value
// and a boolean to check if the value has been set.
func (o *NetworkNode) GetServiceEndpointsOk() ([]ServiceEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceEndpoints, true
}

// SetServiceEndpoints sets field value
func (o *NetworkNode) SetServiceEndpoints(v []ServiceEndpoint) {
	o.ServiceEndpoints = v
}

// GetStake returns the Stake field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NetworkNode) GetStake() int64 {
	if o == nil || o.Stake.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Stake.Get()
}

// GetStakeOk returns a tuple with the Stake field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetStakeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stake.Get(), o.Stake.IsSet()
}

// SetStake sets field value
func (o *NetworkNode) SetStake(v int64) {
	o.Stake.Set(&v)
}

// GetStakeNotRewarded returns the StakeNotRewarded field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NetworkNode) GetStakeNotRewarded() int64 {
	if o == nil || o.StakeNotRewarded.Get() == nil {
		var ret int64
		return ret
	}

	return *o.StakeNotRewarded.Get()
}

// GetStakeNotRewardedOk returns a tuple with the StakeNotRewarded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetStakeNotRewardedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakeNotRewarded.Get(), o.StakeNotRewarded.IsSet()
}

// SetStakeNotRewarded sets field value
func (o *NetworkNode) SetStakeNotRewarded(v int64) {
	o.StakeNotRewarded.Set(&v)
}

// GetStakeRewarded returns the StakeRewarded field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *NetworkNode) GetStakeRewarded() int64 {
	if o == nil || o.StakeRewarded.Get() == nil {
		var ret int64
		return ret
	}

	return *o.StakeRewarded.Get()
}

// GetStakeRewardedOk returns a tuple with the StakeRewarded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkNode) GetStakeRewardedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakeRewarded.Get(), o.StakeRewarded.IsSet()
}

// SetStakeRewarded sets field value
func (o *NetworkNode) SetStakeRewarded(v int64) {
	o.StakeRewarded.Set(&v)
}

// GetStakingPeriod returns the StakingPeriod field value
func (o *NetworkNode) GetStakingPeriod() NetworkNodeStakingPeriod {
	if o == nil {
		var ret NetworkNodeStakingPeriod
		return ret
	}

	return o.StakingPeriod
}

// GetStakingPeriodOk returns a tuple with the StakingPeriod field value
// and a boolean to check if the value has been set.
func (o *NetworkNode) GetStakingPeriodOk() (*NetworkNodeStakingPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingPeriod, true
}

// SetStakingPeriod sets field value
func (o *NetworkNode) SetStakingPeriod(v NetworkNodeStakingPeriod) {
	o.StakingPeriod = v
}

// GetTimestamp returns the Timestamp field value
func (o *NetworkNode) GetTimestamp() TimestampRange {
	if o == nil {
		var ret TimestampRange
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *NetworkNode) GetTimestampOk() (*TimestampRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *NetworkNode) SetTimestamp(v TimestampRange) {
	o.Timestamp = v
}

func (o NetworkNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description.Get()
	toSerialize["file_id"] = o.FileId.Get()
	toSerialize["max_stake"] = o.MaxStake.Get()
	toSerialize["memo"] = o.Memo.Get()
	toSerialize["min_stake"] = o.MinStake.Get()
	toSerialize["node_account_id"] = o.NodeAccountId.Get()
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_cert_hash"] = o.NodeCertHash.Get()
	toSerialize["public_key"] = o.PublicKey.Get()
	toSerialize["reward_rate_start"] = o.RewardRateStart.Get()
	toSerialize["service_endpoints"] = o.ServiceEndpoints
	toSerialize["stake"] = o.Stake.Get()
	toSerialize["stake_not_rewarded"] = o.StakeNotRewarded.Get()
	toSerialize["stake_rewarded"] = o.StakeRewarded.Get()
	toSerialize["staking_period"] = o.StakingPeriod
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableNetworkNode struct {
	value *NetworkNode
	isSet bool
}

func (v NullableNetworkNode) Get() *NetworkNode {
	return v.value
}

func (v *NullableNetworkNode) Set(val *NetworkNode) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkNode(val *NetworkNode) *NullableNetworkNode {
	return &NullableNetworkNode{value: val, isSet: true}
}

func (v NullableNetworkNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


