# LcsPrivacyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lpi** | Pointer to [**Lpi**](Lpi.md) |  | [optional] 
**UnrelatedClass** | Pointer to [**UnrelatedClass**](UnrelatedClass.md) |  | [optional] 
**PlmnOperatorClasses** | Pointer to [**[]PlmnOperatorClass**](PlmnOperatorClass.md) |  | [optional] 
**EvtRptExpectedArea** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**AreaUsageInd** | Pointer to [**AreaUsageInd**](AreaUsageInd.md) |  | [optional] [default to INSIDE_REPORTING]
**UpLocRepIndAf** | Pointer to [**UpLocRepIndAf**](UpLocRepIndAf.md) |  | [optional] [default to USER_PLANE_REPORT_NOT_ALLOWED]

## Methods

### NewLcsPrivacyData

`func NewLcsPrivacyData() *LcsPrivacyData`

NewLcsPrivacyData instantiates a new LcsPrivacyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLcsPrivacyDataWithDefaults

`func NewLcsPrivacyDataWithDefaults() *LcsPrivacyData`

NewLcsPrivacyDataWithDefaults instantiates a new LcsPrivacyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLpi

`func (o *LcsPrivacyData) GetLpi() Lpi`

GetLpi returns the Lpi field if non-nil, zero value otherwise.

### GetLpiOk

`func (o *LcsPrivacyData) GetLpiOk() (*Lpi, bool)`

GetLpiOk returns a tuple with the Lpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpi

`func (o *LcsPrivacyData) SetLpi(v Lpi)`

SetLpi sets Lpi field to given value.

### HasLpi

`func (o *LcsPrivacyData) HasLpi() bool`

HasLpi returns a boolean if a field has been set.

### GetUnrelatedClass

`func (o *LcsPrivacyData) GetUnrelatedClass() UnrelatedClass`

GetUnrelatedClass returns the UnrelatedClass field if non-nil, zero value otherwise.

### GetUnrelatedClassOk

`func (o *LcsPrivacyData) GetUnrelatedClassOk() (*UnrelatedClass, bool)`

GetUnrelatedClassOk returns a tuple with the UnrelatedClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrelatedClass

`func (o *LcsPrivacyData) SetUnrelatedClass(v UnrelatedClass)`

SetUnrelatedClass sets UnrelatedClass field to given value.

### HasUnrelatedClass

`func (o *LcsPrivacyData) HasUnrelatedClass() bool`

HasUnrelatedClass returns a boolean if a field has been set.

### GetPlmnOperatorClasses

`func (o *LcsPrivacyData) GetPlmnOperatorClasses() []PlmnOperatorClass`

GetPlmnOperatorClasses returns the PlmnOperatorClasses field if non-nil, zero value otherwise.

### GetPlmnOperatorClassesOk

`func (o *LcsPrivacyData) GetPlmnOperatorClassesOk() (*[]PlmnOperatorClass, bool)`

GetPlmnOperatorClassesOk returns a tuple with the PlmnOperatorClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnOperatorClasses

`func (o *LcsPrivacyData) SetPlmnOperatorClasses(v []PlmnOperatorClass)`

SetPlmnOperatorClasses sets PlmnOperatorClasses field to given value.

### HasPlmnOperatorClasses

`func (o *LcsPrivacyData) HasPlmnOperatorClasses() bool`

HasPlmnOperatorClasses returns a boolean if a field has been set.

### GetEvtRptExpectedArea

`func (o *LcsPrivacyData) GetEvtRptExpectedArea() GeographicArea`

GetEvtRptExpectedArea returns the EvtRptExpectedArea field if non-nil, zero value otherwise.

### GetEvtRptExpectedAreaOk

`func (o *LcsPrivacyData) GetEvtRptExpectedAreaOk() (*GeographicArea, bool)`

GetEvtRptExpectedAreaOk returns a tuple with the EvtRptExpectedArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtRptExpectedArea

`func (o *LcsPrivacyData) SetEvtRptExpectedArea(v GeographicArea)`

SetEvtRptExpectedArea sets EvtRptExpectedArea field to given value.

### HasEvtRptExpectedArea

`func (o *LcsPrivacyData) HasEvtRptExpectedArea() bool`

HasEvtRptExpectedArea returns a boolean if a field has been set.

### GetAreaUsageInd

`func (o *LcsPrivacyData) GetAreaUsageInd() AreaUsageInd`

GetAreaUsageInd returns the AreaUsageInd field if non-nil, zero value otherwise.

### GetAreaUsageIndOk

`func (o *LcsPrivacyData) GetAreaUsageIndOk() (*AreaUsageInd, bool)`

GetAreaUsageIndOk returns a tuple with the AreaUsageInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaUsageInd

`func (o *LcsPrivacyData) SetAreaUsageInd(v AreaUsageInd)`

SetAreaUsageInd sets AreaUsageInd field to given value.

### HasAreaUsageInd

`func (o *LcsPrivacyData) HasAreaUsageInd() bool`

HasAreaUsageInd returns a boolean if a field has been set.

### GetUpLocRepIndAf

`func (o *LcsPrivacyData) GetUpLocRepIndAf() UpLocRepIndAf`

GetUpLocRepIndAf returns the UpLocRepIndAf field if non-nil, zero value otherwise.

### GetUpLocRepIndAfOk

`func (o *LcsPrivacyData) GetUpLocRepIndAfOk() (*UpLocRepIndAf, bool)`

GetUpLocRepIndAfOk returns a tuple with the UpLocRepIndAf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLocRepIndAf

`func (o *LcsPrivacyData) SetUpLocRepIndAf(v UpLocRepIndAf)`

SetUpLocRepIndAf sets UpLocRepIndAf field to given value.

### HasUpLocRepIndAf

`func (o *LcsPrivacyData) HasUpLocRepIndAf() bool`

HasUpLocRepIndAf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


