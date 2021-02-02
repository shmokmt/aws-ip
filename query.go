package aws_ip

type querier struct {
	Ranges IPRanges
}

// Query creates querier object.
func Query() *querier {
	return &querier{
		Ranges: *ipRanges,
	}
}

func (q *querier) filter(fn func(*IPRange) bool) (result IPRanges) {
	for _, v := range q.Ranges {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Region filters IPRanges by region's name.
func (q *querier) Region(region string) *querier {
	fn := func(ipr *IPRange) bool {
		return ipr.Region == region
	}
	q.Ranges = q.filter(fn)
	return q
}

// Service filters IPRanges by service's name.
func (q *querier) Service(service string) *querier {
	fn := func(ipr *IPRange) bool {
		return ipr.Service == service
	}
	q.Ranges = q.filter(fn)
	return q
}

// NetworkBorderGroup filters IPRanges by network's name.
func (q *querier) NetworkBorderGroup(network string) *querier {
	fn := func(ipr *IPRange) bool {
		return ipr.NetworkBorderGroup == network
	}
	q.Ranges = q.filter(fn)
	return q
}

// Select returns IPRanges.
func (q *querier) Select() IPRanges {
	return q.Ranges
}
