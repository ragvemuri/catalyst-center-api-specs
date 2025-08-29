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

func ReadsitesenergybyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["startTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startTime=%v", val))
		}
		if val, ok := args["endTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endTime=%v", val))
		}
		if val, ok := args["views"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("views=%v", val))
		}
		if val, ok := args["attribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("attribute=%v", val))
		}
		if val, ok := args["deviceCategory"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("deviceCategory=%v", val))
		}
		if val, ok := args["taskId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("taskId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/dna/data/api/v1/energy/sites/%s%s", cfg.BaseURL, id, queryString)
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

func CreateReadsitesenergybyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_dna_data_api_v1_energy_sites_id",
		mcp.WithDescription("Get site energy by ID"),
		mcp.WithString("id", mcp.Required(), mcp.Description("The UUID of the Site. (Ex. \"6bef213c-19ca-4170-8375-b694e251101c\")")),
		mcp.WithNumber("startTime", mcp.Description("Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n\nIf `startTime` is not provided, API will default to one day before `endTime`.\n")),
		mcp.WithNumber("endTime", mcp.Description("End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n\nIf `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.\n")),
		mcp.WithArray("views", mcp.Description("The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites.\n\n### Response data proviced by each view:  \n\n1. **Site**\n[id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude]  \n\n2. **Energy**\n[energyConsumed, estimatedCost, estimatedEmission, carbonIntensity, numberOfDevices]   \n\nWhen this query parameter is not added the default summaries are:  \n\n**[site,energy]**\n\nExamples:\n\nviews=site (single view requested)\n\nviews=site,energy (multiple views requested)\n")),
		mcp.WithArray("attribute", mcp.Description("Supported Attributes:\n\n[id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, energyConsumed, estimatedCost, estimatedEmission, carbonIntensity, numberOfDevices]\n\nIf length of attribute list is too long, please use 'view' param instead.\n\nExamples:\n\nattribute=siteHierarchy (single attribute requested)\n\nattribute=siteHierarchy&attribute=energyConsumed (multiple attributes requested)\n")),
		mcp.WithArray("deviceCategory", mcp.Description("The list of device categories. Note that this filter specifies which devices will be included when calculating energy consumption values, rather than specifying the list of returned sites. \n\nExamples:\n\n`deviceCategory=Switch` (single device category requested)\n\n`deviceCategory=Switch&deviceCategory=Router` (multiple device categories with comma separator)\n")),
		mcp.WithString("X-CALLER-ID", mcp.Description("Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.\n")),
		mcp.WithString("taskId", mcp.Description("used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ReadsitesenergybyidHandler(cfg),
	}
}
