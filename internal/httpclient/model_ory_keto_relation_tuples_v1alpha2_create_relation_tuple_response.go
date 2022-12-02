/*
 * Ory Keto API
 *
 * Documentation for all of Ory Keto's REST APIs. gRPC is documented separately.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse The response from creating a new relationship.
type OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse struct {
	RelationTuple *Relationship `json:"relationTuple,omitempty"`
}

// NewOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse instantiates a new OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse() *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse {
	this := OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse{}
	return &this
}

// NewOryKetoRelationTuplesV1alpha2CreateRelationTupleResponseWithDefaults instantiates a new OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOryKetoRelationTuplesV1alpha2CreateRelationTupleResponseWithDefaults() *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse {
	this := OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse{}
	return &this
}

// GetRelationTuple returns the RelationTuple field value if set, zero value otherwise.
func (o *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) GetRelationTuple() Relationship {
	if o == nil || o.RelationTuple == nil {
		var ret Relationship
		return ret
	}
	return *o.RelationTuple
}

// GetRelationTupleOk returns a tuple with the RelationTuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) GetRelationTupleOk() (*Relationship, bool) {
	if o == nil || o.RelationTuple == nil {
		return nil, false
	}
	return o.RelationTuple, true
}

// HasRelationTuple returns a boolean if a field has been set.
func (o *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) HasRelationTuple() bool {
	if o != nil && o.RelationTuple != nil {
		return true
	}

	return false
}

// SetRelationTuple gets a reference to the given Relationship and assigns it to the RelationTuple field.
func (o *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) SetRelationTuple(v Relationship) {
	o.RelationTuple = &v
}

func (o OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RelationTuple != nil {
		toSerialize["relationTuple"] = o.RelationTuple
	}
	return json.Marshal(toSerialize)
}

type NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse struct {
	value *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse
	isSet bool
}

func (v NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) Get() *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse {
	return v.value
}

func (v *NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) Set(val *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse(val *OryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) *NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse {
	return &NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse{value: val, isSet: true}
}

func (v NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOryKetoRelationTuplesV1alpha2CreateRelationTupleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
