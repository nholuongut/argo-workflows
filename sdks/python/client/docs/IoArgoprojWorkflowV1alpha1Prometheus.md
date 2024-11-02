# IonholuongutWorkflowV1alpha1Prometheus

Prometheus is a prometheus metric to be emitted

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**help** | **str** | Help is a string that describes the metric | 
**name** | **str** | Name is the name of the metric | 
**counter** | [**IonholuongutWorkflowV1alpha1Counter**](IonholuongutWorkflowV1alpha1Counter.md) |  | [optional] 
**gauge** | [**IonholuongutWorkflowV1alpha1Gauge**](IonholuongutWorkflowV1alpha1Gauge.md) |  | [optional] 
**histogram** | [**IonholuongutWorkflowV1alpha1Histogram**](IonholuongutWorkflowV1alpha1Histogram.md) |  | [optional] 
**labels** | [**[IonholuongutWorkflowV1alpha1MetricLabel]**](IonholuongutWorkflowV1alpha1MetricLabel.md) | Labels is a list of metric labels | [optional] 
**when** | **str** | When is a conditional statement that decides when to emit the metric | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


