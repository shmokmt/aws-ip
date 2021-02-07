// The package to get AWS IP ranges.
package aws_ip

// IPRanges has various information related to `ip_prefix`.
type IPRange struct {
	IPPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

// IPRanges is a slice of IPRange.
type IPRanges []*IPRange
