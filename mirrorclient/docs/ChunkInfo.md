# ChunkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialTransactionId** | Pointer to [**TransactionId**](TransactionId.md) |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewChunkInfo

`func NewChunkInfo() *ChunkInfo`

NewChunkInfo instantiates a new ChunkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunkInfoWithDefaults

`func NewChunkInfoWithDefaults() *ChunkInfo`

NewChunkInfoWithDefaults instantiates a new ChunkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialTransactionId

`func (o *ChunkInfo) GetInitialTransactionId() TransactionId`

GetInitialTransactionId returns the InitialTransactionId field if non-nil, zero value otherwise.

### GetInitialTransactionIdOk

`func (o *ChunkInfo) GetInitialTransactionIdOk() (*TransactionId, bool)`

GetInitialTransactionIdOk returns a tuple with the InitialTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialTransactionId

`func (o *ChunkInfo) SetInitialTransactionId(v TransactionId)`

SetInitialTransactionId sets InitialTransactionId field to given value.

### HasInitialTransactionId

`func (o *ChunkInfo) HasInitialTransactionId() bool`

HasInitialTransactionId returns a boolean if a field has been set.

### GetNumber

`func (o *ChunkInfo) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ChunkInfo) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ChunkInfo) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ChunkInfo) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetTotal

`func (o *ChunkInfo) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ChunkInfo) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ChunkInfo) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ChunkInfo) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


