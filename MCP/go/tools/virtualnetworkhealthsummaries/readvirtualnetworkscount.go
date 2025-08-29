package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/open-api-spec-for-assurance-fabric-virtual-networks/mcp-server/config"
	"github.com/open-api-spec-for-assurance-fabric-virtual-networks/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ReadvirtualnetworkscountHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["vnLayer"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vnLayer=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/virtualNetworkHealthSummaries/count%s", cfg.BaseURL, queryString)
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
		var result map[string]interface{}
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

func CreateReadvirtualnetworkscountTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_virtualNetworkHealthSummaries_count",
		mcp.WithDescription("Read Virtual Networks count"),
		mcp.WithString("X-CALLER-ID", mcp.Description("Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.\n")),
		mcp.WithNumber("startTime", mcp.Description("Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n")),
		mcp.WithNumber("endTime", mcp.Description("End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n")),
		mcp.WithArray("siteHierarchy", mcp.Description("The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named \"Global\" (Ex. `Global/AreaName/BuildingName/FloorName`)\n\nThis field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*`\n\nExamples:\n\n`?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested)\n\n`?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)\n")),
		mcp.WithArray("siteHierarchyId", mcp.Description("The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`)\n\nThis field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*`\n\nExamples:\n\n`?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested)\n\n`?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)\n")),
		mcp.WithArray("id", mcp.Description("The list of entity Uuids. (Ex.\"6bef213c-19ca-4170-8375-b694e251101c\")\nExamples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested)\nid=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)\n")),
		mcp.WithString("vnLayer", mcp.Description("VN Layer information covering Layer 3 or Layer 2 VNs.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ReadvirtualnetworkscountHandler(cfg),
	}
}
