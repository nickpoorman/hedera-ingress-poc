# RoyaltyFeeFallbackFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**DenominatingTokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 

## Methods

### NewRoyaltyFeeFallbackFee

`func NewRoyaltyFeeFallbackFee() *RoyaltyFeeFallbackFee`

NewRoyaltyFeeFallbackFee instantiates a new RoyaltyFeeFallbackFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoyaltyFeeFallbackFeeWithDefaults

`func NewRoyaltyFeeFallbackFeeWithDefaults() *RoyaltyFeeFallbackFee`

NewRoyaltyFeeFallbackFeeWithDefaults instantiates a new RoyaltyFeeFallbackFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RoyaltyFeeFallbackFee) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RoyaltyFeeFallbackFee) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RoyaltyFeeFallbackFee) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RoyaltyFeeFallbackFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDenominatingTokenId

`func (o *RoyaltyFeeFallbackFee) GetDenominatingTokenId() string`

GetDenominatingTokenId returns the DenominatingTokenId field if non-nil, zero value otherwise.

### GetDenominatingTokenIdOk

`func (o *RoyaltyFeeFallbackFee) GetDenominatingTokenIdOk() (*string, bool)`

GetDenominatingTokenIdOk returns a tuple with the DenominatingTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatingTokenId

`func (o *RoyaltyFeeFallbackFee) SetDenominatingTokenId(v string)`

SetDenominatingTokenId sets DenominatingTokenId field to given value.

### HasDenominatingTokenId

`func (o *RoyaltyFeeFallbackFee) HasDenominatingTokenId() bool`

HasDenominatingTokenId returns a boolean if a field has been set.

### SetDenominatingTokenIdNil

`func (o *RoyaltyFeeFallbackFee) SetDenominatingTokenIdNil(b bool)`

 SetDenominatingTokenIdNil sets the value for DenominatingTokenId to be an explicit nil

### UnsetDenominatingTokenId
`func (o *RoyaltyFeeFallbackFee) UnsetDenominatingTokenId()`

UnsetDenominatingTokenId ensures that no value is present for DenominatingTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


