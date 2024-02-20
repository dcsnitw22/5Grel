# SpatialValidityCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingAreaList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**Countries** | Pointer to **[]string** |  | [optional] 
**GeographicalServiceArea** | Pointer to [**GeoServiceArea**](GeoServiceArea.md) |  | [optional] 

## Methods

### NewSpatialValidityCond

`func NewSpatialValidityCond() *SpatialValidityCond`

NewSpatialValidityCond instantiates a new SpatialValidityCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpatialValidityCondWithDefaults

`func NewSpatialValidityCondWithDefaults() *SpatialValidityCond`

NewSpatialValidityCondWithDefaults instantiates a new SpatialValidityCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingAreaList

`func (o *SpatialValidityCond) GetTrackingAreaList() []Tai`

GetTrackingAreaList returns the TrackingAreaList field if non-nil, zero value otherwise.

### GetTrackingAreaListOk

`func (o *SpatialValidityCond) GetTrackingAreaListOk() (*[]Tai, bool)`

GetTrackingAreaListOk returns a tuple with the TrackingAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingAreaList

`func (o *SpatialValidityCond) SetTrackingAreaList(v []Tai)`

SetTrackingAreaList sets TrackingAreaList field to given value.

### HasTrackingAreaList

`func (o *SpatialValidityCond) HasTrackingAreaList() bool`

HasTrackingAreaList returns a boolean if a field has been set.

### GetCountries

`func (o *SpatialValidityCond) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *SpatialValidityCond) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *SpatialValidityCond) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *SpatialValidityCond) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetGeographicalServiceArea

`func (o *SpatialValidityCond) GetGeographicalServiceArea() GeoServiceArea`

GetGeographicalServiceArea returns the GeographicalServiceArea field if non-nil, zero value otherwise.

### GetGeographicalServiceAreaOk

`func (o *SpatialValidityCond) GetGeographicalServiceAreaOk() (*GeoServiceArea, bool)`

GetGeographicalServiceAreaOk returns a tuple with the GeographicalServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalServiceArea

`func (o *SpatialValidityCond) SetGeographicalServiceArea(v GeoServiceArea)`

SetGeographicalServiceArea sets GeographicalServiceArea field to given value.

### HasGeographicalServiceArea

`func (o *SpatialValidityCond) HasGeographicalServiceArea() bool`

HasGeographicalServiceArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


