package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// CreateEventIntegrationrequest represents the CreateEventIntegrationrequest schema from the OpenAPI specification
type CreateEventIntegrationrequest struct {
	Description string `json:"Description,omitempty"` // The description of the event integration.
	Eventbridgebus string `json:"EventBridgeBus"` // The EventBridge bus.
	Eventfilter CreateEventIntegrationrequestEventFilter `json:"EventFilter"` // The event filter.
	Name string `json:"Name"` // The name of the event integration.
	Tags map[string]interface{} `json:"Tags,omitempty"` // The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Clienttoken string `json:"ClientToken,omitempty"` // A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If not provided, the Amazon Web Services SDK populates this field. For more information about idempotency, see <a href="https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/">Making retries safe with idempotent APIs</a>.
}

// UpdateEventIntegrationrequest represents the UpdateEventIntegrationrequest schema from the OpenAPI specification
type UpdateEventIntegrationrequest struct {
	Description string `json:"Description,omitempty"` // The description of the event inegration.
}

// ClientAssociationMetadata represents the ClientAssociationMetadata schema from the OpenAPI specification
type ClientAssociationMetadata struct {
}

// CreateDataIntegrationRequest represents the CreateDataIntegrationRequest schema from the OpenAPI specification
type CreateDataIntegrationRequest struct {
	Kmskey interface{} `json:"KmsKey"`
	Scheduleconfig CreateDataIntegrationResponseScheduleConfiguration `json:"ScheduleConfig"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Fileconfiguration CreateDataIntegrationResponseFileConfiguration `json:"FileConfiguration,omitempty"`
	Sourceuri interface{} `json:"SourceURI"`
	Tags interface{} `json:"Tags,omitempty"`
	Name interface{} `json:"Name"`
	Objectconfiguration interface{} `json:"ObjectConfiguration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// CreateDataIntegrationrequest represents the CreateDataIntegrationrequest schema from the OpenAPI specification
type CreateDataIntegrationrequest struct {
	Name string `json:"Name"` // The name of the DataIntegration.
	Clienttoken string `json:"ClientToken,omitempty"` // A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If not provided, the Amazon Web Services SDK populates this field. For more information about idempotency, see <a href="https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/">Making retries safe with idempotent APIs</a>.
	Description string `json:"Description,omitempty"` // A description of the DataIntegration.
	Objectconfiguration map[string]interface{} `json:"ObjectConfiguration,omitempty"` // The configuration for what data should be pulled from the source.
	Scheduleconfig CreateDataIntegrationrequestScheduleConfig `json:"ScheduleConfig"` // The name of the data and how often it should be pulled from the source.
	Sourceuri string `json:"SourceURI"` // The URI of the data source.
	Tags map[string]interface{} `json:"Tags,omitempty"` // The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Kmskey string `json:"KmsKey"` // The KMS key for the DataIntegration.
	Fileconfiguration CreateDataIntegrationrequestFileConfiguration `json:"FileConfiguration,omitempty"` // The configuration for what files should be pulled from the source.
}

// DeleteDataIntegrationRequest represents the DeleteDataIntegrationRequest schema from the OpenAPI specification
type DeleteDataIntegrationRequest struct {
}

// DeleteEventIntegrationResponse represents the DeleteEventIntegrationResponse schema from the OpenAPI specification
type DeleteEventIntegrationResponse struct {
}

// CreateDataIntegrationrequestScheduleConfig represents the CreateDataIntegrationrequestScheduleConfig schema from the OpenAPI specification
type CreateDataIntegrationrequestScheduleConfig struct {
	Firstexecutionfrom interface{} `json:"FirstExecutionFrom,omitempty"`
	Object interface{} `json:"Object,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
}

// ListDataIntegrationAssociationsRequest represents the ListDataIntegrationAssociationsRequest schema from the OpenAPI specification
type ListDataIntegrationAssociationsRequest struct {
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// GetDataIntegrationResponse represents the GetDataIntegrationResponse schema from the OpenAPI specification
type GetDataIntegrationResponse struct {
	Sourceuri interface{} `json:"SourceURI,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Fileconfiguration CreateDataIntegrationResponseFileConfiguration `json:"FileConfiguration,omitempty"`
	Kmskey interface{} `json:"KmsKey,omitempty"`
	Scheduleconfiguration CreateDataIntegrationResponseScheduleConfiguration `json:"ScheduleConfiguration,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Objectconfiguration interface{} `json:"ObjectConfiguration,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// UpdateDataIntegrationResponse represents the UpdateDataIntegrationResponse schema from the OpenAPI specification
type UpdateDataIntegrationResponse struct {
}

// ListEventIntegrationsResponse represents the ListEventIntegrationsResponse schema from the OpenAPI specification
type ListEventIntegrationsResponse struct {
	Eventintegrations interface{} `json:"EventIntegrations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListDataIntegrationAssociationsResponse represents the ListDataIntegrationAssociationsResponse schema from the OpenAPI specification
type ListDataIntegrationAssociationsResponse struct {
	Dataintegrationassociations interface{} `json:"DataIntegrationAssociations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListEventIntegrationsRequest represents the ListEventIntegrationsRequest schema from the OpenAPI specification
type ListEventIntegrationsRequest struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// CreateDataIntegrationResponseFileConfiguration represents the CreateDataIntegrationResponseFileConfiguration schema from the OpenAPI specification
type CreateDataIntegrationResponseFileConfiguration struct {
	Folders interface{} `json:"Folders"`
	Filters interface{} `json:"Filters,omitempty"`
}

// EventIntegrationEventFilter represents the EventIntegrationEventFilter schema from the OpenAPI specification
type EventIntegrationEventFilter struct {
	Source interface{} `json:"Source"`
}

// UpdateEventIntegrationRequest represents the UpdateEventIntegrationRequest schema from the OpenAPI specification
type UpdateEventIntegrationRequest struct {
	Description interface{} `json:"Description,omitempty"`
}

// CreateDataIntegrationrequestFileConfiguration represents the CreateDataIntegrationrequestFileConfiguration schema from the OpenAPI specification
type CreateDataIntegrationrequestFileConfiguration struct {
	Filters interface{} `json:"Filters,omitempty"`
	Folders interface{} `json:"Folders,omitempty"`
}

// ListDataIntegrationsResponse represents the ListDataIntegrationsResponse schema from the OpenAPI specification
type ListDataIntegrationsResponse struct {
	Dataintegrations interface{} `json:"DataIntegrations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EventFilter represents the EventFilter schema from the OpenAPI specification
type EventFilter struct {
	Source interface{} `json:"Source"`
}

// GetEventIntegrationResponseEventFilter represents the GetEventIntegrationResponseEventFilter schema from the OpenAPI specification
type GetEventIntegrationResponseEventFilter struct {
	Source interface{} `json:"Source"`
}

// TagResourcerequest represents the TagResourcerequest schema from the OpenAPI specification
type TagResourcerequest struct {
	Tags map[string]interface{} `json:"tags"` // The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
}

// GetEventIntegrationResponse represents the GetEventIntegrationResponse schema from the OpenAPI specification
type GetEventIntegrationResponse struct {
	Description interface{} `json:"Description,omitempty"`
	Eventbridgebus interface{} `json:"EventBridgeBus,omitempty"`
	Eventfilter GetEventIntegrationResponseEventFilter `json:"EventFilter,omitempty"`
	Eventintegrationarn interface{} `json:"EventIntegrationArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DataIntegrationAssociationSummary represents the DataIntegrationAssociationSummary schema from the OpenAPI specification
type DataIntegrationAssociationSummary struct {
	Dataintegrationarn interface{} `json:"DataIntegrationArn,omitempty"`
	Dataintegrationassociationarn interface{} `json:"DataIntegrationAssociationArn,omitempty"`
	Clientid interface{} `json:"ClientId,omitempty"`
}

// DeleteDataIntegrationResponse represents the DeleteDataIntegrationResponse schema from the OpenAPI specification
type DeleteDataIntegrationResponse struct {
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// CreateDataIntegrationResponseScheduleConfiguration represents the CreateDataIntegrationResponseScheduleConfiguration schema from the OpenAPI specification
type CreateDataIntegrationResponseScheduleConfiguration struct {
	Firstexecutionfrom interface{} `json:"FirstExecutionFrom,omitempty"`
	Object interface{} `json:"Object,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression"`
}

// ListEventIntegrationAssociationsResponse represents the ListEventIntegrationAssociationsResponse schema from the OpenAPI specification
type ListEventIntegrationAssociationsResponse struct {
	Eventintegrationassociations interface{} `json:"EventIntegrationAssociations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteEventIntegrationRequest represents the DeleteEventIntegrationRequest schema from the OpenAPI specification
type DeleteEventIntegrationRequest struct {
}

// ListDataIntegrationsRequest represents the ListDataIntegrationsRequest schema from the OpenAPI specification
type ListDataIntegrationsRequest struct {
}

// FileConfiguration represents the FileConfiguration schema from the OpenAPI specification
type FileConfiguration struct {
	Filters interface{} `json:"Filters,omitempty"`
	Folders interface{} `json:"Folders"`
}

// GetEventIntegrationRequest represents the GetEventIntegrationRequest schema from the OpenAPI specification
type GetEventIntegrationRequest struct {
}

// UpdateDataIntegrationrequest represents the UpdateDataIntegrationrequest schema from the OpenAPI specification
type UpdateDataIntegrationrequest struct {
	Description string `json:"Description,omitempty"` // A description of the DataIntegration.
	Name string `json:"Name,omitempty"` // The name of the DataIntegration.
}

// CreateEventIntegrationResponse represents the CreateEventIntegrationResponse schema from the OpenAPI specification
type CreateEventIntegrationResponse struct {
	Eventintegrationarn interface{} `json:"EventIntegrationArn,omitempty"`
}

// CreateDataIntegrationResponse represents the CreateDataIntegrationResponse schema from the OpenAPI specification
type CreateDataIntegrationResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Scheduleconfiguration CreateDataIntegrationResponseScheduleConfiguration `json:"ScheduleConfiguration,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Fileconfiguration CreateDataIntegrationResponseFileConfiguration `json:"FileConfiguration,omitempty"`
	Kmskey interface{} `json:"KmsKey,omitempty"`
	Objectconfiguration interface{} `json:"ObjectConfiguration,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Sourceuri interface{} `json:"SourceURI,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// FieldsMap represents the FieldsMap schema from the OpenAPI specification
type FieldsMap struct {
}

// UpdateEventIntegrationResponse represents the UpdateEventIntegrationResponse schema from the OpenAPI specification
type UpdateEventIntegrationResponse struct {
}

// CreateEventIntegrationRequest represents the CreateEventIntegrationRequest schema from the OpenAPI specification
type CreateEventIntegrationRequest struct {
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Eventbridgebus interface{} `json:"EventBridgeBus"`
	Eventfilter GetEventIntegrationResponseEventFilter `json:"EventFilter"`
}

// EventIntegration represents the EventIntegration schema from the OpenAPI specification
type EventIntegration struct {
	Description interface{} `json:"Description,omitempty"`
	Eventbridgebus interface{} `json:"EventBridgeBus,omitempty"`
	Eventfilter EventIntegrationEventFilter `json:"EventFilter,omitempty"`
	Eventintegrationarn interface{} `json:"EventIntegrationArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateEventIntegrationrequestEventFilter represents the CreateEventIntegrationrequestEventFilter schema from the OpenAPI specification
type CreateEventIntegrationrequestEventFilter struct {
	Source interface{} `json:"Source,omitempty"`
}

// GetDataIntegrationRequest represents the GetDataIntegrationRequest schema from the OpenAPI specification
type GetDataIntegrationRequest struct {
}

// ObjectConfiguration represents the ObjectConfiguration schema from the OpenAPI specification
type ObjectConfiguration struct {
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// DataIntegrationSummary represents the DataIntegrationSummary schema from the OpenAPI specification
type DataIntegrationSummary struct {
	Sourceuri interface{} `json:"SourceURI,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ScheduleConfiguration represents the ScheduleConfiguration schema from the OpenAPI specification
type ScheduleConfiguration struct {
	Object interface{} `json:"Object,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression"`
	Firstexecutionfrom interface{} `json:"FirstExecutionFrom,omitempty"`
}

// UpdateDataIntegrationRequest represents the UpdateDataIntegrationRequest schema from the OpenAPI specification
type UpdateDataIntegrationRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// EventIntegrationAssociation represents the EventIntegrationAssociation schema from the OpenAPI specification
type EventIntegrationAssociation struct {
	Clientassociationmetadata interface{} `json:"ClientAssociationMetadata,omitempty"`
	Clientid interface{} `json:"ClientId,omitempty"`
	Eventbridgerulename interface{} `json:"EventBridgeRuleName,omitempty"`
	Eventintegrationassociationarn interface{} `json:"EventIntegrationAssociationArn,omitempty"`
	Eventintegrationassociationid interface{} `json:"EventIntegrationAssociationId,omitempty"`
	Eventintegrationname interface{} `json:"EventIntegrationName,omitempty"`
}

// ListEventIntegrationAssociationsRequest represents the ListEventIntegrationAssociationsRequest schema from the OpenAPI specification
type ListEventIntegrationAssociationsRequest struct {
}
