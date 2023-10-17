# StakingReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Amount** | **int64** | The number of tinybars awarded | 
**Timestamp** | **string** |  | 

## Methods

### NewStakingReward

`func NewStakingReward(accountId NullableString, amount int64, timestamp string, ) *StakingReward`

NewStakingReward instantiates a new StakingReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakingRewardWithDefaults

`func NewStakingRewardWithDefaults() *StakingReward`

NewStakingRewardWithDefaults instantiates a new StakingReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *StakingReward) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StakingReward) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StakingReward) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *StakingReward) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *StakingReward) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAmount

`func (o *StakingReward) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StakingReward) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StakingReward) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetTimestamp

`func (o *StakingReward) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StakingReward) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StakingReward) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


