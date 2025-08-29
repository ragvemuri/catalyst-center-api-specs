package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PaginationResponsesortBy represents the PaginationResponsesortBy schema from the OpenAPI specification
type PaginationResponsesortBy struct {
	Name string `json:"name,omitempty"` // Field name by which sort is requested
	Order string `json:"order,omitempty"` // Sort order. 'asc' for ascending and 'desc' for descending
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Response []ErrorObject `json:"response,omitempty"`
	Version string `json:"version,omitempty"` // The version of the response
}

// PaginationResponse represents the PaginationResponse schema from the OpenAPI specification
type PaginationResponse struct {
	Count int `json:"count,omitempty"` // Total number of records related to the resource after applying applicable filtering
	Limit int `json:"limit,omitempty"` // The reference to the limit applied in the Pagination Request.
	Offset int `json:"offset,omitempty"` // The reference to the record offset applied in the Pagination Request.
	Sortby []PaginationResponsesortBy `json:"sortBy,omitempty"` // Reference to the sortBy that was applied in the Pagination Request.
}

// TimeBasedPagination represents the TimeBasedPagination schema from the OpenAPI specification
type TimeBasedPagination struct {
	Offset int `json:"offset,omitempty"` // Starting offset of data to fetch and returned
	Timesortorder string `json:"timeSortOrder,omitempty"` // Sort order. `asc` for ascending and `desc` for descending
	Count int `json:"count,omitempty"` // Number of records returned after applying applicable filtering. Field is ignored for request and updated by API in the response
	Limit int `json:"limit,omitempty"` // Number of records to fetch in a page
}

