package main

import (
	"github.com/open-api-spec-for-assurance-fabric-virtual-networks/mcp-server/config"
	"github.com/open-api-spec-for-assurance-fabric-virtual-networks/mcp-server/models"
	tools_virtualnetworkhealthsummaries "github.com/open-api-spec-for-assurance-fabric-virtual-networks/mcp-server/tools/virtualnetworkhealthsummaries"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_virtualnetworkhealthsummaries.CreateVirtualnetworksummarybyidTool(cfg),
		tools_virtualnetworkhealthsummaries.CreateReadvirtualnetworktrendbyidTool(cfg),
		tools_virtualnetworkhealthsummaries.CreateReadvirtualnetworkssummaryTool(cfg),
		tools_virtualnetworkhealthsummaries.CreateReadvirtualnetworkscountTool(cfg),
	}
}
