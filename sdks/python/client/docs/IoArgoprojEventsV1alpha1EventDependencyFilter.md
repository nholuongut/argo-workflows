# IonholuongutEventsV1alpha1EventDependencyFilter

EventDependencyFilter defines filters and constraints for a io.nholuongut.workflow.v1alpha1.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**context** | [**IonholuongutEventsV1alpha1EventContext**](IonholuongutEventsV1alpha1EventContext.md) |  | [optional] 
**data** | [**[IonholuongutEventsV1alpha1DataFilter]**](IonholuongutEventsV1alpha1DataFilter.md) |  | [optional] 
**exprs** | [**[IonholuongutEventsV1alpha1ExprFilter]**](IonholuongutEventsV1alpha1ExprFilter.md) | Exprs contains the list of expressions evaluated against the event payload. | [optional] 
**time** | [**IonholuongutEventsV1alpha1TimeFilter**](IonholuongutEventsV1alpha1TimeFilter.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