// VirtualNetworkSummary represents the VirtualNetworkSummary schema from the OpenAPI specification
type VirtualNetworkSummary struct {
	Firewallconngoodhealthdevicecount int64 `json:"firewallConnGoodHealthDeviceCount,omitempty"` // Firewall Connectivity Good Health Device Count
	Bgppeergoodhealthpercentage float64 `json:"bgpPeerGoodHealthPercentage,omitempty"` // Bgp Peer Good Health Percentage
	Networkprotocol string `json:"networkProtocol,omitempty"` // Protocol of associated fabric site(s)
	Vnstatusgoodhealthdevicecount int64 `json:"vnStatusGoodHealthDeviceCount,omitempty"` // Vn Status Good Health Device Count
	Totalfabricsites int64 `json:"totalFabricSites,omitempty"` // total number of FabricSites in the Virtual Network
	Vnipoorhealthdevicecount int64 `json:"vniPoorHealthDeviceCount,omitempty"` // Vni Poor Health Device Count
	Vnfabriccontrolplanefairhealthdevicecount int64 `json:"vnFabricControlPlaneFairHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Fair Health Device Count
	Multicastgoodhealthpercentage float64 `json:"multiCastGoodHealthPercentage,omitempty"` // Multicast Good Health Percentage
	Totalhealthdevicecount int64 `json:"totalHealthDeviceCount,omitempty"` // Total number of devices contributing to VN health
	Totalendpoints int64 `json:"totalEndpoints,omitempty"` // total number of Endpoints in the Virtual Network
	Pubsubsessiongoodhealthpercentage float64 `json:"pubsubSessionGoodHealthPercentage,omitempty"` // Pubsub Session Good Health Percentage
	Pubsubsessionfairhealthdevicecount int64 `json:"pubsubSessionFairHealthDeviceCount,omitempty"` // Pubsub Session Fair Health Device Count
	Vnfabriccontrolplanepoorhealthdevicecount int64 `json:"vnFabricControlPlanePoorHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Poor Health Device Count
	Pubsubsessionpoorhealthdevicecount int64 `json:"pubsubSessionPoorHealthDeviceCount,omitempty"` // Pubsub Session Poor Health Device Count
	Goodhealthpercentage float64 `json:"goodHealthPercentage,omitempty"` // Health percentage of sub category overallVnHealthPercentage
	Vnfabriccontrolplanegoodhealthdevicecount int64 `json:"vnFabricControlPlaneGoodHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Good Health Device Count
	Internetavailfairhealthdevicecount int64 `json:"internetAvailFairHealthDeviceCount,omitempty"` // Internet Availability Fair Health Device Count
	Vnservicesfairhealthdevicecount int64 `json:"vnServicesFairHealthDeviceCount,omitempty"` // Vn Services Fair Health Device Count
	Vnfabriccontrolplanenohealthdevicecount int64 `json:"vnFabricControlPlaneNoHealthDeviceCount,omitempty"` // Total number of no data devices in Virtual Network in Vn Fabric Control Plane category
	Name string `json:"name,omitempty"` // Full name of virtual network.
	Vnservicesnohealthdevicecount int64 `json:"vnServicesNoHealthDeviceCount,omitempty"` // Total number of no data devices in Virtual Network in Vn Services category
	Vnstatustotaldevicecount int64 `json:"vnStatusTotalDeviceCount,omitempty"` // Vn Status Total Device Count
	Internetavailtotaldevicecount int64 `json:"internetAvailTotalDeviceCount,omitempty"` // Internet Availability Total Device Count
	Firewallconntotaldevicecount int64 `json:"firewallConnTotalDeviceCount,omitempty"` // Firewall Connectivity Total Device Count
	Multicastgoodhealthdevicecount int64 `json:"multiCastGoodHealthDeviceCount,omitempty"` // Multicast Good Health Device Count
	Vnexittotaldevicecount int64 `json:"vnExitTotalDeviceCount,omitempty"` // Vn Exit Total Device Count
	Bgppeernohealthdevicecount int64 `json:"bgpPeerNoHealthDeviceCount,omitempty"` // Bgp Peer No Health Device Count
	Vnservicestotaldevicecount int64 `json:"vnServicesTotalDeviceCount,omitempty"` // Vn Services Total Device Count
	Multicastpoorhealthdevicecount int64 `json:"multiCastPoorHealthDeviceCount,omitempty"` // Multicast Poor Health Device Count
	Pubsubsessionnohealthdevicecount int64 `json:"pubsubSessionNoHealthDeviceCount,omitempty"` // Pubsub Session No Health Device Count
	Vnserviceshealthpercentage float64 `json:"vnServicesHealthPercentage,omitempty"` // Vn Services Health Percentage
	Internetavailnohealthdevicecount int64 `json:"internetAvailNoHealthDeviceCount,omitempty"` // Internet Availability No Health Device Count
	Internetavailpoorhealthdevicecount int64 `json:"internetAvailPoorHealthDeviceCount,omitempty"` // Internet Availability Poor Health Device Count
	Fairhealthdevicecount int64 `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count
	Vnstatusfairhealthdevicecount int64 `json:"vnStatusFairHealthDeviceCount,omitempty"` // Vn Status Fair Health Device Count
	Bgppeergoodhealthdevicecount int64 `json:"bgpPeerGoodHealthDeviceCount,omitempty"` // Bgp Peer Good Health Device Count
	Vninohealthdevicecount int64 `json:"vniNoHealthDeviceCount,omitempty"` // Vni No Health Device Count
	Internetavailgoodhealthdevicecount int64 `json:"internetAvailGoodHealthDeviceCount,omitempty"` // Internet Availability Good Health Device Count
	Multicastfairhealthdevicecount int64 `json:"multiCastFairHealthDeviceCount,omitempty"` // Multicast Fair Health Device Count
	Vnstatuspoorhealthdevicecount int64 `json:"vnStatusPoorHealthDeviceCount,omitempty"` // Vn Status Poor Health Device Count
	Vnfabriccontrolplanegoodhealthpercentage float64 `json:"vnFabricControlPlaneGoodHealthPercentage,omitempty"` // Vn Fabric Control Plane Good Health Percentage
	Associatedl3vn string `json:"associatedL3Vn,omitempty"` // associatedL3Vn.
	Vnigoodhealthdevicecount int64 `json:"vniGoodHealthDeviceCount,omitempty"` // Vni Good Health Device Count
	Poorhealthdevicecount int64 `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count
	Vnservicesgoodhealthdevicecount int64 `json:"vnServicesGoodHealthDeviceCount,omitempty"` // Vn Services Good Health Device Count
	Vnservicespoorhealthdevicecount int64 `json:"vnServicesPoorHealthDeviceCount,omitempty"` // Vn Services Poor Health Device Count
	Vnexitpoorhealthdevicecount int64 `json:"vnExitPoorHealthDeviceCount,omitempty"` // Vn Exit Poor Health Device Count
	Multicastnohealthdevicecount int64 `json:"multiCastNoHealthDeviceCount,omitempty"` // Multicast Session No Health Device Count
	Vnexithealthpercentage float64 `json:"vnExitHealthPercentage,omitempty"` // Vn Exit Health Percentage
	Bgppeerpoorhealthdevicecount int64 `json:"bgpPeerPoorHealthDeviceCount,omitempty"` // Bgp Peer Poor Health Device Count
	Firewallconngoodhealthpercentage float64 `json:"firewallConnGoodHealthPercentage,omitempty"` // Firewall Connectivity Good Health Percentage
	Vnfabriccontrolplanetotaldevicecount int64 `json:"vnFabricControlPlaneTotalDeviceCount,omitempty"` // Vn Fabric Control Plane Total Device Count
	Vnifairhealthdevicecount int64 `json:"vniFairHealthDeviceCount,omitempty"` // Vni Fair Health Device Count
	Vnexitgoodhealthdevicecount int64 `json:"vnExitGoodHealthDeviceCount,omitempty"` // Vn Exit Good Health Device Count
	Multicasttotaldevicecount int64 `json:"multiCastTotalDeviceCount,omitempty"` // Multicast Total Device Count
	Vnstatusnohealthdevicecount int64 `json:"vnStatusNoHealthDeviceCount,omitempty"` // Total number of no data devices in Virtual Network in Vn Status category
	Layer string `json:"layer,omitempty"` // vn layer.
	Goodhealthdevicecount int64 `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count
	Firewallconnfairhealthdevicecount int64 `json:"firewallConnFairHealthDeviceCount,omitempty"` // Firewall Connectivity Fair Health Device Count
	Pubsubsessiontotaldevicecount int64 `json:"pubsubSessionTotalDeviceCount,omitempty"` // Pubsub Session Total Device Count
	Internetavailgoodhealthpercentage float64 `json:"internetAvailGoodHealthPercentage,omitempty"` // Internet Availability Good Health Percentage
	Vnitotaldevicecount int64 `json:"vniTotalDeviceCount,omitempty"` // Vni Total Device Count
	Firewallconnnohealthdevicecount int64 `json:"firewallConnNoHealthDeviceCount,omitempty"` // Firewall Connectivity No Health Device Count
	Bgppeertotaldevicecount int64 `json:"bgpPeerTotalDeviceCount,omitempty"` // Bgp Peer Total Device Count
	Bgppeerfairhealthdevicecount int64 `json:"bgpPeerFairHealthDeviceCount,omitempty"` // Bgp Peer Fair Health Device Count
	Vnstatushealthpercentage float64 `json:"vnStatusHealthPercentage,omitempty"` // Vn Status Health Percentage
	Firewallconnpoorhealthdevicecount int64 `json:"firewallConnPoorHealthDeviceCount,omitempty"` // Firewall Connectivity Poor Health Device Count
	Vnexitnohealthdevicecount int64 `json:"vnExitNoHealthDeviceCount,omitempty"` // Total number of no data devices in Virtual Network in Vn Exit category
	Nohealthdevicecount int64 `json:"noHealthDeviceCount,omitempty"` // Total number of no data devices in Virtual Network
	Id string `json:"id,omitempty"` // Unique uuid of the virtual network.
	Pubsubsessiongoodhealthdevicecount int64 `json:"pubsubSessionGoodHealthDeviceCount,omitempty"` // Pubsub Session Good Health Device Count
	Vlan string `json:"vlan,omitempty"` // vlan.
	Vnid string `json:"vnid,omitempty"` // vnid.
	Vnigoodhealthpercentage float64 `json:"vniGoodHealthPercentage,omitempty"` // Vni Good Health Percentage
	Vnexitfairhealthdevicecount int64 `json:"vnExitFairHealthDeviceCount,omitempty"` // Vn Exit Fair Health Device Count
}

