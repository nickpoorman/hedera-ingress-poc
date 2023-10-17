# TransactionId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Nonce** | Pointer to **NullableInt32** |  | [optional] 
**Scheduled** | Pointer to **NullableBool** |  | [optional] 
**TransactionValidStart** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionId

`func NewTransactionId() *TransactionId`

NewTransactionId instantiates a new TransactionId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionIdWithDefaults

`func NewTransactionIdWithDefaults() *TransactionId`

NewTransactionIdWithDefaults instantiates a new TransactionId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransactionId) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionId) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionId) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransactionId) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransactionId) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransactionId) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetNonce

`func (o *TransactionId) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransactionId) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransactionId) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TransactionId) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *TransactionId) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *TransactionId) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetScheduled

`func (o *TransactionId) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *TransactionId) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *TransactionId) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *TransactionId) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### SetScheduledNil

`func (o *TransactionId) SetScheduledNil(b bool)`

 SetScheduledNil sets the value for Scheduled to be an explicit nil

### UnsetScheduled
`func (o *TransactionId) UnsetScheduled()`

UnsetScheduled ensures that no value is present for Scheduled, not even an explicit nil
### GetTransactionValidStart

`func (o *TransactionId) GetTransactionValidStart() string`

GetTransactionValidStart returns the TransactionValidStart field if non-nil, zero value otherwise.

### GetTransactionValidStartOk

`func (o *TransactionId) GetTransactionValidStartOk() (*string, bool)`

GetTransactionValidStartOk returns a tuple with the TransactionValidStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionValidStart

`func (o *TransactionId) SetTransactionValidStart(v string)`

SetTransactionValidStart sets TransactionValidStart field to given value.

### HasTransactionValidStart

`func (o *TransactionId) HasTransactionValidStart() bool`

HasTransactionValidStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


