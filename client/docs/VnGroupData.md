# VnGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionTypes** | Pointer to [**PduSessionTypes**](PduSessionTypes.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AppDescriptors** | Pointer to [**[]AppDescriptor**](AppDescriptor.md) |  | [optional] 
**SecondaryAuth** | Pointer to **bool** |  | [optional] 
**DnAaaIpAddressAllocation** | Pointer to **bool** |  | [optional] 
**DnAaaAddress** | Pointer to [**NullableIpAddress**](IpAddress.md) |  | [optional] 
**AdditionalDnAaaAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**DnAaaFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewVnGroupData

`func NewVnGroupData() *VnGroupData`

NewVnGroupData instantiates a new VnGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnGroupDataWithDefaults

`func NewVnGroupDataWithDefaults() *VnGroupData`

NewVnGroupDataWithDefaults instantiates a new VnGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionTypes

`func (o *VnGroupData) GetPduSessionTypes() PduSessionTypes`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *VnGroupData) GetPduSessionTypesOk() (*PduSessionTypes, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *VnGroupData) SetPduSessionTypes(v PduSessionTypes)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *VnGroupData) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetDnn

`func (o *VnGroupData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *VnGroupData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *VnGroupData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *VnGroupData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSingleNssai

`func (o *VnGroupData) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *VnGroupData) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *VnGroupData) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *VnGroupData) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetAppDescriptors

`func (o *VnGroupData) GetAppDescriptors() []AppDescriptor`

GetAppDescriptors returns the AppDescriptors field if non-nil, zero value otherwise.

### GetAppDescriptorsOk

`func (o *VnGroupData) GetAppDescriptorsOk() (*[]AppDescriptor, bool)`

GetAppDescriptorsOk returns a tuple with the AppDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptors

`func (o *VnGroupData) SetAppDescriptors(v []AppDescriptor)`

SetAppDescriptors sets AppDescriptors field to given value.

### HasAppDescriptors

`func (o *VnGroupData) HasAppDescriptors() bool`

HasAppDescriptors returns a boolean if a field has been set.

### GetSecondaryAuth

`func (o *VnGroupData) GetSecondaryAuth() bool`

GetSecondaryAuth returns the SecondaryAuth field if non-nil, zero value otherwise.

### GetSecondaryAuthOk

`func (o *VnGroupData) GetSecondaryAuthOk() (*bool, bool)`

GetSecondaryAuthOk returns a tuple with the SecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuth

`func (o *VnGroupData) SetSecondaryAuth(v bool)`

SetSecondaryAuth sets SecondaryAuth field to given value.

### HasSecondaryAuth

`func (o *VnGroupData) HasSecondaryAuth() bool`

HasSecondaryAuth returns a boolean if a field has been set.

### GetDnAaaIpAddressAllocation

`func (o *VnGroupData) GetDnAaaIpAddressAllocation() bool`

GetDnAaaIpAddressAllocation returns the DnAaaIpAddressAllocation field if non-nil, zero value otherwise.

### GetDnAaaIpAddressAllocationOk

`func (o *VnGroupData) GetDnAaaIpAddressAllocationOk() (*bool, bool)`

GetDnAaaIpAddressAllocationOk returns a tuple with the DnAaaIpAddressAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaIpAddressAllocation

`func (o *VnGroupData) SetDnAaaIpAddressAllocation(v bool)`

SetDnAaaIpAddressAllocation sets DnAaaIpAddressAllocation field to given value.

### HasDnAaaIpAddressAllocation

`func (o *VnGroupData) HasDnAaaIpAddressAllocation() bool`

HasDnAaaIpAddressAllocation returns a boolean if a field has been set.

### GetDnAaaAddress

`func (o *VnGroupData) GetDnAaaAddress() IpAddress`

GetDnAaaAddress returns the DnAaaAddress field if non-nil, zero value otherwise.

### GetDnAaaAddressOk

`func (o *VnGroupData) GetDnAaaAddressOk() (*IpAddress, bool)`

GetDnAaaAddressOk returns a tuple with the DnAaaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaAddress

`func (o *VnGroupData) SetDnAaaAddress(v IpAddress)`

SetDnAaaAddress sets DnAaaAddress field to given value.

### HasDnAaaAddress

`func (o *VnGroupData) HasDnAaaAddress() bool`

HasDnAaaAddress returns a boolean if a field has been set.

### SetDnAaaAddressNil

`func (o *VnGroupData) SetDnAaaAddressNil(b bool)`

 SetDnAaaAddressNil sets the value for DnAaaAddress to be an explicit nil

### UnsetDnAaaAddress
`func (o *VnGroupData) UnsetDnAaaAddress()`

UnsetDnAaaAddress ensures that no value is present for DnAaaAddress, not even an explicit nil
### GetAdditionalDnAaaAddresses

`func (o *VnGroupData) GetAdditionalDnAaaAddresses() []IpAddress`

GetAdditionalDnAaaAddresses returns the AdditionalDnAaaAddresses field if non-nil, zero value otherwise.

### GetAdditionalDnAaaAddressesOk

`func (o *VnGroupData) GetAdditionalDnAaaAddressesOk() (*[]IpAddress, bool)`

GetAdditionalDnAaaAddressesOk returns a tuple with the AdditionalDnAaaAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDnAaaAddresses

`func (o *VnGroupData) SetAdditionalDnAaaAddresses(v []IpAddress)`

SetAdditionalDnAaaAddresses sets AdditionalDnAaaAddresses field to given value.

### HasAdditionalDnAaaAddresses

`func (o *VnGroupData) HasAdditionalDnAaaAddresses() bool`

HasAdditionalDnAaaAddresses returns a boolean if a field has been set.

### GetDnAaaFqdn

`func (o *VnGroupData) GetDnAaaFqdn() string`

GetDnAaaFqdn returns the DnAaaFqdn field if non-nil, zero value otherwise.

### GetDnAaaFqdnOk

`func (o *VnGroupData) GetDnAaaFqdnOk() (*string, bool)`

GetDnAaaFqdnOk returns a tuple with the DnAaaFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnAaaFqdn

`func (o *VnGroupData) SetDnAaaFqdn(v string)`

SetDnAaaFqdn sets DnAaaFqdn field to given value.

### HasDnAaaFqdn

`func (o *VnGroupData) HasDnAaaFqdn() bool`

HasDnAaaFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


