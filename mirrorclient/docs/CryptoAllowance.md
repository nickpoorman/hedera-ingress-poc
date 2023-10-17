# CryptoAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **interface{}** | The amount remaining of the original amount granted in tinybars. | [optional] 
**AmountGranted** | Pointer to **interface{}** | The granted amount of the spender&#39;s allowance in tinybars. | [optional] 
**Owner** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Spender** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Timestamp** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 

## Methods

### NewCryptoAllowance

`func NewCryptoAllowance() *CryptoAllowance`

NewCryptoAllowance instantiates a new CryptoAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoAllowanceWithDefaults

`func NewCryptoAllowanceWithDefaults() *CryptoAllowance`

NewCryptoAllowanceWithDefaults instantiates a new CryptoAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CryptoAllowance) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CryptoAllowance) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CryptoAllowance) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CryptoAllowance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *CryptoAllowance) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *CryptoAllowance) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetAmountGranted

`func (o *CryptoAllowance) GetAmountGranted() interface{}`

GetAmountGranted returns the AmountGranted field if non-nil, zero value otherwise.

### GetAmountGrantedOk

`func (o *CryptoAllowance) GetAmountGrantedOk() (*interface{}, bool)`

GetAmountGrantedOk returns a tuple with the AmountGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountGranted

`func (o *CryptoAllowance) SetAmountGranted(v interface{})`

SetAmountGranted sets AmountGranted field to given value.

### HasAmountGranted

`func (o *CryptoAllowance) HasAmountGranted() bool`

HasAmountGranted returns a boolean if a field has been set.

### SetAmountGrantedNil

`func (o *CryptoAllowance) SetAmountGrantedNil(b bool)`

 SetAmountGrantedNil sets the value for AmountGranted to be an explicit nil

### UnsetAmountGranted
`func (o *CryptoAllowance) UnsetAmountGranted()`

UnsetAmountGranted ensures that no value is present for AmountGranted, not even an explicit nil
### GetOwner

`func (o *CryptoAllowance) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CryptoAllowance) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CryptoAllowance) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CryptoAllowance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CryptoAllowance) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CryptoAllowance) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetSpender

`func (o *CryptoAllowance) GetSpender() string`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *CryptoAllowance) GetSpenderOk() (*string, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *CryptoAllowance) SetSpender(v string)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *CryptoAllowance) HasSpender() bool`

HasSpender returns a boolean if a field has been set.

### SetSpenderNil

`func (o *CryptoAllowance) SetSpenderNil(b bool)`

 SetSpenderNil sets the value for Spender to be an explicit nil

### UnsetSpender
`func (o *CryptoAllowance) UnsetSpender()`

UnsetSpender ensures that no value is present for Spender, not even an explicit nil
### GetTimestamp

`func (o *CryptoAllowance) GetTimestamp() TimestampRange`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CryptoAllowance) GetTimestampOk() (*TimestampRange, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CryptoAllowance) SetTimestamp(v TimestampRange)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CryptoAllowance) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


