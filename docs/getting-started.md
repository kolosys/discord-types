# Getting Started

This guide will help you get up and running quickly with this package.

## Installation

```bash
go get github.com/kolosys/discord-types
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "fmt"
    
    "github.com/kolosys/discord-types/discord"
    "github.com/kolosys/discord-types/payloads"
)

func main() {
    // Using core types
    var userID discord.Snowflake = "123456789012345678"

    // Using payload types
    user := payloads.User{
        ID:            userID,
        Username:      "example",
        Discriminator: "0001",
        GlobalName:    StringPtr("Example User"),
    }

    fmt.Printf("User: %s (%s)\n", user.Username, user.ID)
}

func StringPtr(s string) *string {
    return &s
}
```

## Next Steps

- [API Reference](api-reference.md)
- [Examples](examples.md)
- [GitHub Repository](https://github.com/kolosys/ion)
