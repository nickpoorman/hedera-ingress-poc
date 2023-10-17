# FixedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllCollectorsAreExempt** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to **int64** |  | [optional] 
**CollectorAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**DenominatingTokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 

## Methods

### NewFixedFee

`func NewFixedFee() *FixedFee`

NewFixedFee instantiates a new FixedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedFeeWithDefaults

`func NewFixedFeeWithDefaults() *FixedFee`

NewFixedFeeWithDefaults instantiates a new FixedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllCollectorsAreExempt

`func (o *FixedFee) GetAllCollectorsAreExempt() bool`

GetAllCollectorsAreExempt returns the AllCollectorsAreExempt field if non-nil, zero value otherwise.

### GetAllCollectorsAreExemptOk

`func (o *FixedFee) GetAllCollectorsAreExemptOk() (*bool, bool)`

GetAllCollectorsAreExemptOk returns a tuple with the AllCollectorsAreExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCollectorsAreExempt

`func (o *FixedFee) SetAllCollectorsAreExempt(v bool)`

SetAllCollectorsAreExempt sets AllCollectorsAreExempt field to given value.

### HasAllCollectorsAreExempt

`func (o *FixedFee) HasAllCollectorsAreExempt() bool`

HasAllCollectorsAreExempt returns a boolean if a field has been set.

### GetAmount

`func (o *FixedFee) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FixedFee) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FixedFee) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FixedFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCollectorAccountId

`func (o *FixedFee) GetCollectorAccountId() string`

GetCollectorAccountId returns the CollectorAccountId field if non-nil, zero value otherwise.

### GetCollectorAccountIdOk

`func (o *FixedFee) GetCollectorAccountIdOk() (*string, bool)`

GetCollectorAccountIdOk returns a tuple with the CollectorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorAccountId

`func (o *FixedFee) SetCollectorAccountId(v string)`

SetCollectorAccountId sets CollectorAccountId field to given value.

### HasCollectorAccountId

`func (o *FixedFee) HasCollectorAccountId() bool`

HasCollectorAccountId returns a boolean if a field has been set.

### SetCollectorAccountIdNil

`func (o *FixedFee) SetCollectorAccountIdNil(b bool)`

 SetCollectorAccountIdNil sets the value for CollectorAccountId to be an explicit nil

### UnsetCollectorAccountId
`func (o *FixedFee) UnsetCollectorAccountId()`

UnsetCollectorAccountId ensures that no value is present for CollectorAccountId, not even an explicit nil
### GetDenominatingTokenId

`func (o *FixedFee) GetDenominatingTokenId() string`

GetDenominatingTokenId returns the DenominatingTokenId field if non-nil, zero value otherwise.

### GetDenominatingTokenIdOk

`func (o *FixedFee) GetDenominatingTokenIdOk() (*string, bool)`

GetDenominatingTokenIdOk returns a tuple with the DenominatingTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatingTokenId

`func (o *FixedFee) SetDenominatingTokenId(v string)`

SetDenominatingTokenId sets DenominatingTokenId field to given value.

### HasDenominatingTokenId

`func (o *FixedFee) HasDenominatingTokenId() bool`

HasDenominatingTokenId returns a boolean if a field has been set.

### SetDenominatingTokenIdNil

`func (o *FixedFee) SetDenominatingTokenIdNil(b bool)`

 SetDenominatingTokenIdNil sets the value for DenominatingTokenId to be an explicit nil

### UnsetDenominatingTokenId
`func (o *FixedFee) UnsetDenominatingTokenId()`

UnsetDenominatingTokenId ensures that no value is present for DenominatingTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


