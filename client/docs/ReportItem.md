# ReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Contains a JSON pointer value (as defined in IETF RFC 6901) that references a  location of a resource to which the modification is subject.  | 
**Reason** | Pointer to **string** | A human-readable reason providing details on the reported modification failure.  The reason string should identify the operation that failed using the operation&#39;s  array index to assist in correlation of the invalid parameter with the failed  operation, e.g. \&quot;Replacement value invalid for attribute (failed operation index&#x3D; 4)\&quot;.  | [optional] 

## Methods

### NewReportItem

`func NewReportItem(path string, ) *ReportItem`

NewReportItem instantiates a new ReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportItemWithDefaults

`func NewReportItemWithDefaults() *ReportItem`

NewReportItemWithDefaults instantiates a new ReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ReportItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReportItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReportItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetReason

`func (o *ReportItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReportItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReportItem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReportItem) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


