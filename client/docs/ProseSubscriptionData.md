# ProseSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProseServiceAuth** | Pointer to [**ProseServiceAuth**](ProseServiceAuth.md) |  | [optional] 
**NrUePc5Ambr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ProseAllowedPlmn** | Pointer to [**[]ProSeAllowedPlmn**](ProSeAllowedPlmn.md) |  | [optional] 

## Methods

### NewProseSubscriptionData

`func NewProseSubscriptionData() *ProseSubscriptionData`

NewProseSubscriptionData instantiates a new ProseSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseSubscriptionDataWithDefaults

`func NewProseSubscriptionDataWithDefaults() *ProseSubscriptionData`

NewProseSubscriptionDataWithDefaults instantiates a new ProseSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProseServiceAuth

`func (o *ProseSubscriptionData) GetProseServiceAuth() ProseServiceAuth`

GetProseServiceAuth returns the ProseServiceAuth field if non-nil, zero value otherwise.

### GetProseServiceAuthOk

`func (o *ProseSubscriptionData) GetProseServiceAuthOk() (*ProseServiceAuth, bool)`

GetProseServiceAuthOk returns a tuple with the ProseServiceAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseServiceAuth

`func (o *ProseSubscriptionData) SetProseServiceAuth(v ProseServiceAuth)`

SetProseServiceAuth sets ProseServiceAuth field to given value.

### HasProseServiceAuth

`func (o *ProseSubscriptionData) HasProseServiceAuth() bool`

HasProseServiceAuth returns a boolean if a field has been set.

### GetNrUePc5Ambr

`func (o *ProseSubscriptionData) GetNrUePc5Ambr() string`

GetNrUePc5Ambr returns the NrUePc5Ambr field if non-nil, zero value otherwise.

### GetNrUePc5AmbrOk

`func (o *ProseSubscriptionData) GetNrUePc5AmbrOk() (*string, bool)`

GetNrUePc5AmbrOk returns a tuple with the NrUePc5Ambr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrUePc5Ambr

`func (o *ProseSubscriptionData) SetNrUePc5Ambr(v string)`

SetNrUePc5Ambr sets NrUePc5Ambr field to given value.

### HasNrUePc5Ambr

`func (o *ProseSubscriptionData) HasNrUePc5Ambr() bool`

HasNrUePc5Ambr returns a boolean if a field has been set.

### GetProseAllowedPlmn

`func (o *ProseSubscriptionData) GetProseAllowedPlmn() []ProSeAllowedPlmn`

GetProseAllowedPlmn returns the ProseAllowedPlmn field if non-nil, zero value otherwise.

### GetProseAllowedPlmnOk

`func (o *ProseSubscriptionData) GetProseAllowedPlmnOk() (*[]ProSeAllowedPlmn, bool)`

GetProseAllowedPlmnOk returns a tuple with the ProseAllowedPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAllowedPlmn

`func (o *ProseSubscriptionData) SetProseAllowedPlmn(v []ProSeAllowedPlmn)`

SetProseAllowedPlmn sets ProseAllowedPlmn field to given value.

### HasProseAllowedPlmn

`func (o *ProseSubscriptionData) HasProseAllowedPlmn() bool`

HasProseAllowedPlmn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


