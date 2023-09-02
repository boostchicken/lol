# Documentation for BoostchickenLOL

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://lol.boostchicken.io*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *DefaultApi* | [**addCommand**](Apis/DefaultApi.md#addcommand) | **PUT** /add/{command}/{type} | Add a command  to the config |
*DefaultApi* | [**deleteCommand**](Apis/DefaultApi.md#deletecommand) | **DELETE** /delete/{command} | Delete a command from the config |
*DefaultApi* | [**getConfig**](Apis/DefaultApi.md#getconfig) | **GET** /config |  |
*DefaultApi* | [**getHistory**](Apis/DefaultApi.md#gethistory) | **GET** /history | Get all history tab entries (max 250) |
*DefaultApi* | [**getLiveConfig**](Apis/DefaultApi.md#getliveconfig) | **GET** /liveconfig | Get current configuration in JSON for UI |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Commands_inner](./Models/Commands_inner.md)
 - [HistoryEntry](./Models/HistoryEntry.md)
 - [LOLConfig](./Models/LOLConfig.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
