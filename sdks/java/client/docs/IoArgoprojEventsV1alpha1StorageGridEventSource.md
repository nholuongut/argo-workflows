

# IonholuongutEventsV1alpha1StorageGridEventSource


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**apiURL** | **String** | APIURL is the url of the storagegrid api. |  [optional]
**authToken** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]
**bucket** | **String** | Name of the bucket to register notifications for. |  [optional]
**events** | **List&lt;String&gt;** |  |  [optional]
**filter** | [**IonholuongutEventsV1alpha1StorageGridFilter**](IonholuongutEventsV1alpha1StorageGridFilter.md) |  |  [optional]
**metadata** | **Map&lt;String, String&gt;** |  |  [optional]
**region** | **String** |  |  [optional]
**topicArn** | **String** |  |  [optional]
**webhook** | [**IonholuongutEventsV1alpha1WebhookContext**](IonholuongutEventsV1alpha1WebhookContext.md) |  |  [optional]



