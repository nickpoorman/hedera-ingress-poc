# StakingRewardTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Amount** | **int64** | The number of tinybars awarded | 

## Methods

### NewStakingRewardTransfer

`func NewStakingRewardTransfer(account NullableString, amount int64, ) *StakingRewardTransfer`

NewStakingRewardTransfer instantiates a new StakingRewardTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakingRewardTransferWithDefaults

`func NewStakingRewardTransferWithDefaults() *StakingRewardTransfer`

NewStakingRewardTransferWithDefaults instantiates a new StakingRewardTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *StakingRewardTransfer) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *StakingRewardTransfer) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *StakingRewardTransfer) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *StakingRewardTransfer) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *StakingRewardTransfer) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAmount

`func (o *StakingRewardTransfer) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StakingRewardTransfer) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StakingRewardTransfer) SetAmount(v int64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


