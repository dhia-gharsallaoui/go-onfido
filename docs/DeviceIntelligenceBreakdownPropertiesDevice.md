# DeviceIntelligenceBreakdownPropertiesDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdkVersion** | Pointer to **string** | The SDK version that was used. | [optional] 
**SdkSource** | Pointer to **string** | The SDK used to upload the media. | [optional] 
**AuthenticationType** | Pointer to **string** | The token used to authenticate the request. | [optional] 
**RawModel** | Pointer to **string** | The model as set by the phone manufacturer (for Android and iOS) or the browser manufacturer (for Web). The model can be presented in name or number form depending on each manufacturer implementation. | [optional] 
**Os** | Pointer to **string** | The operating system of the device. The value came from manufacturer implementation (for Android and iOS) or browser&#39;s user agent (for Web). | [optional] 
**Browser** | Pointer to **string** | The browser name reported by the browser&#39;s user agent. | [optional] 
**Emulator** | Pointer to **bool** | Whether the device is an emulator. | [optional] 
**RandomizedDevice** | Pointer to **bool** | Whether the device is providing false randomized device and network information. | [optional] 
**FakeNetworkRequest** | Pointer to **bool** | Whether device is using stolen security tokens to send the network information. | [optional] 
**TrueOs** | Pointer to **string** | The true operating system of the device. | [optional] 
**OsAnomaly** | Pointer to **string** | The likelihood of an operating system anomaly between the true OS and the OS sent by the device. | [optional] 
**Rooted** | Pointer to **bool** | Whether the device is rooted. | [optional] 
**RemoteSoftware** | Pointer to **bool** | Whether the device is controlled via remote software. | [optional] 
**IpReputation** | Pointer to **string** | Whether there is highly suspicious traffic related to the IP address. The risk depends on the overall ratio of clear checks on a given IP. | [optional] 
**DeviceFingerprintReuse** | Pointer to **int32** | The number of times the device was used to create a report for a new applicant. A value greater than 1 indicates potential device reuse. | [optional] 
**SingleDeviceUsed** | Pointer to **NullableBool** | Whether the document or biometric media were uploaded from a single device. | [optional] 
**DocumentCapture** | Pointer to **string** | Whether the document media were live captured from the device camera. | [optional] 
**BiometricCapture** | Pointer to **string** | Whether the biometric media were live captured from the device camera. | [optional] 

## Methods

### NewDeviceIntelligenceBreakdownPropertiesDevice

`func NewDeviceIntelligenceBreakdownPropertiesDevice() *DeviceIntelligenceBreakdownPropertiesDevice`

NewDeviceIntelligenceBreakdownPropertiesDevice instantiates a new DeviceIntelligenceBreakdownPropertiesDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIntelligenceBreakdownPropertiesDeviceWithDefaults

`func NewDeviceIntelligenceBreakdownPropertiesDeviceWithDefaults() *DeviceIntelligenceBreakdownPropertiesDevice`

NewDeviceIntelligenceBreakdownPropertiesDeviceWithDefaults instantiates a new DeviceIntelligenceBreakdownPropertiesDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdkVersion

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetSdkSource

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetSdkSource() string`

GetSdkSource returns the SdkSource field if non-nil, zero value otherwise.

### GetSdkSourceOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetSdkSourceOk() (*string, bool)`

GetSdkSourceOk returns a tuple with the SdkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkSource

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetSdkSource(v string)`

SetSdkSource sets SdkSource field to given value.

### HasSdkSource

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasSdkSource() bool`

HasSdkSource returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetRawModel

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRawModel() string`

GetRawModel returns the RawModel field if non-nil, zero value otherwise.

### GetRawModelOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRawModelOk() (*string, bool)`

GetRawModelOk returns a tuple with the RawModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawModel

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetRawModel(v string)`

SetRawModel sets RawModel field to given value.

### HasRawModel

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasRawModel() bool`

HasRawModel returns a boolean if a field has been set.

### GetOs

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetBrowser

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetEmulator

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetEmulator() bool`

GetEmulator returns the Emulator field if non-nil, zero value otherwise.

### GetEmulatorOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetEmulatorOk() (*bool, bool)`

GetEmulatorOk returns a tuple with the Emulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulator

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetEmulator(v bool)`

SetEmulator sets Emulator field to given value.

### HasEmulator

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasEmulator() bool`

HasEmulator returns a boolean if a field has been set.

### GetRandomizedDevice

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRandomizedDevice() bool`

GetRandomizedDevice returns the RandomizedDevice field if non-nil, zero value otherwise.

### GetRandomizedDeviceOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRandomizedDeviceOk() (*bool, bool)`

GetRandomizedDeviceOk returns a tuple with the RandomizedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizedDevice

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetRandomizedDevice(v bool)`

SetRandomizedDevice sets RandomizedDevice field to given value.

### HasRandomizedDevice

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasRandomizedDevice() bool`

HasRandomizedDevice returns a boolean if a field has been set.

### GetFakeNetworkRequest

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetFakeNetworkRequest() bool`

GetFakeNetworkRequest returns the FakeNetworkRequest field if non-nil, zero value otherwise.

### GetFakeNetworkRequestOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetFakeNetworkRequestOk() (*bool, bool)`

