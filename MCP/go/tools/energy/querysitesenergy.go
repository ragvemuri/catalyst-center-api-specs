package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/assurance-sites-energy-api/mcp-server/config"
	"github.com/assurance-sites-energy-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func QuerysitesenergyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.SitesEnergyQueryRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/dna/data/api/v1/energy/sites/query", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateQuerysitesenergyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_dna_data_api_v1_energy_sites_query",
		mcp.WithDescription("Submit request to query sites energy"),
		mcp.WithString("X-CALLER-ID", mcp.Description("Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.\n")),
		mcp.WithNumber("endTime", mcp.Description("Input parameter: End time to which the API queries the dataset related to the resource. It must be specified in terms of milliseconds since UNIX epoch. Value is inclusive.\n\nIf `endTime` is not provided, API will default to current time.\n")),
		mcp.WithNumber("startTime", mcp.Description("Input parameter: Start time from which the API queries the dataset related to the resource. It must be specified in terms of milliseconds since UNIX epoch. Value is inclusive.\n\nIf `startTime` is not provided, API will default to current time.\n")),
		mcp.WithObject("page", mcp.Description("Input parameter: Pagination model with support for aggregate sort by.")),
		mcp.WithArray("views", mcp.Description("")),
		mcp.WithArray("aggregateAttributes", mcp.Description("")),
		mcp.WithArray("attributes", mcp.Description("")),
		mcp.WithArray("filters", mcp.Description("Input parameter: List of filters to apply when querying the site energy data.\n\nSupported operators are [in, eq, like]\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    QuerysitesenergyHandler(cfg),
	}
}
