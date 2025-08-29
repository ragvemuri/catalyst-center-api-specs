package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SiteEnergyResponse represents the SiteEnergyResponse schema from the OpenAPI specification
type SiteEnergyResponse struct {
	Response SiteEnergy `json:"response,omitempty"` // Assurance site and corresponding energy details
	Version string `json:"version,omitempty"` // The version of the response
}

// AsyncInfo represents the AsyncInfo schema from the OpenAPI specification
type AsyncInfo struct {
	Taskid string `json:"taskId,omitempty"` // the specific task id associated with the specific request made
	Tasklocation string `json:"taskLocation,omitempty"` // url resource where the client can fetch their task's lifecycle updates
}

// SiteEnergyFilter represents the SiteEnergyFilter schema from the OpenAPI specification
type SiteEnergyFilter struct {
	Logicaloperator string `json:"logicalOperator,omitempty"` // Operator to use when attempting to apply a logical conjunction of more than 1 filter Logical operations include: 'and', 'or'.
	Operator string `json:"operator,omitempty"` // Type of filter operator to use for querying data | in and out operator takes multiple values and applies the filters
	Value map[string]interface{} `json:"value,omitempty"` // Field value(s) to filter the data set. Array of values is used for "in" or "out" operator. Values will be of whatever type the specific field being filtered is defined with. For other operators, filter value is of whatever type the specific field being filtered is defined with. In the case of an "and" or "or" operator, this values array will be ignored, and the values arrays in each of the *nested filters* will be used.
	Filters []SiteEnergyFilter `json:"filters,omitempty"` // Nested array of filters in case of AND/OR based filters. Only one level of nesting will be supported. Structure of nested filter is the same as parent with all operators supported except AND or OR.
	Key string `json:"key,omitempty"` // Field names which are supported by this API for filter keys.
}

// SiteEnergyPaginationWithAggregateSortBy represents the SiteEnergyPaginationWithAggregateSortBy schema from the OpenAPI specification
type SiteEnergyPaginationWithAggregateSortBy struct {
	Count int `json:"count,omitempty"` // Total number of records related to the resource after applying applicable filtering
	Limit int `json:"limit,omitempty"` // The reference to the limit applied in the Pagination Request.
	Offset int `json:"offset,omitempty"` // The reference to the record offset applied in the Pagination Request.
	Sortby []SiteEnergySortByObj `json:"sortBy,omitempty"`
}

// CountIntegerResponse represents the CountIntegerResponse schema from the OpenAPI specification
type CountIntegerResponse struct {
	Response CountIntegerResponseresponse `json:"response,omitempty"`
	Version string `json:"version,omitempty"` // The version of the response
}

// ErrorObject represents the ErrorObject schema from the OpenAPI specification
type ErrorObject struct {
	Detail string `json:"detail,omitempty"` // A more detailed explanation of the error condition the parameter and its value, that caused the condition and why it caused it.
	Errorcode int `json:"errorCode,omitempty"` // Application specific error code returned by the server
	Message string `json:"message,omitempty"` // Brief message about the error condition
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Response []ErrorObject `json:"response,omitempty"`
	Version string `json:"version,omitempty"` // The version of the response
}

// AsyncInfoContainer represents the AsyncInfoContainer schema from the OpenAPI specification
type AsyncInfoContainer struct {
	Version string `json:"version,omitempty"`
	Response AsyncInfo `json:"response,omitempty"` // Generic model with information on the asynchronous task/process that was created to handle an API request
}

// SiteEnergyAggregateAttribute represents the SiteEnergyAggregateAttribute schema from the OpenAPI specification
type SiteEnergyAggregateAttribute struct {
	Function string `json:"function,omitempty"` // Type of aggregate function to apply on the field when querying data
	Name string `json:"name,omitempty"` // Field name on which the aggregate function should be applied when querying the data. The fields supported for aggregation in the API response object are, [energyConsumed, estimatedEmission, estimatedCost, carbonIntensity]
}

// SiteEnergySortByObj represents the SiteEnergySortByObj schema from the OpenAPI specification
type SiteEnergySortByObj struct {
	Name string `json:"name,omitempty"` // Field names which are supported by this API as attributes or filter keys
	Order string `json:"order,omitempty"` // The sort order of the field ascending or descending.
	Function string `json:"function,omitempty"` // Type of aggregate function to apply on the field when querying data
}

