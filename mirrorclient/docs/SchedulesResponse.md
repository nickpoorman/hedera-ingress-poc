# SchedulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to [**[]Schedule**](Schedule.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewSchedulesResponse

`func NewSchedulesResponse() *SchedulesResponse`

NewSchedulesResponse instantiates a new SchedulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulesResponseWithDefaults

`func NewSchedulesResponseWithDefaults() *SchedulesResponse`

NewSchedulesResponseWithDefaults instantiates a new SchedulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *SchedulesResponse) GetSchedules() []Schedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *SchedulesResponse) GetSchedulesOk() (*[]Schedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *SchedulesResponse) SetSchedules(v []Schedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *SchedulesResponse) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetLinks

`func (o *SchedulesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SchedulesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SchedulesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SchedulesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


