# A2xSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NrA2xServicesAuth** | Pointer to [**NrA2xAuth**](NrA2xAuth.md) |  | [optional] 
**LteA2xServicesAuth** | Pointer to [**LteA2xAuth**](LteA2xAuth.md) |  | [optional] 
**NrUePc5Ambr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**LtePc5Ambr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewA2xSubscriptionData

`func NewA2xSubscriptionData() *A2xSubscriptionData`

NewA2xSubscriptionData instantiates a new A2xSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewA2xSubscriptionDataWithDefaults

`func NewA2xSubscriptionDataWithDefaults() *A2xSubscriptionData`

NewA2xSubscriptionDataWithDefaults instantiates a new A2xSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNrA2xServicesAuth

`func (o *A2xSubscriptionData) GetNrA2xServicesAuth() NrA2xAuth`

GetNrA2xServicesAuth returns the NrA2xServicesAuth field if non-nil, zero value otherwise.

### GetNrA2xServicesAuthOk

`func (o *A2xSubscriptionData) GetNrA2xServicesAuthOk() (*NrA2xAuth, bool)`

GetNrA2xServicesAuthOk returns a tuple with the NrA2xServicesAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrA2xServicesAuth

`func (o *A2xSubscriptionData) SetNrA2xServicesAuth(v NrA2xAuth)`

SetNrA2xServicesAuth sets NrA2xServicesAuth field to given value.

### HasNrA2xServicesAuth

`func (o *A2xSubscriptionData) HasNrA2xServicesAuth() bool`

HasNrA2xServicesAuth returns a boolean if a field has been set.

### GetLteA2xServicesAuth

`func (o *A2xSubscriptionData) GetLteA2xServicesAuth() LteA2xAuth`

GetLteA2xServicesAuth returns the LteA2xServicesAuth field if non-nil, zero value otherwise.

### GetLteA2xServicesAuthOk

`func (o *A2xSubscriptionData) GetLteA2xServicesAuthOk() (*LteA2xAuth, bool)`

GetLteA2xServicesAuthOk returns a tuple with the LteA2xServicesAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteA2xServicesAuth

`func (o *A2xSubscriptionData) SetLteA2xServicesAuth(v LteA2xAuth)`

SetLteA2xServicesAuth sets LteA2xServicesAuth field to given value.

### HasLteA2xServicesAuth

`func (o *A2xSubscriptionData) HasLteA2xServicesAuth() bool`

HasLteA2xServicesAuth returns a boolean if a field has been set.

### GetNrUePc5Ambr

`func (o *A2xSubscriptionData) GetNrUePc5Ambr() string`

GetNrUePc5Ambr returns the NrUePc5Ambr field if non-nil, zero value otherwise.

### GetNrUePc5AmbrOk

`func (o *A2xSubscriptionData) GetNrUePc5AmbrOk() (*string, bool)`

GetNrUePc5AmbrOk returns a tuple with the NrUePc5Ambr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrUePc5Ambr

`func (o *A2xSubscriptionData) SetNrUePc5Ambr(v string)`

SetNrUePc5Ambr sets NrUePc5Ambr field to given value.

### HasNrUePc5Ambr

`func (o *A2xSubscriptionData) HasNrUePc5Ambr() bool`

HasNrUePc5Ambr returns a boolean if a field has been set.

### GetLtePc5Ambr

`func (o *A2xSubscriptionData) GetLtePc5Ambr() string`

GetLtePc5Ambr returns the LtePc5Ambr field if non-nil, zero value otherwise.

### GetLtePc5AmbrOk

`func (o *A2xSubscriptionData) GetLtePc5AmbrOk() (*string, bool)`

GetLtePc5AmbrOk returns a tuple with the LtePc5Ambr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtePc5Ambr

`func (o *A2xSubscriptionData) SetLtePc5Ambr(v string)`

SetLtePc5Ambr sets LtePc5Ambr field to given value.

### HasLtePc5Ambr

`func (o *A2xSubscriptionData) HasLtePc5Ambr() bool`

HasLtePc5Ambr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


