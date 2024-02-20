# Tmgi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsServiceId** | **string** | MBS Service ID | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 

## Methods

### NewTmgi

`func NewTmgi(mbsServiceId string, plmnId PlmnId, ) *Tmgi`

NewTmgi instantiates a new Tmgi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmgiWithDefaults

`func NewTmgiWithDefaults() *Tmgi`

NewTmgiWithDefaults instantiates a new Tmgi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsServiceId

`func (o *Tmgi) GetMbsServiceId() string`

GetMbsServiceId returns the MbsServiceId field if non-nil, zero value otherwise.

### GetMbsServiceIdOk

`func (o *Tmgi) GetMbsServiceIdOk() (*string, bool)`

GetMbsServiceIdOk returns a tuple with the MbsServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceId

`func (o *Tmgi) SetMbsServiceId(v string)`

SetMbsServiceId sets MbsServiceId field to given value.


### GetPlmnId

`func (o *Tmgi) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Tmgi) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Tmgi) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


