package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/assurance-sites-energy-api/mcp-server/config"
	"github.com/assurance-sites-energy-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CountsitesenergyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["startTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startTime=%v", val))
		}
		if val, ok := args["endTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endTime=%v", val))
		}
		if val, ok := args["siteHierarchy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("siteHierarchy=%v", val))
		}
		if val, ok := args["siteHierarchyId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("siteHierarchyId=%v", val))
		}
		if val, ok := args["siteName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("siteName=%v", val))
		}
		if val, ok := args["siteType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("siteType=%v", val))
		}
		if val, ok := args["deviceCategory"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("deviceCategory=%v", val))
		}
		if val, ok := args["siteId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("siteId=%v", val))
		}
		if val, ok := args["taskId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("taskId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/dna/data/api/v1/energy/sites/count%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No specific authentication scheme defined - add fallback authentication
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", "Bearer "+cfg.BearerToken)
		} else if cfg.APIKey != "" {
			req.Header.Set("Authorization", "Bearer "+cfg.APIKey)
		} else if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", "Basic "+cfg.BasicAuth)
		}
		// Note: If no auth tokens provided, requests will be made without authentication
		
		// Add custom headers if provided
		
		// Set client identification headers
		req.Header.Set("X-Request-Source", "Codeglide-MCP-generator")
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-CALLER-ID"]; ok {
			req.Header.Set("X-CALLER-ID", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CountIntegerResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCountsitesenergyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_dna_data_api_v1_energy_sites_count",
		mcp.WithDescription("Count sites energy"),
		mcp.WithNumber("startTime", mcp.Description("Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n\nIf `startTime` is not provided, API will default to one day before `endTime`.\n")),
		mcp.WithNumber("endTime", mcp.Description("End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n\nIf `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.\n")),
		mcp.WithArray("siteHierarchy", mcp.Description("The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named \"Global\" (Ex. `Global/AreaName/BuildingName/FloorName`)\n\nThis field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*`\n\nExamples:\n\n`?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested)\n\n`?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)\n")),
		mcp.WithArray("siteHierarchyId", mcp.Description("The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`)\n\nThis field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*`\n\nExamples:\n\n`?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested)\n\n`?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)\n")),
		mcp.WithArray("siteName", mcp.Description("The name of the site. (Ex. `FloorName`)\n\nThis field supports wildcard asterisk (`*`) character search support. E.g. `*San*, *San, San*`\n\nExamples:\n\n`?siteName=building1` (single siteName requested)\n\n`?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)\n")),
		mcp.WithArray("siteType", mcp.Description("The type of the site. A site can be an area, building, or floor.\n\nDefault when not provided will be `[floor,building,area]`\n\nExamples:\n\n`?siteType=area` (single siteType requested)\n\n`?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)\n")),
		mcp.WithArray("deviceCategory", mcp.Description("The list of device categories. Note that this filter specifies which devices will be included when calculating energy consumption values, rather than specifying the list of returned sites. \n\nExamples:\n\n`deviceCategory=Switch` (single device category requested)\n\n`deviceCategory=Switch&deviceCategory=Router` (multiple device categories with comma separator)\n")),
		mcp.WithArray("siteId", mcp.Description("The UUID of the site. (Ex. `flooruuid`)\n\nExamples:\n\n`?siteId=id1` (single id requested)\n\n`?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)\n")),
		mcp.WithString("X-CALLER-ID", mcp.Description("Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.\n")),
		mcp.WithString("taskId", mcp.Description("used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CountsitesenergyHandler(cfg),
	}
}
