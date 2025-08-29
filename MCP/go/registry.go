package main

import (
	"github.com/assurance-sites-energy-api/mcp-server/config"
	"github.com/assurance-sites-energy-api/mcp-server/models"
	tools_energy "github.com/assurance-sites-energy-api/mcp-server/tools/energy"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_energy.CreateCountsitesenergyTool(cfg),
		tools_energy.CreateQuerysitesenergytaskTool(cfg),
		tools_energy.CreateQuerysitesenergyTool(cfg),
		tools_energy.CreateQuerysitesenergycountTool(cfg),
		tools_energy.CreateQuerysitesenergycounttaskTool(cfg),
		tools_energy.CreateReadsitesenergybyidTool(cfg),
		tools_energy.CreateReadsitesenergyTool(cfg),
	}
}
