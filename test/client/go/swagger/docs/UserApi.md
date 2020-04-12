# \UserApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAuth**](UserApi.md#UserAuth) | **Post** /api/v1/user/auth | 
[**UserInfo**](UserApi.md#UserInfo) | **Post** /api/v1/user/info | 
[**UserRefresh**](UserApi.md#UserRefresh) | **Post** /api/v1/user/refresh | 
[**UserRegister**](UserApi.md#UserRegister) | **Post** /api/v1/user/register | 
[**UserTest**](UserApi.md#UserTest) | **Post** /api/v1/user/test | 


# **UserAuth**
> BinacsApiUserV1UserAuthResp UserAuth(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BinacsApiUserV1UserAuthReq**](BinacsApiUserV1UserAuthReq.md)|  | 

### Return type

[**BinacsApiUserV1UserAuthResp**](binacs_api_user_v1UserAuthResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserInfo**
> BinacsApiUserV1UserInfoResp UserInfo(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BinacsApiUserV1UserInfoReq**](BinacsApiUserV1UserInfoReq.md)|  | 

### Return type

[**BinacsApiUserV1UserInfoResp**](binacs_api_user_v1UserInfoResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRefresh**
> BinacsApiUserV1UserRefreshResp UserRefresh(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BinacsApiUserV1UserRefreshReq**](BinacsApiUserV1UserRefreshReq.md)|  | 

### Return type

[**BinacsApiUserV1UserRefreshResp**](binacs_api_user_v1UserRefreshResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRegister**
> BinacsApiUserV1UserRegisterResp UserRegister(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BinacsApiUserV1UserRegisterReq**](BinacsApiUserV1UserRegisterReq.md)|  | 

### Return type

[**BinacsApiUserV1UserRegisterResp**](binacs_api_user_v1UserRegisterResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserTest**
> BinacsApiUserV1UserTestResp UserTest(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BinacsApiUserV1UserTestReq**](BinacsApiUserV1UserTestReq.md)|  | 

### Return type

[**BinacsApiUserV1UserTestResp**](binacs_api_user_v1UserTestResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

