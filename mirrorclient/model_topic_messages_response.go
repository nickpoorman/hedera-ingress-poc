/*
Hedera Mirror Node REST API

The Mirror Node REST API offers the ability to query cryptocurrency transactions and account information from a Hedera managed mirror node.  Base url: [/api/v1](/api/v1)  OpenAPI Spec: [/api/v1/docs/openapi.yml](/api/v1/docs/openapi.yml)

API version: 0.89.0
Contact: mirrornode@hedera.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mirrorclient

import (
	"encoding/json"
)

// checks if the TopicMessagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopicMessagesResponse{}

// TopicMessagesResponse struct for TopicMessagesResponse
type TopicMessagesResponse struct {
	Messages []TopicMessage `json:"messages,omitempty"`
	Links *Links `json:"links,omitempty"`
}

// NewTopicMessagesResponse instantiates a new TopicMessagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicMessagesResponse() *TopicMessagesResponse {
	this := TopicMessagesResponse{}
	return &this
}

// NewTopicMessagesResponseWithDefaults instantiates a new TopicMessagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicMessagesResponseWithDefaults() *TopicMessagesResponse {
	this := TopicMessagesResponse{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *TopicMessagesResponse) GetMessages() []TopicMessage {
	if o == nil || IsNil(o.Messages) {
		var ret []TopicMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessagesResponse) GetMessagesOk() ([]TopicMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *TopicMessagesResponse) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []TopicMessage and assigns it to the Messages field.
func (o *TopicMessagesResponse) SetMessages(v []TopicMessage) {
	o.Messages = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TopicMessagesResponse) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicMessagesResponse) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TopicMessagesResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *TopicMessagesResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o TopicMessagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopicMessagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTopicMessagesResponse struct {
	value *TopicMessagesResponse
	isSet bool
}

func (v NullableTopicMessagesResponse) Get() *TopicMessagesResponse {
	return v.value
}

func (v *NullableTopicMessagesResponse) Set(val *TopicMessagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicMessagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicMessagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicMessagesResponse(val *TopicMessagesResponse) *NullableTopicMessagesResponse {
	return &NullableTopicMessagesResponse{value: val, isSet: true}
}

func (v NullableTopicMessagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicMessagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


