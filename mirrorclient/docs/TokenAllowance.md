# TokenAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount remaining of the original amount granted. | [optional] 
**AmountGranted** | Pointer to **int64** | The granted amount of the spender&#39;s allowance. | [optional] 
**Owner** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Spender** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Timestamp** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 
**TokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 

## Methods

### NewTokenAllowance

`func NewTokenAllowance() *TokenAllowance`

NewTokenAllowance instantiates a new TokenAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAllowanceWithDefaults

`func NewTokenAllowanceWithDefaults() *TokenAllowance`

NewTokenAllowanceWithDefaults instantiates a new TokenAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TokenAllowance) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenAllowance) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenAllowance) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TokenAllowance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountGranted

`func (o *TokenAllowance) GetAmountGranted() int64`

GetAmountGranted returns the AmountGranted field if non-nil, zero value otherwise.

### GetAmountGrantedOk

`func (o *TokenAllowance) GetAmountGrantedOk() (*int64, bool)`

GetAmountGrantedOk returns a tuple with the AmountGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGranted

`func (o *TokenAllowance) SetAmountGranted(v int64)`

SetAmountGranted sets AmountGranted field to given value.

### HasAmountGranted

`func (o *TokenAllowance) HasAmountGranted() bool`

HasAmountGranted returns a boolean if a field has been set.

### GetOwner

`func (o *TokenAllowance) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TokenAllowance) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TokenAllowance) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TokenAllowance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *TokenAllowance) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *TokenAllowance) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSpender

`func (o *TokenAllowance) GetSpender() string`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *TokenAllowance) GetSpenderOk() (*string, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *TokenAllowance) SetSpender(v string)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *TokenAllowance) HasSpender() bool`

HasSpender returns a boolean if a field has been set.

### SetSpenderNil

`func (o *TokenAllowance) SetSpenderNil(b bool)`

 SetSpenderNil sets the value for Spender to be an explicit nil

### UnsetSpender
`func (o *TokenAllowance) UnsetSpender()`

UnsetSpender ensures that no value is present for Spender, not even an explicit nil
### GetTimestamp

`func (o *TokenAllowance) GetTimestamp() TimestampRange`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TokenAllowance) GetTimestampOk() (*TimestampRange, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TokenAllowance) SetTimestamp(v TimestampRange)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TokenAllowance) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTokenId

`func (o *TokenAllowance) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenAllowance) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenAllowance) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenAllowance) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *TokenAllowance) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TokenAllowance) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


