# \ChannelApi

All URIs are relative to *http://127.0.0.1:10080/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelCreate**](ChannelApi.md#ChannelCreate) | **Post** /channelcreate | Create a new channel
[**ChannelList**](ChannelApi.md#ChannelList) | **Get** /channellist | Gets list of channels



## ChannelCreate

> ChannelCreate(ctx, channelName, optional)

Create a new channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelName** | **string**| Channel&#39;s name | 
 **optional** | ***ChannelCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChannelCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelTopic** | **optional.String**| Channel&#39;s topic | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelList

> ChannelListResponse ChannelList(ctx, )

Gets list of channels

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

