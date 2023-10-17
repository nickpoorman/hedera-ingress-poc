# BalanceTokensInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Balance** | Pointer to **int64** |  | [optional] 

## Methods

### NewBalanceTokensInner

`func NewBalanceTokensInner() *BalanceTokensInner`

NewBalanceTokensInner instantiates a new BalanceTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceTokensInnerWithDefaults

`func NewBalanceTokensInnerWithDefaults() *BalanceTokensInner`

NewBalanceTokensInnerWithDefaults instantiates a new BalanceTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *BalanceTokensInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BalanceTokensInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BalanceTokensInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *BalanceTokensInner) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *BalanceTokensInner) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *BalanceTokensInner) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetBalance

`func (o *BalanceTokensInner) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceTokensInner) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceTokensInner) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BalanceTokensInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


