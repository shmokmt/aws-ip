// Tha package to get AWS IP ranges.
package aws_ip

// Querier executes query.
type Querier struct {
	Ranges IPRanges
}

// Query creates Querier.
func Query() *Querier {
	return &Querier{
		Ranges: *ipRanges,
	}
}

func (q *Querier) filter(fn func(*IPRange) bool) (result IPRanges) {
	for _, v := range q.Ranges {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Region filters IPRanges by region's name.
func (q *Querier) Region(region string) *Querier {
	fn := func(ipr *IPRange) bool {
		return ipr.Region == region
	}
	q.Ranges = q.filter(fn)
	return q
}

// Service filters IPRanges by service's name.
func (q *Querier) Service(service string) *Querier {
	fn := func(ipr *IPRange) bool {
		return ipr.Service == service
	}
	q.Ranges = q.filter(fn)
	return q
}

// NetworkBorderGroup filters IPRanges by network's name.
func (q *Querier) NetworkBorderGroup(network string) *Querier {
	fn := func(ipr *IPRange) bool {
		return ipr.NetworkBorderGroup == network
	}
	q.Ranges = q.filter(fn)
	return q
}

// Select returns IPRanges.
func (q *Querier) Select() IPRanges {
	return q.Ranges
}
