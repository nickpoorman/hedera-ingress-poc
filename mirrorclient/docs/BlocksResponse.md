# BlocksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | Pointer to [**[]Block**](Block.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewBlocksResponse

`func NewBlocksResponse() *BlocksResponse`

NewBlocksResponse instantiates a new BlocksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocksResponseWithDefaults

`func NewBlocksResponseWithDefaults() *BlocksResponse`

NewBlocksResponseWithDefaults instantiates a new BlocksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *BlocksResponse) GetBlocks() []Block`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BlocksResponse) GetBlocksOk() (*[]Block, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BlocksResponse) SetBlocks(v []Block)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BlocksResponse) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetLinks

`func (o *BlocksResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BlocksResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BlocksResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BlocksResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


