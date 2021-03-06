# Go API client for teamspeak

Teamspeak WebQuery

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.12.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./teamspeak"
```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:10080/1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ChannelApi* | [**ChannelCreate**](docs/ChannelApi.md#channelcreate) | **Post** /channelcreate | Create a new channel
*ChannelApi* | [**ChannelDelete**](docs/ChannelApi.md#channeldelete) | **Post** /channeldelete | Delete channel
*ChannelApi* | [**ChannelEdit**](docs/ChannelApi.md#channeledit) | **Post** /channeledit | Edit channel
*ChannelApi* | [**ChannelInfo**](docs/ChannelApi.md#channelinfo) | **Get** /channelinfo | Get info about a channel


## Documentation For Models

 - [ChannelCreateResponse](docs/ChannelCreateResponse.md)
 - [ChannelCreateResponseBody](docs/ChannelCreateResponseBody.md)
 - [ChannelInfoResponse](docs/ChannelInfoResponse.md)
 - [ChannelInfoResponseBody](docs/ChannelInfoResponseBody.md)
 - [MinimalResponse](docs/MinimalResponse.md)
 - [Status](docs/Status.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



