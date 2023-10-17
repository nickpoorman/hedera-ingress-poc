# TransactionDetailAllOfAssessedCustomFees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CollectorAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**EffectivePayerAccountIds** | Pointer to **[]string** |  | [optional] 
**TokenId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 

## Methods

### NewTransactionDetailAllOfAssessedCustomFees

`func NewTransactionDetailAllOfAssessedCustomFees() *TransactionDetailAllOfAssessedCustomFees`

NewTransactionDetailAllOfAssessedCustomFees instantiates a new TransactionDetailAllOfAssessedCustomFees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDetailAllOfAssessedCustomFeesWithDefaults

`func NewTransactionDetailAllOfAssessedCustomFeesWithDefaults() *TransactionDetailAllOfAssessedCustomFees`

NewTransactionDetailAllOfAssessedCustomFeesWithDefaults instantiates a new TransactionDetailAllOfAssessedCustomFees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionDetailAllOfAssessedCustomFees) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionDetailAllOfAssessedCustomFees) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionDetailAllOfAssessedCustomFees) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionDetailAllOfAssessedCustomFees) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCollectorAccountId

`func (o *TransactionDetailAllOfAssessedCustomFees) GetCollectorAccountId() string`

GetCollectorAccountId returns the CollectorAccountId field if non-nil, zero value otherwise.

### GetCollectorAccountIdOk

`func (o *TransactionDetailAllOfAssessedCustomFees) GetCollectorAccountIdOk() (*string, bool)`

GetCollectorAccountIdOk returns a tuple with the CollectorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorAccountId

`func (o *TransactionDetailAllOfAssessedCustomFees) SetCollectorAccountId(v string)`

SetCollectorAccountId sets CollectorAccountId field to given value.

### HasCollectorAccountId

`func (o *TransactionDetailAllOfAssessedCustomFees) HasCollectorAccountId() bool`

HasCollectorAccountId returns a boolean if a field has been set.

### SetCollectorAccountIdNil

`func (o *TransactionDetailAllOfAssessedCustomFees) SetCollectorAccountIdNil(b bool)`

 SetCollectorAccountIdNil sets the value for CollectorAccountId to be an explicit nil

### UnsetCollectorAccountId
`func (o *TransactionDetailAllOfAssessedCustomFees) UnsetCollectorAccountId()`

UnsetCollectorAccountId ensures that no value is present for CollectorAccountId, not even an explicit nil
### GetEffectivePayerAccountIds

`func (o *TransactionDetailAllOfAssessedCustomFees) GetEffectivePayerAccountIds() []*string`

GetEffectivePayerAccountIds returns the EffectivePayerAccountIds field if non-nil, zero value otherwise.

### GetEffectivePayerAccountIdsOk

`func (o *TransactionDetailAllOfAssessedCustomFees) GetEffectivePayerAccountIdsOk() (*[]*string, bool)`

GetEffectivePayerAccountIdsOk returns a tuple with the EffectivePayerAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePayerAccountIds

`func (o *TransactionDetailAllOfAssessedCustomFees) SetEffectivePayerAccountIds(v []*string)`

SetEffectivePayerAccountIds sets EffectivePayerAccountIds field to given value.

### HasEffectivePayerAccountIds

`func (o *TransactionDetailAllOfAssessedCustomFees) HasEffectivePayerAccountIds() bool`

HasEffectivePayerAccountIds returns a boolean if a field has been set.

### GetTokenId

`func (o *TransactionDetailAllOfAssessedCustomFees) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionDetailAllOfAssessedCustomFees) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionDetailAllOfAssessedCustomFees) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TransactionDetailAllOfAssessedCustomFees) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *TransactionDetailAllOfAssessedCustomFees) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *TransactionDetailAllOfAssessedCustomFees) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


