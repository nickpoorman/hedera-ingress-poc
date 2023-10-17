# NetworkNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | a memo associated with the address book | 
**FileId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**MaxStake** | **NullableInt64** | The maximum stake (rewarded or not rewarded) this node can have as consensus weight | 
**Memo** | **NullableString** | memo | 
**MinStake** | **NullableInt64** | The minimum stake (rewarded or not rewarded) this node must reach before having non-zero consensus weight  | 
**NodeAccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**NodeId** | **int64** | An identifier for the node | 
**NodeCertHash** | **NullableString** | hex encoded hash of the node&#39;s TLS certificate | 
**PublicKey** | **NullableString** | hex encoded X509 RSA public key used to verify stream file signature | 
**RewardRateStart** | **NullableInt64** | The total tinybars earned by this node per whole hbar in the last staking period | 
**ServiceEndpoints** | [**[]ServiceEndpoint**](ServiceEndpoint.md) |  | 
**Stake** | **NullableInt64** | The node consensus weight at the beginning of the staking period | 
**StakeNotRewarded** | **NullableInt64** | The sum (balance + stakedToMe) for all accounts staked to this node with declineReward&#x3D;true at the beginning of the staking period  | 
**StakeRewarded** | **NullableInt64** | The sum (balance + staked) for all accounts staked to the node that are not declining rewards at the beginning of the staking period  | 
**StakingPeriod** | [**NetworkNodeStakingPeriod**](NetworkNodeStakingPeriod.md) |  | 
**Timestamp** | [**TimestampRange**](TimestampRange.md) |  | 

## Methods

### NewNetworkNode

`func NewNetworkNode(description NullableString, fileId NullableString, maxStake NullableInt64, memo NullableString, minStake NullableInt64, nodeAccountId NullableString, nodeId int64, nodeCertHash NullableString, publicKey NullableString, rewardRateStart NullableInt64, serviceEndpoints []ServiceEndpoint, stake NullableInt64, stakeNotRewarded NullableInt64, stakeRewarded NullableInt64, stakingPeriod NetworkNodeStakingPeriod, timestamp TimestampRange, ) *NetworkNode`

NewNetworkNode instantiates a new NetworkNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNodeWithDefaults

`func NewNetworkNodeWithDefaults() *NetworkNode`

NewNetworkNodeWithDefaults instantiates a new NetworkNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkNode) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *NetworkNode) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkNode) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFileId

`func (o *NetworkNode) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *NetworkNode) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *NetworkNode) SetFileId(v string)`

SetFileId sets FileId field to given value.


### SetFileIdNil

`func (o *NetworkNode) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *NetworkNode) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetMaxStake

`func (o *NetworkNode) GetMaxStake() int64`

GetMaxStake returns the MaxStake field if non-nil, zero value otherwise.

### GetMaxStakeOk

`func (o *NetworkNode) GetMaxStakeOk() (*int64, bool)`

GetMaxStakeOk returns a tuple with the MaxStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStake

`func (o *NetworkNode) SetMaxStake(v int64)`

SetMaxStake sets MaxStake field to given value.


### SetMaxStakeNil

`func (o *NetworkNode) SetMaxStakeNil(b bool)`

 SetMaxStakeNil sets the value for MaxStake to be an explicit nil

### UnsetMaxStake
`func (o *NetworkNode) UnsetMaxStake()`

UnsetMaxStake ensures that no value is present for MaxStake, not even an explicit nil
### GetMemo

`func (o *NetworkNode) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NetworkNode) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NetworkNode) SetMemo(v string)`

SetMemo sets Memo field to given value.


### SetMemoNil

`func (o *NetworkNode) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *NetworkNode) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMinStake

`func (o *NetworkNode) GetMinStake() int64`

GetMinStake returns the MinStake field if non-nil, zero value otherwise.

### GetMinStakeOk

`func (o *NetworkNode) GetMinStakeOk() (*int64, bool)`

GetMinStakeOk returns a tuple with the MinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStake

`func (o *NetworkNode) SetMinStake(v int64)`

SetMinStake sets MinStake field to given value.


### SetMinStakeNil

`func (o *NetworkNode) SetMinStakeNil(b bool)`

 SetMinStakeNil sets the value for MinStake to be an explicit nil

### UnsetMinStake
`func (o *NetworkNode) UnsetMinStake()`

UnsetMinStake ensures that no value is present for MinStake, not even an explicit nil
### GetNodeAccountId

`func (o *NetworkNode) GetNodeAccountId() string`

GetNodeAccountId returns the NodeAccountId field if non-nil, zero value otherwise.

### GetNodeAccountIdOk

`func (o *NetworkNode) GetNodeAccountIdOk() (*string, bool)`

GetNodeAccountIdOk returns a tuple with the NodeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAccountId

`func (o *NetworkNode) SetNodeAccountId(v string)`

SetNodeAccountId sets NodeAccountId field to given value.


### SetNodeAccountIdNil

`func (o *NetworkNode) SetNodeAccountIdNil(b bool)`

 SetNodeAccountIdNil sets the value for NodeAccountId to be an explicit nil

### UnsetNodeAccountId
`func (o *NetworkNode) UnsetNodeAccountId()`

UnsetNodeAccountId ensures that no value is present for NodeAccountId, not even an explicit nil
### GetNodeId

