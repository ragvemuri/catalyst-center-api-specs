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

func VirtualnetworksummarybyidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["endTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endTime=%v", val))
		}
		if val, ok := args["startTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startTime=%v", val))
		}
		if val, ok := args["attribute"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("attribute=%v", val))
		}
		if val, ok := args["view"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("view=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/virtualNetworkHealthSummaries/%s%s", cfg.BaseURL, id, queryString)
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

func CreateVirtualnetworksummarybyidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_virtualNetworkHealthSummaries_id",
		mcp.WithDescription("Read virtual network with its health summary from id"),
		mcp.WithString("id", mcp.Required(), mcp.Description("unique virtual networks id")),
		mcp.WithString("X-CALLER-ID", mcp.Description("Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.\n")),
		mcp.WithNumber("endTime", mcp.Description("End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n")),
		mcp.WithNumber("startTime", mcp.Description("Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.\n")),
		mcp.WithArray("attribute", mcp.Description("Supported Attributes:\n[id,name,networkProtocol,vlan,vnid,layer,associatedL3Vn,totalEndpoints,totalFabricSites,goodHealthPercentage,goodHealthDeviceCount,totalHealthDeviceCount,poorHealthDeviceCount,fairHealthDeviceCount,noHealthDeviceCount,vnFabricControlPlaneGoodHealthPercentage,vnFabricControlPlaneGoodHealthDeviceCount, vnFabricControlPlaneTotalDeviceCount, vnFabricControlPlanePoorHealthDeviceCount, vnFabricControlPlaneFairHealthDeviceCount,vnFabricControlPlaneNoHealthDeviceCount, vnServicesHealthPercentage, vnServicesTotalDeviceCount, vnServicesGoodHealthDeviceCount, vnServicesPoorHealthDeviceCount, vnServicesFairHealthDeviceCount, vnServicesNoHealthDeviceCount, vnExitHealthPercentage, vnExitTotalDeviceCount, vnExitGoodHealthDeviceCount, vnExitPoorHealthDeviceCount, vnExitFairHealthDeviceCount, vnExitNoHealthDeviceCount , vnStatusHealthPercentage, vnStatusTotalDeviceCount,vnStatusGoodHealthDeviceCount, vnStatusPoorHealthDeviceCount, vnStatusFairHealthDeviceCount, vnStatusNoHealthDeviceCount, pubsubSessionGoodHealthPercentage, pubsubSessionTotalDeviceCount, pubsubSessionGoodHealthDeviceCount, pubsubSessionPoorHealthDeviceCount, pubsubSessionFairHealthDeviceCount, pubsubSessionNoHealthDeviceCount, multiCastGoodHealthPercentage, multiCastTotalDeviceCount, multiCastGoodHealthDeviceCount, multiCastPoorHealthDeviceCount, multiCastFairHealthDeviceCount, multiCastNoHealthDeviceCount, internetAvailGoodHealthPercentage, internetAvailTotalDeviceCount,internetAvailGoodHealthDeviceCount, internetAvailPoorHealthDeviceCount, internetAvailFairHealthDeviceCount, internetAvailNoHealthDeviceCount, bgpPeerGoodHealthPercentage, bgpPeerTotalDeviceCount, bgpPeerGoodHealthDeviceCount,bgpPeerPoorHealthDeviceCount,bgpPeerFairHealthDeviceCount, bgpPeerNoHealthDeviceCount,vniGoodHealthPercentage,vniTotalDeviceCount, vniGoodHealthDeviceCount, vniPoorHealthDeviceCount,vniFairHealthDeviceCount, vniNoHealthDeviceCount, firewallConnGoodHealthPercentage, firewallConnGoodHealthDeviceCount, firewallConnPoorHealthDeviceCount,firewallConnNoHealthDeviceCount, firewallConnFairHealthDeviceCount, firewallConnTotalDeviceCount]\n\nExamples:\n\nattribute=goodHealthPercentage (single attribute requested)\n\nattribute=goodHealthPercentage&attribute=totalEndpoints (multiple attributes requested) \n")),
		mcp.WithArray("view", mcp.Description("The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with virtual networks.\n\n### Response data proviced by each view:  \n\n1. **vnInfo**\n[id, name, networkProtocol, vlan, vnid, layer, associatedL3Vn, totalEndpoints, totalFabricSites]  \n\n2. **overallHealth**\n[id, name, goodHealthPercentage, goodHealthDeviceCount, totalHealthDeviceCount,poorHealthDeviceCount,fairHealthDeviceCount,noHealthDeviceCount,vnFabricControlPlaneGoodHealthPercentage,vnFabricControlPlaneGoodHealthDeviceCount, vnFabricControlPlaneTotalDeviceCount, vnFabricControlPlanePoorHealthDeviceCount, vnFabricControlPlaneFairHealthDeviceCount,vnFabricControlPlaneNoHealthDeviceCount, vnServicesHealthPercentage, vnServicesTotalDeviceCount, vnServicesGoodHealthDeviceCount, vnServicesPoorHealthDeviceCount, vnServicesFairHealthDeviceCount, vnServicesNoHealthDeviceCount, vnExitHealthPercentage, vnExitTotalDeviceCount, vnExitGoodHealthDeviceCount, vnExitPoorHealthDeviceCount, vnExitFairHealthDeviceCount, vnExitNoHealthDeviceCount, vnStatusHealthPercentage, vnStatusTotalDeviceCount,vnStatusGoodHealthDeviceCount, vnStatusPoorHealthDeviceCount, vnStatusFairHealthDeviceCount, vnStatusNoHealthDeviceCount]\n\n3. **vnControlPlaneHealth**\n[id, name, vnFabricControlPlaneGoodHealthPercentage, vnFabricControlPlaneGoodHealthDeviceCount, vnFabricControlPlaneTotalDeviceCount, vnFabricControlPlanePoorHealthDeviceCount, vnFabricControlPlaneFairHealthDeviceCount, vnFabricControlPlaneNoHealthDeviceCount, pubsubSessionGoodHealthPercentage, pubsubSessionTotalDeviceCount, pubsubSessionGoodHealthDeviceCount,pubsubSessionPoorHealthDeviceCount, pubsubSessionFairHealthDeviceCount, pubsubSessionNoHealthDeviceCount]\n\n4. **vnServicesHealth**\n[id, name, vnServicesHealthPercentage, vnServicesTotalDeviceCount, vnServicesGoodHealthDeviceCount,vnServicesPoorHealthDeviceCount, vnServicesFairHealthDeviceCount, vnServicesNoHealthDeviceCount, multiCastGoodHealthPercentage, multiCastTotalDeviceCount, multiCastGoodHealthDeviceCount,multiCastPoorHealthDeviceCount, multiCastFairHealthDeviceCount, multiCastNoHealthDeviceCount, internetAvailGoodHealthPercentage, internetAvailTotalDeviceCount, internetAvailGoodHealthDeviceCount,internetAvailPoorHealthDeviceCount, internetAvailFairHealthDeviceCount, internetAvailNoHealthDeviceCount,\n  firewallConnGoodHealthPercentage, firewallConnGoodHealthDeviceCount, firewallConnPoorHealthDeviceCount,\n  firewallConnNoHealthDeviceCount, firewallConnFairHealthDeviceCount, firewallConnTotalDeviceCount]\n\n5. **vnExitHealth**\n[id, name, vnExitHealthPercentage, vnExitTotalDeviceCount, vnExitGoodHealthDeviceCount, vnExitPoorHealthDeviceCount, vnExitFairHealthDeviceCount, vnExitNoHealthDeviceCount, bgpPeerGoodHealthPercentage, bgpPeerTotalDeviceCount, bgpPeerGoodHealthDeviceCount,bgpPeerPoorHealthDeviceCount ,bgpPeerFairHealthDeviceCount, bgpPeerNoHealthDeviceCount]\n\n6.**vnStatusHealth**\n[id, name, vnStatusTotalDeviceCount, vnStatusGoodHealthDeviceCount, vnStatusPoorHealthDeviceCount, vnStatusFairHealthDeviceCount, vnStatusNoHealthDeviceCount, vniGoodHealthPercentage, vniTotalDeviceCount, vniGoodHealthDeviceCount,vniPoorHealthDeviceCount, vniFairHealthDeviceCount, vniNoHealthDeviceCount]\n\nWhen this query parameter is not added the default summaries are:  \n\n**[vnInfo, overallHealth]**\n\nExamples:\n\nview=overallHealth (single view requested)\n\nview=vnInfo&view=overallHealth&view=vnServicesHealth (multiple views requested)\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    VirtualnetworksummarybyidHandler(cfg),
	}
}
