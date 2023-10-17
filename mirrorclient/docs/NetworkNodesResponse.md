# NetworkNodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]NetworkNode**](NetworkNode.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewNetworkNodesResponse

`func NewNetworkNodesResponse(nodes []NetworkNode, links Links, ) *NetworkNodesResponse`

NewNetworkNodesResponse instantiates a new NetworkNodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNodesResponseWithDefaults

`func NewNetworkNodesResponseWithDefaults() *NetworkNodesResponse`

NewNetworkNodesResponseWithDefaults instantiates a new NetworkNodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *NetworkNodesResponse) GetNodes() []NetworkNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NetworkNodesResponse) GetNodesOk() (*[]NetworkNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NetworkNodesResponse) SetNodes(v []NetworkNode)`

SetNodes sets Nodes field to given value.


### GetLinks

`func (o *NetworkNodesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkNodesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkNodesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