`func (o *NetworkNode) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NetworkNode) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NetworkNode) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.


### GetNodeCertHash

`func (o *NetworkNode) GetNodeCertHash() string`

GetNodeCertHash returns the NodeCertHash field if non-nil, zero value otherwise.

### GetNodeCertHashOk

`func (o *NetworkNode) GetNodeCertHashOk() (*string, bool)`

GetNodeCertHashOk returns a tuple with the NodeCertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCertHash

`func (o *NetworkNode) SetNodeCertHash(v string)`

SetNodeCertHash sets NodeCertHash field to given value.


### SetNodeCertHashNil

`func (o *NetworkNode) SetNodeCertHashNil(b bool)`

 SetNodeCertHashNil sets the value for NodeCertHash to be an explicit nil

### UnsetNodeCertHash
`func (o *NetworkNode) UnsetNodeCertHash()`

UnsetNodeCertHash ensures that no value is present for NodeCertHash, not even an explicit nil
### GetPublicKey

`func (o *NetworkNode) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *NetworkNode) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *NetworkNode) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *NetworkNode) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *NetworkNode) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetRewardRateStart

`func (o *NetworkNode) GetRewardRateStart() int64`

GetRewardRateStart returns the RewardRateStart field if non-nil, zero value otherwise.

### GetRewardRateStartOk

`func (o *NetworkNode) GetRewardRateStartOk() (*int64, bool)`

GetRewardRateStartOk returns a tuple with the RewardRateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardRateStart

`func (o *NetworkNode) SetRewardRateStart(v int64)`

SetRewardRateStart sets RewardRateStart field to given value.


### SetRewardRateStartNil

`func (o *NetworkNode) SetRewardRateStartNil(b bool)`

 SetRewardRateStartNil sets the value for RewardRateStart to be an explicit nil

### UnsetRewardRateStart
`func (o *NetworkNode) UnsetRewardRateStart()`

UnsetRewardRateStart ensures that no value is present for RewardRateStart, not even an explicit nil
### GetServiceEndpoints

`func (o *NetworkNode) GetServiceEndpoints() []ServiceEndpoint`

GetServiceEndpoints returns the ServiceEndpoints field if non-nil, zero value otherwise.

### GetServiceEndpointsOk

`func (o *NetworkNode) GetServiceEndpointsOk() (*[]ServiceEndpoint, bool)`

GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoints

`func (o *NetworkNode) SetServiceEndpoints(v []ServiceEndpoint)`

SetServiceEndpoints sets ServiceEndpoints field to given value.


### GetStake

`func (o *NetworkNode) GetStake() int64`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *NetworkNode) GetStakeOk() (*int64, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *NetworkNode) SetStake(v int64)`

SetStake sets Stake field to given value.


### SetStakeNil

`func (o *NetworkNode) SetStakeNil(b bool)`

 SetStakeNil sets the value for Stake to be an explicit nil

### UnsetStake
`func (o *NetworkNode) UnsetStake()`

UnsetStake ensures that no value is present for Stake, not even an explicit nil
### GetStakeNotRewarded

`func (o *NetworkNode) GetStakeNotRewarded() int64`

GetStakeNotRewarded returns the StakeNotRewarded field if non-nil, zero value otherwise.

### GetStakeNotRewardedOk

`func (o *NetworkNode) GetStakeNotRewardedOk() (*int64, bool)`

GetStakeNotRewardedOk returns a tuple with the StakeNotRewarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeNotRewarded

`func (o *NetworkNode) SetStakeNotRewarded(v int64)`

SetStakeNotRewarded sets StakeNotRewarded field to given value.


### SetStakeNotRewardedNil

`func (o *NetworkNode) SetStakeNotRewardedNil(b bool)`

 SetStakeNotRewardedNil sets the value for StakeNotRewarded to be an explicit nil

### UnsetStakeNotRewarded
`func (o *NetworkNode) UnsetStakeNotRewarded()`

UnsetStakeNotRewarded ensures that no value is present for StakeNotRewarded, not even an explicit nil
### GetStakeRewarded

`func (o *NetworkNode) GetStakeRewarded() int64`

GetStakeRewarded returns the StakeRewarded field if non-nil, zero value otherwise.

### GetStakeRewardedOk

`func (o *NetworkNode) GetStakeRewardedOk() (*int64, bool)`

GetStakeRewardedOk returns a tuple with the StakeRewarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeRewarded

`func (o *NetworkNode) SetStakeRewarded(v int64)`

SetStakeRewarded sets StakeRewarded field to given value.


### SetStakeRewardedNil

`func (o *NetworkNode) SetStakeRewardedNil(b bool)`

 SetStakeRewardedNil sets the value for StakeRewarded to be an explicit nil

### UnsetStakeRewarded
`func (o *NetworkNode) UnsetStakeRewarded()`

UnsetStakeRewarded ensures that no value is present for StakeRewarded, not even an explicit nil
### GetStakingPeriod

`func (o *NetworkNode) GetStakingPeriod() NetworkNodeStakingPeriod`

GetStakingPeriod returns the StakingPeriod field if non-nil, zero value otherwise.

### GetStakingPeriodOk

`func (o *NetworkNode) GetStakingPeriodOk() (*NetworkNodeStakingPeriod, bool)`

GetStakingPeriodOk returns a tuple with the StakingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingPeriod

`func (o *NetworkNode) SetStakingPeriod(v NetworkNodeStakingPeriod)`

SetStakingPeriod sets StakingPeriod field to given value.


### GetTimestamp

`func (o *NetworkNode) GetTimestamp() TimestampRange`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NetworkNode) GetTimestampOk() (*TimestampRange, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NetworkNode) SetTimestamp(v TimestampRange)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


