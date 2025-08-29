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

func ReadvirtualnetworktrendbyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["trendInterval"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("trendInterval=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order=%v", val))
		}
		if val, ok := args["attribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("attribute=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/virtualNetworkHealthSummaries/%s/trendAnalytics%s", cfg.BaseURL, id, queryString)
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

func CreateReadvirtualnetworktrendbyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_virtualNetworkHealthSummaries_id_trendAnalytics",
		mcp.WithDescription("The Trend analytics data for a virtual network in the specified time range"),
		mcp.WithString("id", mcp.Required(), mcp.Description("unique virtual network id")),
		mcp.WithString("X-CALLER-ID", mcp.Description("Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.\n")),
		mcp.WithNumber("startTime", mcp.Description("Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n")),
		mcp.WithNumber("endTime", mcp.Description("End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n")),
		mcp.WithString("trendInterval", mcp.Required(), mcp.Description("The time window to aggregate the metrics. \nInterval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days\n")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of records to return")),
		mcp.WithNumber("offset", mcp.Description("Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.")),
		mcp.WithString("order", mcp.Description("The sort order of the field ascending or descending.")),
		mcp.WithArray("attribute", mcp.Description("Supported Attributes:\n\n[goodHealthPercentage,goodHealthDeviceCount,totalHealthDeviceCount,poorHealthDeviceCount,fairHealthDeviceCount,noHealthDeviceCount,vnFabricControlPlaneGoodHealthPercentage,vnFabricControlPlaneGoodHealthDeviceCount, vnFabricControlPlaneTotalDeviceCount, vnFabricControlPlanePoorHealthDeviceCount, vnFabricControlPlaneFairHealthDeviceCount,vnFabricControlPlaneNoHealthDeviceCount, vnServicesHealthPercentage, vnServicesTotalDeviceCount, vnServicesGoodHealthDeviceCount, vnServicesPoorHealthDeviceCount, vnServicesFairHealthDeviceCount, vnServicesNoHealthDeviceCount, vnExitHealthPercentage, vnExitTotalDeviceCount, vnExitGoodHealthDeviceCount, vnExitPoorHealthDeviceCount, vnExitFairHealthDeviceCount, vnExitNoHealthDeviceCount , vnStatusHealthPercentage, vnStatusTotalDeviceCount,vnStatusGoodHealthDeviceCount, vnStatusPoorHealthDeviceCount, vnStatusFairHealthDeviceCount, vnStatusNoHealthDeviceCount, pubsubSessionGoodHealthPercentage, pubsubSessionTotalDeviceCount, pubsubSessionGoodHealthDeviceCount, pubsubSessionPoorHealthDeviceCount, pubsubSessionFairHealthDeviceCount, pubsubSessionNoHealthDeviceCount, multiCastGoodHealthPercentage, multiCastTotalDeviceCount, multiCastGoodHealthDeviceCount, multiCastPoorHealthDeviceCount, multiCastFairHealthDeviceCount, multiCastNoHealthDeviceCount, internetAvailGoodHealthPercentage, internetAvailTotalDeviceCount,internetAvailGoodHealthDeviceCount, internetAvailPoorHealthDeviceCount, internetAvailFairHealthDeviceCount, internetAvailNoHealthDeviceCount, bgpPeerGoodHealthPercentage, bgpPeerTotalDeviceCount, bgpPeerGoodHealthDeviceCount,bgpPeerPoorHealthDeviceCount,bgpPeerFairHealthDeviceCount, bgpPeerNoHealthDeviceCount,vniGoodHealthPercentage,vniTotalDeviceCount, vniGoodHealthDeviceCount, vniPoorHealthDeviceCount,vniFairHealthDeviceCount, vniNoHealthDeviceCount, firewallConnGoodHealthPercentage,firewallConnGoodHealthDeviceCount, firewallConnPoorHealthDeviceCount, firewallConnNoHealthDeviceCount,firewallConnFairHealthDeviceCount, firewallConnTotalDeviceCount]\n\nExamples:\n\nattribute=goodHealthPercentage (single attribute requested)\n\nattribute=goodHealthPercentage&attribute=totalFabricSites (multiple attributes requested)         \n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ReadvirtualnetworktrendbyidHandler(cfg),
	}
}
