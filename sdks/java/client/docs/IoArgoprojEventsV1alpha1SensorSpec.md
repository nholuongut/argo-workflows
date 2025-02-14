

# IonholuongutEventsV1alpha1SensorSpec


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dependencies** | [**List&lt;IonholuongutEventsV1alpha1EventDependency&gt;**](IonholuongutEventsV1alpha1EventDependency.md) | Dependencies is a list of the events that this sensor is dependent on. |  [optional]
**errorOnFailedRound** | **Boolean** | ErrorOnFailedRound if set to true, marks sensor state as &#x60;error&#x60; if the previous trigger round fails. Once sensor state is set to &#x60;error&#x60;, no further triggers will be processed. |  [optional]
**eventBusName** | **String** |  |  [optional]
**replicas** | **Integer** |  |  [optional]
**template** | [**IonholuongutEventsV1alpha1Template**](IonholuongutEventsV1alpha1Template.md) |  |  [optional]
**triggers** | [**List&lt;IonholuongutEventsV1alpha1Trigger&gt;**](IonholuongutEventsV1alpha1Trigger.md) | Triggers is a list of the things that this sensor evokes. These are the outputs from this sensor. |  [optional]



