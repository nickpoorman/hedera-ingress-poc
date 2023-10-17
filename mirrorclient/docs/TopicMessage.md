# TopicMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChunkInfo** | Pointer to [**NullableChunkInfo**](ChunkInfo.md) |  | [optional] 
**ConsensusTimestamp** | **string** |  | 
**Message** | **string** |  | 
**PayerAccountId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 
**RunningHash** | **string** |  | 
**RunningHashVersion** | **int32** |  | 
**SequenceNumber** | **int64** |  | 
**TopicId** | **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | 

## Methods

### NewTopicMessage

`func NewTopicMessage(consensusTimestamp string, message string, payerAccountId NullableString, runningHash string, runningHashVersion int32, sequenceNumber int64, topicId NullableString, ) *TopicMessage`

NewTopicMessage instantiates a new TopicMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicMessageWithDefaults

`func NewTopicMessageWithDefaults() *TopicMessage`

NewTopicMessageWithDefaults instantiates a new TopicMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunkInfo

`func (o *TopicMessage) GetChunkInfo() ChunkInfo`

GetChunkInfo returns the ChunkInfo field if non-nil, zero value otherwise.

### GetChunkInfoOk

`func (o *TopicMessage) GetChunkInfoOk() (*ChunkInfo, bool)`

GetChunkInfoOk returns a tuple with the ChunkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkInfo

`func (o *TopicMessage) SetChunkInfo(v ChunkInfo)`

SetChunkInfo sets ChunkInfo field to given value.

### HasChunkInfo

`func (o *TopicMessage) HasChunkInfo() bool`

HasChunkInfo returns a boolean if a field has been set.

### SetChunkInfoNil

`func (o *TopicMessage) SetChunkInfoNil(b bool)`

 SetChunkInfoNil sets the value for ChunkInfo to be an explicit nil

### UnsetChunkInfo
`func (o *TopicMessage) UnsetChunkInfo()`

UnsetChunkInfo ensures that no value is present for ChunkInfo, not even an explicit nil
### GetConsensusTimestamp

`func (o *TopicMessage) GetConsensusTimestamp() string`

GetConsensusTimestamp returns the ConsensusTimestamp field if non-nil, zero value otherwise.

### GetConsensusTimestampOk

`func (o *TopicMessage) GetConsensusTimestampOk() (*string, bool)`

GetConsensusTimestampOk returns a tuple with the ConsensusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusTimestamp

`func (o *TopicMessage) SetConsensusTimestamp(v string)`

SetConsensusTimestamp sets ConsensusTimestamp field to given value.


### GetMessage

`func (o *TopicMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TopicMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TopicMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPayerAccountId

`func (o *TopicMessage) GetPayerAccountId() string`

GetPayerAccountId returns the PayerAccountId field if non-nil, zero value otherwise.

### GetPayerAccountIdOk

`func (o *TopicMessage) GetPayerAccountIdOk() (*string, bool)`

GetPayerAccountIdOk returns a tuple with the PayerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAccountId

`func (o *TopicMessage) SetPayerAccountId(v string)`

SetPayerAccountId sets PayerAccountId field to given value.


### SetPayerAccountIdNil

`func (o *TopicMessage) SetPayerAccountIdNil(b bool)`

 SetPayerAccountIdNil sets the value for PayerAccountId to be an explicit nil

### UnsetPayerAccountId
`func (o *TopicMessage) UnsetPayerAccountId()`

UnsetPayerAccountId ensures that no value is present for PayerAccountId, not even an explicit nil
### GetRunningHash

`func (o *TopicMessage) GetRunningHash() string`

GetRunningHash returns the RunningHash field if non-nil, zero value otherwise.

### GetRunningHashOk

`func (o *TopicMessage) GetRunningHashOk() (*string, bool)`

GetRunningHashOk returns a tuple with the RunningHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningHash

`func (o *TopicMessage) SetRunningHash(v string)`

SetRunningHash sets RunningHash field to given value.


### GetRunningHashVersion

`func (o *TopicMessage) GetRunningHashVersion() int32`

GetRunningHashVersion returns the RunningHashVersion field if non-nil, zero value otherwise.

### GetRunningHashVersionOk

`func (o *TopicMessage) GetRunningHashVersionOk() (*int32, bool)`

GetRunningHashVersionOk returns a tuple with the RunningHashVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningHashVersion

`func (o *TopicMessage) SetRunningHashVersion(v int32)`

SetRunningHashVersion sets RunningHashVersion field to given value.


### GetSequenceNumber

`func (o *TopicMessage) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *TopicMessage) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *TopicMessage) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetTopicId

`func (o *TopicMessage) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *TopicMessage) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *TopicMessage) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.


### SetTopicIdNil

`func (o *TopicMessage) SetTopicIdNil(b bool)`

 SetTopicIdNil sets the value for TopicId to be an explicit nil

### UnsetTopicId
`func (o *TopicMessage) UnsetTopicId()`

UnsetTopicId ensures that no value is present for TopicId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


