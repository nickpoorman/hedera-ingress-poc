# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**ConsensusTimestamp** | Pointer to **string** |  | [optional] 
**CreatorAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**ExecutedTimestamp** | Pointer to **NullableString** |  | [optional] 
**ExpirationTime** | Pointer to **NullableString** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**PayerAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**ScheduleId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Signatures** | Pointer to [**[]ScheduleSignature**](ScheduleSignature.md) |  | [optional] 
**TransactionBody** | Pointer to **string** |  | [optional] 
**WaitForExpiry** | Pointer to **bool** |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminKey

`func (o *Schedule) GetAdminKey() Key`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *Schedule) GetAdminKeyOk() (*Key, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *Schedule) SetAdminKey(v Key)`

SetAdminKey sets AdminKey field to given value.

### HasAdminKey

`func (o *Schedule) HasAdminKey() bool`

HasAdminKey returns a boolean if a field has been set.

### SetAdminKeyNil

`func (o *Schedule) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *Schedule) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil
### GetConsensusTimestamp

`func (o *Schedule) GetConsensusTimestamp() string`

GetConsensusTimestamp returns the ConsensusTimestamp field if non-nil, zero value otherwise.

### GetConsensusTimestampOk

`func (o *Schedule) GetConsensusTimestampOk() (*string, bool)`

GetConsensusTimestampOk returns a tuple with the ConsensusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusTimestamp

`func (o *Schedule) SetConsensusTimestamp(v string)`

SetConsensusTimestamp sets ConsensusTimestamp field to given value.

### HasConsensusTimestamp

`func (o *Schedule) HasConsensusTimestamp() bool`

HasConsensusTimestamp returns a boolean if a field has been set.

### GetCreatorAccountId

`func (o *Schedule) GetCreatorAccountId() string`

GetCreatorAccountId returns the CreatorAccountId field if non-nil, zero value otherwise.

### GetCreatorAccountIdOk

`func (o *Schedule) GetCreatorAccountIdOk() (*string, bool)`

GetCreatorAccountIdOk returns a tuple with the CreatorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorAccountId

`func (o *Schedule) SetCreatorAccountId(v string)`

SetCreatorAccountId sets CreatorAccountId field to given value.

### HasCreatorAccountId

`func (o *Schedule) HasCreatorAccountId() bool`

HasCreatorAccountId returns a boolean if a field has been set.

### SetCreatorAccountIdNil

`func (o *Schedule) SetCreatorAccountIdNil(b bool)`

 SetCreatorAccountIdNil sets the value for CreatorAccountId to be an explicit nil

### UnsetCreatorAccountId
`func (o *Schedule) UnsetCreatorAccountId()`

UnsetCreatorAccountId ensures that no value is present for CreatorAccountId, not even an explicit nil
### GetDeleted

`func (o *Schedule) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Schedule) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Schedule) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Schedule) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetExecutedTimestamp

`func (o *Schedule) GetExecutedTimestamp() string`

GetExecutedTimestamp returns the ExecutedTimestamp field if non-nil, zero value otherwise.

### GetExecutedTimestampOk

`func (o *Schedule) GetExecutedTimestampOk() (*string, bool)`

GetExecutedTimestampOk returns a tuple with the ExecutedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedTimestamp

`func (o *Schedule) SetExecutedTimestamp(v string)`

SetExecutedTimestamp sets ExecutedTimestamp field to given value.

### HasExecutedTimestamp

`func (o *Schedule) HasExecutedTimestamp() bool`

HasExecutedTimestamp returns a boolean if a field has been set.

### SetExecutedTimestampNil

`func (o *Schedule) SetExecutedTimestampNil(b bool)`

 SetExecutedTimestampNil sets the value for ExecutedTimestamp to be an explicit nil

### UnsetExecutedTimestamp
`func (o *Schedule) UnsetExecutedTimestamp()`

UnsetExecutedTimestamp ensures that no value is present for ExecutedTimestamp, not even an explicit nil
### GetExpirationTime

`func (o *Schedule) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *Schedule) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *Schedule) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *Schedule) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *Schedule) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *Schedule) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetMemo

`func (o *Schedule) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Schedule) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Schedule) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *Schedule) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPayerAccountId

`func (o *Schedule) GetPayerAccountId() string`

GetPayerAccountId returns the PayerAccountId field if non-nil, zero value otherwise.

### GetPayerAccountIdOk

`func (o *Schedule) GetPayerAccountIdOk() (*string, bool)`

GetPayerAccountIdOk returns a tuple with the PayerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAccountId

`func (o *Schedule) SetPayerAccountId(v string)`

SetPayerAccountId sets PayerAccountId field to given value.

### HasPayerAccountId

`func (o *Schedule) HasPayerAccountId() bool`

HasPayerAccountId returns a boolean if a field has been set.

### SetPayerAccountIdNil

`func (o *Schedule) SetPayerAccountIdNil(b bool)`

 SetPayerAccountIdNil sets the value for PayerAccountId to be an explicit nil

### UnsetPayerAccountId
`func (o *Schedule) UnsetPayerAccountId()`

UnsetPayerAccountId ensures that no value is present for PayerAccountId, not even an explicit nil
### GetScheduleId

`func (o *Schedule) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *Schedule) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *Schedule) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *Schedule) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *Schedule) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *Schedule) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetSignatures

`func (o *Schedule) GetSignatures() []ScheduleSignature`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *Schedule) GetSignaturesOk() (*[]ScheduleSignature, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *Schedule) SetSignatures(v []ScheduleSignature)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *Schedule) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetTransactionBody

`func (o *Schedule) GetTransactionBody() string`

GetTransactionBody returns the TransactionBody field if non-nil, zero value otherwise.

### GetTransactionBodyOk

`func (o *Schedule) GetTransactionBodyOk() (*string, bool)`

GetTransactionBodyOk returns a tuple with the TransactionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBody

`func (o *Schedule) SetTransactionBody(v string)`

SetTransactionBody sets TransactionBody field to given value.

### HasTransactionBody

`func (o *Schedule) HasTransactionBody() bool`

HasTransactionBody returns a boolean if a field has been set.

### GetWaitForExpiry

`func (o *Schedule) GetWaitForExpiry() bool`

GetWaitForExpiry returns the WaitForExpiry field if non-nil, zero value otherwise.

### GetWaitForExpiryOk

`func (o *Schedule) GetWaitForExpiryOk() (*bool, bool)`

GetWaitForExpiryOk returns a tuple with the WaitForExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForExpiry

`func (o *Schedule) SetWaitForExpiry(v bool)`

SetWaitForExpiry sets WaitForExpiry field to given value.

### HasWaitForExpiry

`func (o *Schedule) HasWaitForExpiry() bool`

HasWaitForExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


