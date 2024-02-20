# AppSpecificExpectedUeBehaviourData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**TrafficFilters** | Pointer to [**[]FlowInfo**](FlowInfo.md) |  | [optional] 
**ExpectedInactivityTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ConfidenceLevel** | Pointer to **string** |  | [optional] 
**AccuracyLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSpecificExpectedUeBehaviourData

`func NewAppSpecificExpectedUeBehaviourData() *AppSpecificExpectedUeBehaviourData`

NewAppSpecificExpectedUeBehaviourData instantiates a new AppSpecificExpectedUeBehaviourData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSpecificExpectedUeBehaviourDataWithDefaults

`func NewAppSpecificExpectedUeBehaviourDataWithDefaults() *AppSpecificExpectedUeBehaviourData`

NewAppSpecificExpectedUeBehaviourDataWithDefaults instantiates a new AppSpecificExpectedUeBehaviourData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppSpecificExpectedUeBehaviourData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppSpecificExpectedUeBehaviourData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppSpecificExpectedUeBehaviourData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppSpecificExpectedUeBehaviourData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetTrafficFilters

`func (o *AppSpecificExpectedUeBehaviourData) GetTrafficFilters() []FlowInfo`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *AppSpecificExpectedUeBehaviourData) GetTrafficFiltersOk() (*[]FlowInfo, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *AppSpecificExpectedUeBehaviourData) SetTrafficFilters(v []FlowInfo)`

SetTrafficFilters sets TrafficFilters field to given value.

### HasTrafficFilters

`func (o *AppSpecificExpectedUeBehaviourData) HasTrafficFilters() bool`

HasTrafficFilters returns a boolean if a field has been set.

### GetExpectedInactivityTime

`func (o *AppSpecificExpectedUeBehaviourData) GetExpectedInactivityTime() int32`

GetExpectedInactivityTime returns the ExpectedInactivityTime field if non-nil, zero value otherwise.

### GetExpectedInactivityTimeOk

`func (o *AppSpecificExpectedUeBehaviourData) GetExpectedInactivityTimeOk() (*int32, bool)`

GetExpectedInactivityTimeOk returns a tuple with the ExpectedInactivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedInactivityTime

`func (o *AppSpecificExpectedUeBehaviourData) SetExpectedInactivityTime(v int32)`

SetExpectedInactivityTime sets ExpectedInactivityTime field to given value.

### HasExpectedInactivityTime

`func (o *AppSpecificExpectedUeBehaviourData) HasExpectedInactivityTime() bool`

HasExpectedInactivityTime returns a boolean if a field has been set.

### GetValidityTime

`func (o *AppSpecificExpectedUeBehaviourData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AppSpecificExpectedUeBehaviourData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AppSpecificExpectedUeBehaviourData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *AppSpecificExpectedUeBehaviourData) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetConfidenceLevel

`func (o *AppSpecificExpectedUeBehaviourData) GetConfidenceLevel() string`

GetConfidenceLevel returns the ConfidenceLevel field if non-nil, zero value otherwise.

### GetConfidenceLevelOk

`func (o *AppSpecificExpectedUeBehaviourData) GetConfidenceLevelOk() (*string, bool)`

GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceLevel

`func (o *AppSpecificExpectedUeBehaviourData) SetConfidenceLevel(v string)`

SetConfidenceLevel sets ConfidenceLevel field to given value.

### HasConfidenceLevel

`func (o *AppSpecificExpectedUeBehaviourData) HasConfidenceLevel() bool`

HasConfidenceLevel returns a boolean if a field has been set.

### GetAccuracyLevel

`func (o *AppSpecificExpectedUeBehaviourData) GetAccuracyLevel() string`

GetAccuracyLevel returns the AccuracyLevel field if non-nil, zero value otherwise.

### GetAccuracyLevelOk

`func (o *AppSpecificExpectedUeBehaviourData) GetAccuracyLevelOk() (*string, bool)`

GetAccuracyLevelOk returns a tuple with the AccuracyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyLevel

`func (o *AppSpecificExpectedUeBehaviourData) SetAccuracyLevel(v string)`

SetAccuracyLevel sets AccuracyLevel field to given value.

### HasAccuracyLevel

`func (o *AppSpecificExpectedUeBehaviourData) HasAccuracyLevel() bool`

HasAccuracyLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


