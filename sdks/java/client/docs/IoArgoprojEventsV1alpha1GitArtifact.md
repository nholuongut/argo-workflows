

# IonholuongutEventsV1alpha1GitArtifact


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**branch** | **String** |  |  [optional]
**cloneDirectory** | **String** | Directory to clone the repository. We clone complete directory because GitArtifact is not limited to any specific Git service providers. Hence we don&#39;t use any specific git provider client. |  [optional]
**creds** | [**IonholuongutEventsV1alpha1GitCreds**](IonholuongutEventsV1alpha1GitCreds.md) |  |  [optional]
**filePath** | **String** |  |  [optional]
**ref** | **String** |  |  [optional]
**remote** | [**IonholuongutEventsV1alpha1GitRemoteConfig**](IonholuongutEventsV1alpha1GitRemoteConfig.md) |  |  [optional]
**sshKeySecret** | [**io.kubernetes.client.openapi.models.V1SecretKeySelector**](io.kubernetes.client.openapi.models.V1SecretKeySelector.md) |  |  [optional]
**tag** | **String** |  |  [optional]
**url** | **String** |  |  [optional]



