

# CreateDataIntegrationRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | The name of the DataIntegration. |  |
|**description** | **String** | A description of the DataIntegration. |  [optional] |
|**kmsKey** | **String** | The KMS key for the DataIntegration. |  |
|**sourceURI** | **String** | The URI of the data source. |  |
|**scheduleConfig** | [**CreateDataIntegrationRequestScheduleConfig**](CreateDataIntegrationRequestScheduleConfig.md) |  |  |
|**tags** | **Map&lt;String, String&gt;** | The tags used to organize, track, or control access for this resource. For example, { \&quot;tags\&quot;: {\&quot;key1\&quot;:\&quot;value1\&quot;, \&quot;key2\&quot;:\&quot;value2\&quot;} }. |  [optional] |
|**clientToken** | **String** | A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If not provided, the Amazon Web Services SDK populates this field. For more information about idempotency, see &lt;a href&#x3D;\&quot;https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/\&quot;&gt;Making retries safe with idempotent APIs&lt;/a&gt;. |  [optional] |
|**fileConfiguration** | [**CreateDataIntegrationRequestFileConfiguration**](CreateDataIntegrationRequestFileConfiguration.md) |  |  [optional] |
|**objectConfiguration** | **Map&lt;String, Map&lt;String, List&lt;String&gt;&gt;&gt;** | The configuration for what data should be pulled from the source. |  [optional] |



