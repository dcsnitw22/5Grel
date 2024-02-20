# TimeSyncData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | **bool** |  | 
**UuTimeSyncErrBdgt** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 
**CoverageArea** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**ClockQualityDetailLevel** | Pointer to [**ClockQualityDetailLevel**](ClockQualityDetailLevel.md) |  | [optional] 
**ClockQualityAcceptanceCriteria** | Pointer to [**[]ClockQualityAcceptanceCriterion**](ClockQualityAcceptanceCriterion.md) |  | [optional] 

## Methods

### NewTimeSyncData

`func NewTimeSyncData(authorized bool, ) *TimeSyncData`

NewTimeSyncData instantiates a new TimeSyncData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncDataWithDefaults

`func NewTimeSyncDataWithDefaults() *TimeSyncData`

NewTimeSyncDataWithDefaults instantiates a new TimeSyncData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *TimeSyncData) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *TimeSyncData) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *TimeSyncData) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.


### GetUuTimeSyncErrBdgt

`func (o *TimeSyncData) GetUuTimeSyncErrBdgt() int32`

GetUuTimeSyncErrBdgt returns the UuTimeSyncErrBdgt field if non-nil, zero value otherwise.

### GetUuTimeSyncErrBdgtOk

`func (o *TimeSyncData) GetUuTimeSyncErrBdgtOk() (*int32, bool)`

GetUuTimeSyncErrBdgtOk returns a tuple with the UuTimeSyncErrBdgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuTimeSyncErrBdgt

`func (o *TimeSyncData) SetUuTimeSyncErrBdgt(v int32)`

SetUuTimeSyncErrBdgt sets UuTimeSyncErrBdgt field to given value.

### HasUuTimeSyncErrBdgt

`func (o *TimeSyncData) HasUuTimeSyncErrBdgt() bool`

HasUuTimeSyncErrBdgt returns a boolean if a field has been set.

### GetTempVals

`func (o *TimeSyncData) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *TimeSyncData) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *TimeSyncData) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *TimeSyncData) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.

### GetCoverageArea

`func (o *TimeSyncData) GetCoverageArea() []Tai`

GetCoverageArea returns the CoverageArea field if non-nil, zero value otherwise.

### GetCoverageAreaOk

`func (o *TimeSyncData) GetCoverageAreaOk() (*[]Tai, bool)`

GetCoverageAreaOk returns a tuple with the CoverageArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageArea

`func (o *TimeSyncData) SetCoverageArea(v []Tai)`

SetCoverageArea sets CoverageArea field to given value.

### HasCoverageArea

`func (o *TimeSyncData) HasCoverageArea() bool`

HasCoverageArea returns a boolean if a field has been set.

### GetClockQualityDetailLevel

`func (o *TimeSyncData) GetClockQualityDetailLevel() ClockQualityDetailLevel`

GetClockQualityDetailLevel returns the ClockQualityDetailLevel field if non-nil, zero value otherwise.

### GetClockQualityDetailLevelOk

`func (o *TimeSyncData) GetClockQualityDetailLevelOk() (*ClockQualityDetailLevel, bool)`

GetClockQualityDetailLevelOk returns a tuple with the ClockQualityDetailLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockQualityDetailLevel

`func (o *TimeSyncData) SetClockQualityDetailLevel(v ClockQualityDetailLevel)`

SetClockQualityDetailLevel sets ClockQualityDetailLevel field to given value.

### HasClockQualityDetailLevel

`func (o *TimeSyncData) HasClockQualityDetailLevel() bool`

HasClockQualityDetailLevel returns a boolean if a field has been set.

### GetClockQualityAcceptanceCriteria

`func (o *TimeSyncData) GetClockQualityAcceptanceCriteria() []ClockQualityAcceptanceCriterion`

GetClockQualityAcceptanceCriteria returns the ClockQualityAcceptanceCriteria field if non-nil, zero value otherwise.

### GetClockQualityAcceptanceCriteriaOk

`func (o *TimeSyncData) GetClockQualityAcceptanceCriteriaOk() (*[]ClockQualityAcceptanceCriterion, bool)`

GetClockQualityAcceptanceCriteriaOk returns a tuple with the ClockQualityAcceptanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockQualityAcceptanceCriteria

`func (o *TimeSyncData) SetClockQualityAcceptanceCriteria(v []ClockQualityAcceptanceCriterion)`

SetClockQualityAcceptanceCriteria sets ClockQualityAcceptanceCriteria field to given value.

### HasClockQualityAcceptanceCriteria

`func (o *TimeSyncData) HasClockQualityAcceptanceCriteria() bool`

HasClockQualityAcceptanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


