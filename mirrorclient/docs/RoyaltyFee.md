# RoyaltyFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllCollectorsAreExempt** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to [**RoyaltyFeeAmount**](RoyaltyFeeAmount.md) |  | [optional] 
**CollectorAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**FallbackFee** | Pointer to [**RoyaltyFeeFallbackFee**](RoyaltyFeeFallbackFee.md) |  | [optional] 

## Methods

### NewRoyaltyFee

`func NewRoyaltyFee() *RoyaltyFee`

NewRoyaltyFee instantiates a new RoyaltyFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoyaltyFeeWithDefaults

`func NewRoyaltyFeeWithDefaults() *RoyaltyFee`

NewRoyaltyFeeWithDefaults instantiates a new RoyaltyFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllCollectorsAreExempt

`func (o *RoyaltyFee) GetAllCollectorsAreExempt() bool`

GetAllCollectorsAreExempt returns the AllCollectorsAreExempt field if non-nil, zero value otherwise.

### GetAllCollectorsAreExemptOk

`func (o *RoyaltyFee) GetAllCollectorsAreExemptOk() (*bool, bool)`

GetAllCollectorsAreExemptOk returns a tuple with the AllCollectorsAreExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllCollectorsAreExempt

`func (o *RoyaltyFee) SetAllCollectorsAreExempt(v bool)`

SetAllCollectorsAreExempt sets AllCollectorsAreExempt field to given value.

### HasAllCollectorsAreExempt

`func (o *RoyaltyFee) HasAllCollectorsAreExempt() bool`

HasAllCollectorsAreExempt returns a boolean if a field has been set.

### GetAmount

`func (o *RoyaltyFee) GetAmount() RoyaltyFeeAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RoyaltyFee) GetAmountOk() (*RoyaltyFeeAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RoyaltyFee) SetAmount(v RoyaltyFeeAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RoyaltyFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCollectorAccountId

`func (o *RoyaltyFee) GetCollectorAccountId() string`

GetCollectorAccountId returns the CollectorAccountId field if non-nil, zero value otherwise.

### GetCollectorAccountIdOk

`func (o *RoyaltyFee) GetCollectorAccountIdOk() (*string, bool)`

GetCollectorAccountIdOk returns a tuple with the CollectorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorAccountId

`func (o *RoyaltyFee) SetCollectorAccountId(v string)`

SetCollectorAccountId sets CollectorAccountId field to given value.

### HasCollectorAccountId

`func (o *RoyaltyFee) HasCollectorAccountId() bool`

HasCollectorAccountId returns a boolean if a field has been set.

### SetCollectorAccountIdNil

`func (o *RoyaltyFee) SetCollectorAccountIdNil(b bool)`

 SetCollectorAccountIdNil sets the value for CollectorAccountId to be an explicit nil

### UnsetCollectorAccountId
`func (o *RoyaltyFee) UnsetCollectorAccountId()`

UnsetCollectorAccountId ensures that no value is present for CollectorAccountId, not even an explicit nil
### GetFallbackFee

`func (o *RoyaltyFee) GetFallbackFee() RoyaltyFeeFallbackFee`

GetFallbackFee returns the FallbackFee field if non-nil, zero value otherwise.

### GetFallbackFeeOk

`func (o *RoyaltyFee) GetFallbackFeeOk() (*RoyaltyFeeFallbackFee, bool)`

GetFallbackFeeOk returns a tuple with the FallbackFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackFee

`func (o *RoyaltyFee) SetFallbackFee(v RoyaltyFeeFallbackFee)`

SetFallbackFee sets FallbackFee field to given value.

### HasFallbackFee

`func (o *RoyaltyFee) HasFallbackFee() bool`

HasFallbackFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


