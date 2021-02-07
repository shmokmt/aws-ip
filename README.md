# aws-ip

![Test](https://github.com/shmokmt/aws-ip/workflows/Test/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/shmokmt/aws-ip.svg)](https://pkg.go.dev/github.com/shmokmt/aws-ip)

The package to get [AWS IP address ranges](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html).

## SYNOPSIS

```go
package main

import (
    awsip "github.com/shmokmt/aws-ip"
)

func main() {
    prefixes := awsip.Query().Service("CODEBUILD").Select()
    for _, p := range prefixes {
        fmt.Println(p.IPPrefix)
        fmt.Println(p.Region)
        fmt.Println(p.Service)
        fmt.Println(p.NetworkBorderGroup)
    }
}
```

### LICENSE

MIT
