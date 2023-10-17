# FractionalFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllCollectorsAreExempt** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to [**FractionalFeeAmount**](FractionalFeeAmount.md) |  | [optional] 
**CollectorAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**DenominatingTokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Maximum** | Pointer to **NullableInt64** |  | [optional] 
**Minimum** | Pointer to **int64** |  | [optional] 
**NetOfTransfers** | Pointer to **bool** |  | [optional] 

## Methods

### NewFractionalFee

`func NewFractionalFee() *FractionalFee`

NewFractionalFee instantiates a new FractionalFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFractionalFeeWithDefaults

`func NewFractionalFeeWithDefaults() *FractionalFee`

NewFractionalFeeWithDefaults instantiates a new FractionalFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllCollectorsAreExempt

`func (o *FractionalFee) GetAllCollectorsAreExempt() bool`

GetAllCollectorsAreExempt returns the AllCollectorsAreExempt field if non-nil, zero value otherwise.

### GetAllCollectorsAreExemptOk

`func (o *FractionalFee) GetAllCollectorsAreExemptOk() (*bool, bool)`

GetAllCollectorsAreExemptOk returns a tuple with the AllCollectorsAreExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCollectorsAreExempt

`func (o *FractionalFee) SetAllCollectorsAreExempt(v bool)`

SetAllCollectorsAreExempt sets AllCollectorsAreExempt field to given value.

### HasAllCollectorsAreExempt

`func (o *FractionalFee) HasAllCollectorsAreExempt() bool`

HasAllCollectorsAreExempt returns a boolean if a field has been set.

### GetAmount

`func (o *FractionalFee) GetAmount() FractionalFeeAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FractionalFee) GetAmountOk() (*FractionalFeeAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FractionalFee) SetAmount(v FractionalFeeAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FractionalFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCollectorAccountId

`func (o *FractionalFee) GetCollectorAccountId() string`

GetCollectorAccountId returns the CollectorAccountId field if non-nil, zero value otherwise.

### GetCollectorAccountIdOk

`func (o *FractionalFee) GetCollectorAccountIdOk() (*string, bool)`

GetCollectorAccountIdOk returns a tuple with the CollectorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorAccountId

`func (o *FractionalFee) SetCollectorAccountId(v string)`

SetCollectorAccountId sets CollectorAccountId field to given value.

### HasCollectorAccountId

`func (o *FractionalFee) HasCollectorAccountId() bool`

HasCollectorAccountId returns a boolean if a field has been set.

### SetCollectorAccountIdNil

`func (o *FractionalFee) SetCollectorAccountIdNil(b bool)`

 SetCollectorAccountIdNil sets the value for CollectorAccountId to be an explicit nil

### UnsetCollectorAccountId
`func (o *FractionalFee) UnsetCollectorAccountId()`

UnsetCollectorAccountId ensures that no value is present for CollectorAccountId, not even an explicit nil
### GetDenominatingTokenId

`func (o *FractionalFee) GetDenominatingTokenId() string`

GetDenominatingTokenId returns the DenominatingTokenId field if non-nil, zero value otherwise.

### GetDenominatingTokenIdOk

`func (o *FractionalFee) GetDenominatingTokenIdOk() (*string, bool)`

GetDenominatingTokenIdOk returns a tuple with the DenominatingTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatingTokenId

`func (o *FractionalFee) SetDenominatingTokenId(v string)`

SetDenominatingTokenId sets DenominatingTokenId field to given value.

### HasDenominatingTokenId

`func (o *FractionalFee) HasDenominatingTokenId() bool`

HasDenominatingTokenId returns a boolean if a field has been set.

### SetDenominatingTokenIdNil

`func (o *FractionalFee) SetDenominatingTokenIdNil(b bool)`

 SetDenominatingTokenIdNil sets the value for DenominatingTokenId to be an explicit nil

### UnsetDenominatingTokenId
`func (o *FractionalFee) UnsetDenominatingTokenId()`

UnsetDenominatingTokenId ensures that no value is present for DenominatingTokenId, not even an explicit nil
### GetMaximum

`func (o *FractionalFee) GetMaximum() int64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *FractionalFee) GetMaximumOk() (*int64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *FractionalFee) SetMaximum(v int64)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *FractionalFee) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *FractionalFee) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *FractionalFee) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetMinimum

`func (o *FractionalFee) GetMinimum() int64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *FractionalFee) GetMinimumOk() (*int64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *FractionalFee) SetMinimum(v int64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *FractionalFee) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetNetOfTransfers

`func (o *FractionalFee) GetNetOfTransfers() bool`

GetNetOfTransfers returns the NetOfTransfers field if non-nil, zero value otherwise.

### GetNetOfTransfersOk

`func (o *FractionalFee) GetNetOfTransfersOk() (*bool, bool)`

GetNetOfTransfersOk returns a tuple with the NetOfTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOfTransfers

`func (o *FractionalFee) SetNetOfTransfers(v bool)`

SetNetOfTransfers sets NetOfTransfers field to given value.

### HasNetOfTransfers

`func (o *FractionalFee) HasNetOfTransfers() bool`

HasNetOfTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