GetFakeNetworkRequestOk returns a tuple with the FakeNetworkRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFakeNetworkRequest

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetFakeNetworkRequest(v bool)`

SetFakeNetworkRequest sets FakeNetworkRequest field to given value.

### HasFakeNetworkRequest

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasFakeNetworkRequest() bool`

HasFakeNetworkRequest returns a boolean if a field has been set.

### GetTrueOs

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetTrueOs() string`

GetTrueOs returns the TrueOs field if non-nil, zero value otherwise.

### GetTrueOsOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetTrueOsOk() (*string, bool)`

GetTrueOsOk returns a tuple with the TrueOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueOs

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetTrueOs(v string)`

SetTrueOs sets TrueOs field to given value.

### HasTrueOs

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasTrueOs() bool`

HasTrueOs returns a boolean if a field has been set.

### GetOsAnomaly

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetOsAnomaly() string`

GetOsAnomaly returns the OsAnomaly field if non-nil, zero value otherwise.

### GetOsAnomalyOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetOsAnomalyOk() (*string, bool)`

GetOsAnomalyOk returns a tuple with the OsAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAnomaly

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetOsAnomaly(v string)`

SetOsAnomaly sets OsAnomaly field to given value.

### HasOsAnomaly

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasOsAnomaly() bool`

HasOsAnomaly returns a boolean if a field has been set.

### GetRooted

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### GetRemoteSoftware

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRemoteSoftware() bool`

GetRemoteSoftware returns the RemoteSoftware field if non-nil, zero value otherwise.

### GetRemoteSoftwareOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetRemoteSoftwareOk() (*bool, bool)`

GetRemoteSoftwareOk returns a tuple with the RemoteSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSoftware

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetRemoteSoftware(v bool)`

SetRemoteSoftware sets RemoteSoftware field to given value.

### HasRemoteSoftware

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasRemoteSoftware() bool`

HasRemoteSoftware returns a boolean if a field has been set.

### GetIpReputation

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetIpReputation() string`

GetIpReputation returns the IpReputation field if non-nil, zero value otherwise.

### GetIpReputationOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetIpReputationOk() (*string, bool)`

GetIpReputationOk returns a tuple with the IpReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReputation

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetIpReputation(v string)`

SetIpReputation sets IpReputation field to given value.

### HasIpReputation

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasIpReputation() bool`

HasIpReputation returns a boolean if a field has been set.

### GetDeviceFingerprintReuse

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetDeviceFingerprintReuse() int32`

GetDeviceFingerprintReuse returns the DeviceFingerprintReuse field if non-nil, zero value otherwise.

### GetDeviceFingerprintReuseOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetDeviceFingerprintReuseOk() (*int32, bool)`

GetDeviceFingerprintReuseOk returns a tuple with the DeviceFingerprintReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprintReuse

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetDeviceFingerprintReuse(v int32)`

SetDeviceFingerprintReuse sets DeviceFingerprintReuse field to given value.

### HasDeviceFingerprintReuse

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasDeviceFingerprintReuse() bool`

HasDeviceFingerprintReuse returns a boolean if a field has been set.

### GetSingleDeviceUsed

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetSingleDeviceUsed() bool`

GetSingleDeviceUsed returns the SingleDeviceUsed field if non-nil, zero value otherwise.

### GetSingleDeviceUsedOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetSingleDeviceUsedOk() (*bool, bool)`

GetSingleDeviceUsedOk returns a tuple with the SingleDeviceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleDeviceUsed

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetSingleDeviceUsed(v bool)`

SetSingleDeviceUsed sets SingleDeviceUsed field to given value.

### HasSingleDeviceUsed

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasSingleDeviceUsed() bool`

HasSingleDeviceUsed returns a boolean if a field has been set.

### SetSingleDeviceUsedNil

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetSingleDeviceUsedNil(b bool)`

 SetSingleDeviceUsedNil sets the value for SingleDeviceUsed to be an explicit nil

### UnsetSingleDeviceUsed
`func (o *DeviceIntelligenceBreakdownPropertiesDevice) UnsetSingleDeviceUsed()`

UnsetSingleDeviceUsed ensures that no value is present for SingleDeviceUsed, not even an explicit nil
### GetDocumentCapture

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetDocumentCapture() string`

GetDocumentCapture returns the DocumentCapture field if non-nil, zero value otherwise.

### GetDocumentCaptureOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetDocumentCaptureOk() (*string, bool)`

GetDocumentCaptureOk returns a tuple with the DocumentCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCapture

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetDocumentCapture(v string)`

SetDocumentCapture sets DocumentCapture field to given value.

### HasDocumentCapture

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasDocumentCapture() bool`

HasDocumentCapture returns a boolean if a field has been set.

### GetBiometricCapture

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetBiometricCapture() string`

GetBiometricCapture returns the BiometricCapture field if non-nil, zero value otherwise.

### GetBiometricCaptureOk

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) GetBiometricCaptureOk() (*string, bool)`

GetBiometricCaptureOk returns a tuple with the BiometricCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricCapture

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) SetBiometricCapture(v string)`

SetBiometricCapture sets BiometricCapture field to given value.

### HasBiometricCapture

`func (o *DeviceIntelligenceBreakdownPropertiesDevice) HasBiometricCapture() bool`

HasBiometricCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


