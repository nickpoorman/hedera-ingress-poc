# TokenDistributionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Balance** | **int64** |  | 

## Methods

### NewTokenDistributionInner

`func NewTokenDistributionInner(account NullableString, balance int64, ) *TokenDistributionInner`

NewTokenDistributionInner instantiates a new TokenDistributionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDistributionInnerWithDefaults

`func NewTokenDistributionInnerWithDefaults() *TokenDistributionInner`

NewTokenDistributionInnerWithDefaults instantiates a new TokenDistributionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TokenDistributionInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TokenDistributionInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TokenDistributionInner) SetAccount(v string)`

SetAccount sets Account field to given value.


### SetAccountNil

`func (o *TokenDistributionInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TokenDistributionInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetBalance

`func (o *TokenDistributionInner) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TokenDistributionInner) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TokenDistributionInner) SetBalance(v int64)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


