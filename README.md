# aws-ip

The package to get [AWS IP address ranges](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html).

## SYNOPSIS

```go
package main

import (
    awsip "github.com/shmokmt/aws-ip"
)

func main() {
    prefixes := awsip.Query().Service("CODEBUILD").NetworkBorderGroup().Select()
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
