package main

import (
	"github.com/amazon-appintegrations-service/mcp-server/config"
	"github.com/amazon-appintegrations-service/mcp-server/models"
	tools_dataintegrations "github.com/amazon-appintegrations-service/mcp-server/tools/dataintegrations"
	tools_eventintegrations "github.com/amazon-appintegrations-service/mcp-server/tools/eventintegrations"
	tools_tags "github.com/amazon-appintegrations-service/mcp-server/tools/tags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_dataintegrations.CreateListdataintegrationassociationsTool(cfg),
		tools_eventintegrations.CreateListeventintegrationsTool(cfg),
		tools_eventintegrations.CreateCreateeventintegrationTool(cfg),
		tools_eventintegrations.CreateDeleteeventintegrationTool(cfg),
		tools_eventintegrations.CreateGeteventintegrationTool(cfg),
		tools_eventintegrations.CreateUpdateeventintegrationTool(cfg),
		tools_eventintegrations.CreateListeventintegrationassociationsTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_dataintegrations.CreateListdataintegrationsTool(cfg),
		tools_dataintegrations.CreateCreatedataintegrationTool(cfg),
		tools_dataintegrations.CreateDeletedataintegrationTool(cfg),
		tools_dataintegrations.CreateGetdataintegrationTool(cfg),
		tools_dataintegrations.CreateUpdatedataintegrationTool(cfg),
	}
}
