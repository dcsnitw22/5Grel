# RangingSlPosQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HAccuracy** | Pointer to **float32** | Indicates value of accuracy. | [optional] 
**VAccuracy** | Pointer to **float32** | Indicates value of accuracy. | [optional] 
**RelativeHAccuracy** | Pointer to **float32** | Indicates value of accuracy. | [optional] 
**RelativeVAccuracy** | Pointer to **float32** | Indicates value of accuracy. | [optional] 
**DistanceAccuracy** | Pointer to **float32** | Indicates value of accuracy. | [optional] 
**DirectionAccuracy** | Pointer to **float32** | Indicates value of accuracy. | [optional] 
**VerticalRequested** | Pointer to **bool** |  | [optional] 
**ResponseTime** | Pointer to [**ResponseTime**](ResponseTime.md) |  | [optional] 
**LcsQosClass** | Pointer to [**LcsQosClass**](LcsQosClass.md) |  | [optional] 

## Methods

### NewRangingSlPosQos

`func NewRangingSlPosQos() *RangingSlPosQos`

NewRangingSlPosQos instantiates a new RangingSlPosQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangingSlPosQosWithDefaults

`func NewRangingSlPosQosWithDefaults() *RangingSlPosQos`

NewRangingSlPosQosWithDefaults instantiates a new RangingSlPosQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHAccuracy

`func (o *RangingSlPosQos) GetHAccuracy() float32`

GetHAccuracy returns the HAccuracy field if non-nil, zero value otherwise.

### GetHAccuracyOk

`func (o *RangingSlPosQos) GetHAccuracyOk() (*float32, bool)`

GetHAccuracyOk returns a tuple with the HAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHAccuracy

`func (o *RangingSlPosQos) SetHAccuracy(v float32)`

SetHAccuracy sets HAccuracy field to given value.

### HasHAccuracy

`func (o *RangingSlPosQos) HasHAccuracy() bool`

HasHAccuracy returns a boolean if a field has been set.

### GetVAccuracy

`func (o *RangingSlPosQos) GetVAccuracy() float32`

GetVAccuracy returns the VAccuracy field if non-nil, zero value otherwise.

### GetVAccuracyOk

`func (o *RangingSlPosQos) GetVAccuracyOk() (*float32, bool)`

GetVAccuracyOk returns a tuple with the VAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAccuracy

`func (o *RangingSlPosQos) SetVAccuracy(v float32)`

SetVAccuracy sets VAccuracy field to given value.

### HasVAccuracy

`func (o *RangingSlPosQos) HasVAccuracy() bool`

HasVAccuracy returns a boolean if a field has been set.

### GetRelativeHAccuracy

`func (o *RangingSlPosQos) GetRelativeHAccuracy() float32`

GetRelativeHAccuracy returns the RelativeHAccuracy field if non-nil, zero value otherwise.

### GetRelativeHAccuracyOk

`func (o *RangingSlPosQos) GetRelativeHAccuracyOk() (*float32, bool)`

GetRelativeHAccuracyOk returns a tuple with the RelativeHAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeHAccuracy

`func (o *RangingSlPosQos) SetRelativeHAccuracy(v float32)`

SetRelativeHAccuracy sets RelativeHAccuracy field to given value.

### HasRelativeHAccuracy

`func (o *RangingSlPosQos) HasRelativeHAccuracy() bool`

HasRelativeHAccuracy returns a boolean if a field has been set.

### GetRelativeVAccuracy

`func (o *RangingSlPosQos) GetRelativeVAccuracy() float32`

GetRelativeVAccuracy returns the RelativeVAccuracy field if non-nil, zero value otherwise.

### GetRelativeVAccuracyOk

`func (o *RangingSlPosQos) GetRelativeVAccuracyOk() (*float32, bool)`

GetRelativeVAccuracyOk returns a tuple with the RelativeVAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeVAccuracy

`func (o *RangingSlPosQos) SetRelativeVAccuracy(v float32)`

SetRelativeVAccuracy sets RelativeVAccuracy field to given value.

### HasRelativeVAccuracy

`func (o *RangingSlPosQos) HasRelativeVAccuracy() bool`

HasRelativeVAccuracy returns a boolean if a field has been set.

### GetDistanceAccuracy

`func (o *RangingSlPosQos) GetDistanceAccuracy() float32`

GetDistanceAccuracy returns the DistanceAccuracy field if non-nil, zero value otherwise.

### GetDistanceAccuracyOk

`func (o *RangingSlPosQos) GetDistanceAccuracyOk() (*float32, bool)`

GetDistanceAccuracyOk returns a tuple with the DistanceAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceAccuracy

`func (o *RangingSlPosQos) SetDistanceAccuracy(v float32)`

SetDistanceAccuracy sets DistanceAccuracy field to given value.

### HasDistanceAccuracy

`func (o *RangingSlPosQos) HasDistanceAccuracy() bool`

HasDistanceAccuracy returns a boolean if a field has been set.

### GetDirectionAccuracy

`func (o *RangingSlPosQos) GetDirectionAccuracy() float32`

GetDirectionAccuracy returns the DirectionAccuracy field if non-nil, zero value otherwise.

### GetDirectionAccuracyOk

`func (o *RangingSlPosQos) GetDirectionAccuracyOk() (*float32, bool)`

GetDirectionAccuracyOk returns a tuple with the DirectionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionAccuracy

`func (o *RangingSlPosQos) SetDirectionAccuracy(v float32)`

SetDirectionAccuracy sets DirectionAccuracy field to given value.

### HasDirectionAccuracy

`func (o *RangingSlPosQos) HasDirectionAccuracy() bool`

HasDirectionAccuracy returns a boolean if a field has been set.

### GetVerticalRequested

`func (o *RangingSlPosQos) GetVerticalRequested() bool`

GetVerticalRequested returns the VerticalRequested field if non-nil, zero value otherwise.

### GetVerticalRequestedOk

`func (o *RangingSlPosQos) GetVerticalRequestedOk() (*bool, bool)`

GetVerticalRequestedOk returns a tuple with the VerticalRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalRequested

`func (o *RangingSlPosQos) SetVerticalRequested(v bool)`

SetVerticalRequested sets VerticalRequested field to given value.

### HasVerticalRequested

`func (o *RangingSlPosQos) HasVerticalRequested() bool`

HasVerticalRequested returns a boolean if a field has been set.

### GetResponseTime

`func (o *RangingSlPosQos) GetResponseTime() ResponseTime`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *RangingSlPosQos) GetResponseTimeOk() (*ResponseTime, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *RangingSlPosQos) SetResponseTime(v ResponseTime)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *RangingSlPosQos) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetLcsQosClass

`func (o *RangingSlPosQos) GetLcsQosClass() LcsQosClass`

GetLcsQosClass returns the LcsQosClass field if non-nil, zero value otherwise.

### GetLcsQosClassOk

`func (o *RangingSlPosQos) GetLcsQosClassOk() (*LcsQosClass, bool)`

GetLcsQosClassOk returns a tuple with the LcsQosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsQosClass

`func (o *RangingSlPosQos) SetLcsQosClass(v LcsQosClass)`

SetLcsQosClass sets LcsQosClass field to given value.

### HasLcsQosClass

`func (o *RangingSlPosQos) HasLcsQosClass() bool`

HasLcsQosClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


