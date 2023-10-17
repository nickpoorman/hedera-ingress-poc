# BlockTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 
**To** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 

## Methods

### NewBlockTimestamp

`func NewBlockTimestamp() *BlockTimestamp`

NewBlockTimestamp instantiates a new BlockTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockTimestampWithDefaults

`func NewBlockTimestampWithDefaults() *BlockTimestamp`

NewBlockTimestampWithDefaults instantiates a new BlockTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BlockTimestamp) GetFrom() TimestampRange`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BlockTimestamp) GetFromOk() (*TimestampRange, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BlockTimestamp) SetFrom(v TimestampRange)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BlockTimestamp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *BlockTimestamp) GetTo() TimestampRange`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BlockTimestamp) GetToOk() (*TimestampRange, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BlockTimestamp) SetTo(v TimestampRange)`

SetTo sets To field to given value.

### HasTo

`func (o *BlockTimestamp) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


