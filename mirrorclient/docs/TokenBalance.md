# TokenBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**Balance** | **int64** |  | 

## Methods

### NewTokenBalance

`func NewTokenBalance(tokenId NullableString, balance int64, ) *TokenBalance`

NewTokenBalance instantiates a new TokenBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBalanceWithDefaults

`func NewTokenBalanceWithDefaults() *TokenBalance`

NewTokenBalanceWithDefaults instantiates a new TokenBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenBalance) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenBalance) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenBalance) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### SetTokenIdNil

`func (o *TokenBalance) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TokenBalance) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetBalance

`func (o *TokenBalance) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TokenBalance) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TokenBalance) SetBalance(v int64)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


