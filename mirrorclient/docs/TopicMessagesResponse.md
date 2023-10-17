# TopicMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]TopicMessage**](TopicMessage.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewTopicMessagesResponse

`func NewTopicMessagesResponse() *TopicMessagesResponse`

NewTopicMessagesResponse instantiates a new TopicMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicMessagesResponseWithDefaults

`func NewTopicMessagesResponseWithDefaults() *TopicMessagesResponse`

NewTopicMessagesResponseWithDefaults instantiates a new TopicMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *TopicMessagesResponse) GetMessages() []TopicMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *TopicMessagesResponse) GetMessagesOk() (*[]TopicMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *TopicMessagesResponse) SetMessages(v []TopicMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *TopicMessagesResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetLinks

`func (o *TopicMessagesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TopicMessagesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TopicMessagesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TopicMessagesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


