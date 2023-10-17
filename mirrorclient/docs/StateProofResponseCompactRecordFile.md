# StateProofResponseCompactRecordFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Head** | **string** | The record file head | 
**StartRunningHashObject** | **string** | The start running hash object | 
**EndRunningHashObject** | **string** | THe end running hash object | 
**HashesBefore** | **[]string** | The hashes of the transactions before the transaction in query, in consensus timestamp ascending order  | 
**HashesAfter** | **[]string** | The hashes of the transactions after the transaction in query, in consensus timestamp ascending order  | 
**RecordStreamObject** | **string** | The record stream object of the transaction in query | 
**BlockNumber** | **NullableString** | The block number, in base64 encoding. Only present if version is 6 | 

## Methods

### NewStateProofResponseCompactRecordFile

`func NewStateProofResponseCompactRecordFile(head string, startRunningHashObject string, endRunningHashObject string, hashesBefore []string, hashesAfter []string, recordStreamObject string, blockNumber NullableString, ) *StateProofResponseCompactRecordFile`

NewStateProofResponseCompactRecordFile instantiates a new StateProofResponseCompactRecordFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateProofResponseCompactRecordFileWithDefaults

`func NewStateProofResponseCompactRecordFileWithDefaults() *StateProofResponseCompactRecordFile`

NewStateProofResponseCompactRecordFileWithDefaults instantiates a new StateProofResponseCompactRecordFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHead

`func (o *StateProofResponseCompactRecordFile) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *StateProofResponseCompactRecordFile) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *StateProofResponseCompactRecordFile) SetHead(v string)`

SetHead sets Head field to given value.


### GetStartRunningHashObject

`func (o *StateProofResponseCompactRecordFile) GetStartRunningHashObject() string`

GetStartRunningHashObject returns the StartRunningHashObject field if non-nil, zero value otherwise.

### GetStartRunningHashObjectOk

`func (o *StateProofResponseCompactRecordFile) GetStartRunningHashObjectOk() (*string, bool)`

GetStartRunningHashObjectOk returns a tuple with the StartRunningHashObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRunningHashObject

`func (o *StateProofResponseCompactRecordFile) SetStartRunningHashObject(v string)`

SetStartRunningHashObject sets StartRunningHashObject field to given value.


### GetEndRunningHashObject

`func (o *StateProofResponseCompactRecordFile) GetEndRunningHashObject() string`

GetEndRunningHashObject returns the EndRunningHashObject field if non-nil, zero value otherwise.

### GetEndRunningHashObjectOk

`func (o *StateProofResponseCompactRecordFile) GetEndRunningHashObjectOk() (*string, bool)`

GetEndRunningHashObjectOk returns a tuple with the EndRunningHashObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndRunningHashObject

`func (o *StateProofResponseCompactRecordFile) SetEndRunningHashObject(v string)`

SetEndRunningHashObject sets EndRunningHashObject field to given value.


### GetHashesBefore

`func (o *StateProofResponseCompactRecordFile) GetHashesBefore() []string`

GetHashesBefore returns the HashesBefore field if non-nil, zero value otherwise.

### GetHashesBeforeOk

`func (o *StateProofResponseCompactRecordFile) GetHashesBeforeOk() (*[]string, bool)`

GetHashesBeforeOk returns a tuple with the HashesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashesBefore

`func (o *StateProofResponseCompactRecordFile) SetHashesBefore(v []string)`

SetHashesBefore sets HashesBefore field to given value.


### GetHashesAfter

`func (o *StateProofResponseCompactRecordFile) GetHashesAfter() []string`

GetHashesAfter returns the HashesAfter field if non-nil, zero value otherwise.

### GetHashesAfterOk

`func (o *StateProofResponseCompactRecordFile) GetHashesAfterOk() (*[]string, bool)`

GetHashesAfterOk returns a tuple with the HashesAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashesAfter

`func (o *StateProofResponseCompactRecordFile) SetHashesAfter(v []string)`

SetHashesAfter sets HashesAfter field to given value.


### GetRecordStreamObject

`func (o *StateProofResponseCompactRecordFile) GetRecordStreamObject() string`

GetRecordStreamObject returns the RecordStreamObject field if non-nil, zero value otherwise.

### GetRecordStreamObjectOk

`func (o *StateProofResponseCompactRecordFile) GetRecordStreamObjectOk() (*string, bool)`

GetRecordStreamObjectOk returns a tuple with the RecordStreamObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStreamObject

`func (o *StateProofResponseCompactRecordFile) SetRecordStreamObject(v string)`

SetRecordStreamObject sets RecordStreamObject field to given value.


### GetBlockNumber

`func (o *StateProofResponseCompactRecordFile) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *StateProofResponseCompactRecordFile) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *StateProofResponseCompactRecordFile) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### SetBlockNumberNil

`func (o *StateProofResponseCompactRecordFile) SetBlockNumberNil(b bool)`

 SetBlockNumberNil sets the value for BlockNumber to be an explicit nil

### UnsetBlockNumber
`func (o *StateProofResponseCompactRecordFile) UnsetBlockNumber()`

UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