// VirtualNetworkSummary represents the VirtualNetworkSummary schema from the OpenAPI specification
type VirtualNetworkSummary struct {
	Response VirtualNetworkSummary `json:"response,omitempty"` // Summary of virtual networks with health info.
	Version string `json:"version,omitempty"` // The version of the response
}

// VirtualNetworkSummaryTrend represents the VirtualNetworkSummaryTrend schema from the OpenAPI specification
type VirtualNetworkSummaryTrend struct {
	Page TimeBasedPagination `json:"page,omitempty"`
	Response []VirtualNetworkTrend `json:"response,omitempty"` // Successfully return list of Virtual networks health trend data
	Version string `json:"version,omitempty"`
}

// VirtualNetworkAttributeValuePair represents the VirtualNetworkAttributeValuePair schema from the OpenAPI specification
type VirtualNetworkAttributeValuePair struct {
	Name string `json:"name,omitempty"` // Supported filter attributes related to clients
	Value map[string]interface{} `json:"value,omitempty"` // This is the value of the attribute requested for analytics API.
}

// VirtualNetworksSummaries represents the VirtualNetworksSummaries schema from the OpenAPI specification
type VirtualNetworksSummaries struct {
	Page PaginationResponse `json:"page,omitempty"`
	Response []VirtualNetworkSummary `json:"response,omitempty"` // Successfully returned list of Virtual Networks
	Version string `json:"version,omitempty"` // The version of the response
}

// CountIntegerResponseresponse represents the CountIntegerResponseresponse schema from the OpenAPI specification
type CountIntegerResponseresponse struct {
	Count int64 `json:"count,omitempty"` // The total number of records related to the resource
}

// VirtualNetworkTrend represents the VirtualNetworkTrend schema from the OpenAPI specification
type VirtualNetworkTrend struct {
	Attributes []VirtualNetworkAttributeValuePair `json:"attributes,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"` // For trend API timestamp is the default groupBy attribute.
}

// CountIntegerResponse represents the CountIntegerResponse schema from the OpenAPI specification
type CountIntegerResponse struct {
	Response CountIntegerResponseresponse `json:"response,omitempty"`
	Version string `json:"version,omitempty"` // The version of the response
}

// ErrorObject represents the ErrorObject schema from the OpenAPI specification
type ErrorObject struct {
	Message string `json:"message,omitempty"` // Brief message about the error condition
	Detail string `json:"detail,omitempty"` // A more detailed explanation of the error condition the parameter and its value, that caused the condition and why it caused it.
	Errorcode int `json:"errorCode,omitempty"` // Application specific error code returned by the server
}