// SiteEnergyAggregateAttributesResponseinner represents the SiteEnergyAggregateAttributesResponseinner schema from the OpenAPI specification
type SiteEnergyAggregateAttributesResponseinner struct {
	Function string `json:"function,omitempty"` // Type of aggregate function to apply on the field when querying data
	Name string `json:"name,omitempty"` // Field name on which the aggregate function should be applied when querying the data. The fields supported for aggregation in the API response object are, [energyConsumed, estimatedEmission, estimatedCost, carbonIntensity]
	Value float64 `json:"value,omitempty"` // The value based on the functions requested in the input data
}

// SitesEnergyQueryResponse represents the SitesEnergyQueryResponse schema from the OpenAPI specification
type SitesEnergyQueryResponse struct {
	Page SiteEnergyPaginationWithAggregateSortBy `json:"page,omitempty"` // Pagination model with support for aggregate sort by.
	Response []interface{} `json:"response,omitempty"` // Successfully returned list of Sites Energy information
	Version string `json:"version,omitempty"` // The version of the response
}

// CountIntegerResponseresponse represents the CountIntegerResponseresponse schema from the OpenAPI specification
type CountIntegerResponseresponse struct {
	Count int64 `json:"count,omitempty"` // The total number of records related to the resource
}

// SitesEnergyQueryRequest represents the SitesEnergyQueryRequest schema from the OpenAPI specification
type SitesEnergyQueryRequest struct {
	Starttime int64 `json:"startTime,omitempty"` // Start time from which the API queries the dataset related to the resource. It must be specified in terms of milliseconds since UNIX epoch. Value is inclusive. If `startTime` is not provided, API will default to current time.
	Endtime int64 `json:"endTime,omitempty"` // End time to which the API queries the dataset related to the resource. It must be specified in terms of milliseconds since UNIX epoch. Value is inclusive. If `endTime` is not provided, API will default to current time.
}

// StartAndEndTime represents the StartAndEndTime schema from the OpenAPI specification
type StartAndEndTime struct {
	Starttime int64 `json:"startTime,omitempty"` // Start time from which the API queries the dataset related to the resource. It must be specified in terms of milliseconds since UNIX epoch. Value is inclusive. If `startTime` is not provided, API will default to current time.
	Endtime int64 `json:"endTime,omitempty"` // End time to which the API queries the dataset related to the resource. It must be specified in terms of milliseconds since UNIX epoch. Value is inclusive. If `endTime` is not provided, API will default to current time.
}

// SiteEnergy represents the SiteEnergy schema from the OpenAPI specification
type SiteEnergy struct {
	Id string `json:"id,omitempty"` // Unique ID of the Site.
	Numberofdevices float64 `json:"numberOfDevices,omitempty"` // number of devices for the sites that are returning energy data.
	Sitetype string `json:"siteType,omitempty"` // Type of site For aggregation requests, the value may be null since the unique identifier may not apply
	Energyconsumed float64 `json:"energyConsumed,omitempty"` // Total energy consumed in kWh.
	Estimatedcost float64 `json:"estimatedCost,omitempty"` // Estimate of the total financial cost.
	Sitehierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy.
	Sitehierarchyid string `json:"siteHierarchyId,omitempty"` // Site hierarchy ID.
	Carbonintensity float64 `json:"carbonIntensity,omitempty"` // Carbon intensity.
	Latitude float64 `json:"latitude,omitempty"` // Longitude of the site.
	Longitude float64 `json:"longitude,omitempty"` // Longitude of the site.
	Sitename string `json:"siteName,omitempty"` // Site name.
	Devicecategories []string `json:"deviceCategories,omitempty"` // Device categories.
	Estimatedemission float64 `json:"estimatedEmission,omitempty"` // Estimate of the total emission in kg CO2.
}

// SitesEnergyResponse represents the SitesEnergyResponse schema from the OpenAPI specification
type SitesEnergyResponse struct {
	Version string `json:"version,omitempty"` // The version of the response
	Page SiteEnergyPaginationWithAggregateSortBy `json:"page,omitempty"` // Pagination model with support for aggregate sort by.
	Response []SiteEnergy `json:"response,omitempty"` // Successfully returned list of Sites Energy information
}
