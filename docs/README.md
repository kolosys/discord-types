# discord-api-types (Go)

[![Go Reference](https://pkg.go.dev/badge/github.com/kolosys/discord-types.svg)](https://pkg.go.dev/github.com/kolosys/discord-types)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolosys/discord-types)](https://goreportcard.com/report/github.com/kolosys/discord-types)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Discord API types for Go that are kept up to date for use in Discord bot library creation.

This library provides comprehensive Go type definitions for the Discord API v10, including:

- **Payloads**: User, Guild, Channel, Message, Auto Moderation, Monetization, Polls, Soundboard, and all Discord structures
- **REST API**: Request/response types and route constants
- **Gateway**: Complete WebSocket event types, dispatch events, and payloads (70+ events)
- **Voice**: Voice Gateway v8 with DAVE protocol support and E2E encryption
- **Utilities**: Helper functions and constants

## Features

✨ **Complete API v10 Coverage** - All Discord API v10 types including latest features  
🛡️ **Auto Moderation** - Full support for Discord's auto-moderation system  
📊 **Polls** - Complete poll creation, voting, and management  
🔊 **Soundboard** - Guild soundboard sound management  
💰 **Monetization** - SKUs, subscriptions, and entitlements  
📅 **Scheduled Events** - Event scheduling with recurrence support  
🧵 **Threads** - Complete thread lifecycle management  
🎭 **Stage Instances** - Voice stage management  
🔐 **Voice Gateway v8** - Complete voice support with DAVE E2E encryption  
⚡ **70+ Gateway Events** - All Discord WebSocket events  
🤖 **Complete RPC API** - Rich Presence, voice control, and Discord client interaction  
🔒 **Type Safety** - Compile-time validation with Go interfaces

## Installation

```bash
go get github.com/kolosys/discord-types
```

## Usage

### Basic Types

```go
package main

import (
    "fmt"
    "github.com/kolosys/discord-types/discord"
    "github.com/kolosys/discord-types/payloads"
    "github.com/kolosys/discord-types/rest"
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

    // Using REST routes
    userRoute := rest.Routes.User(userID)
    fmt.Printf("User route: %s\n", userRoute)
}

func StringPtr(s string) *string {
    return &s
}
```

### Gateway Events

```go
package main

import (
    "fmt"
    "github.com/kolosys/discord-types/gateway"
)

func handleGatewayEvent(event gateway.GatewayReceivePayload) {
    switch e := event.(type) {
    case gateway.ReadyDispatch:
        fmt.Printf("Bot connected as: %s\n", e.D.User.Username)

    case gateway.MessageCreateDispatch:
        fmt.Printf("Message from %s: %s\n",
            e.D.Author.Username,
            e.D.Content)

    case gateway.GuildCreateDispatch:
        fmt.Printf("Guild: %s (%d members)\n",
            e.D.Name,
            e.D.MemberCount)

    case gateway.AutoModerationActionExecutionDispatch:
        fmt.Printf("Auto-mod action: Rule %s triggered by user %s\n",
            e.D.RuleID,
            e.D.UserID)

    case gateway.MessagePollVoteAddDispatch:
        fmt.Printf("Poll vote: User %s voted for answer %d\n",
            e.D.UserID,
            e.D.AnswerID)
    }
}
```

### Voice Gateway

```go
package main

import (
    "github.com/kolosys/discord-types/voice"
    "github.com/kolosys/discord-types/discord"
)

func handleVoiceEvent(event voice.VoiceReceivePayload) {
    switch e := event.(type) {
    case voice.VoiceReady:
        fmt.Printf("Voice ready: SSRC=%d, IP=%s, Port=%d\n",
            e.D.SSRC, e.D.IP, e.D.Port)

    case voice.VoiceSpeaking:
        fmt.Printf("User %s is speaking (SSRC: %d)\n",
            e.D.UserID, e.D.SSRC)

    case voice.VoiceClientsConnect:
        fmt.Printf("Clients connected: %v\n", e.D.UserIDs)
    }
}

func sendVoiceIdentify(serverID, userID discord.Snowflake, sessionID, token string) voice.VoiceIdentify {
    return voice.VoiceIdentify{
        Op: voice.VoiceOpcodeIdentify,
        D: voice.VoiceIdentifyData{
            ServerID:  serverID,
            UserID:    userID,
            SessionID: sessionID,
            Token:     token,
        },
    }
}
```

### RPC (Rich Presence Client)

```go
package main

import (
    "github.com/kolosys/discord-types/rpc"
    "github.com/kolosys/discord-types/discord"
)

func sendRPCCommand(clientID discord.Snowflake) rpc.RPCCommandPayload {
    return rpc.RPCCommandPayload{
        Cmd: rpc.RPCCommandSetActivity,
        Args: map[string]interface{}{
            "activity": map[string]interface{}{
                "details": "Playing a game",
                "state":   "In a match",
                "timestamps": map[string]interface{}{
                    "start": 1234567890,
                },
            },
        },
        Nonce: "unique-nonce-123",
    }
}

func handleRPCEvent(payload rpc.RPCReceivePayload) {
    switch p := payload.(type) {
    case rpc.RPCEventPayload:
        switch p.Evt {
        case rpc.RPCEventReady:
            fmt.Println("RPC client ready")
        case rpc.RPCEventVoiceSettingsUpdate:
            fmt.Println("Voice settings updated")
        case rpc.RPCEventActivityJoinRequest:
            fmt.Println("Activity join requested")
        }
    case rpc.RPCErrorPayload:
        fmt.Printf("RPC Error: %s (Code: %d)\n", p.Data.Message, p.Data.Code)
    }
}
```

## Package Structure

- `discord-types` - Core types (Snowflake, Permissions, etc.)
- `discord-types/payloads` - Complete Discord object structures including:
  - Users, Guilds, Channels, Messages
  - Auto Moderation rules and actions
  - Monetization (SKUs, Subscriptions, Entitlements)
  - Polls and voting
  - Soundboard sounds
  - Guild Scheduled Events with recurrence
  - Audit Logs
  - OAuth2 scopes
  - Voice states and regions
  - Templates and Teams
  - Stage Instances
- `discord-types/rest` - REST API routes and request/response types
- `discord-types/gateway` - Complete WebSocket support including:
  - 70+ dispatch event types
  - Gateway connection management
  - Comprehensive event data structures
  - Send/receive payload interfaces
- `discord-types/voice` - Voice Gateway v8 with:
  - DAVE protocol support
  - E2E encryption
  - Voice connection management
  - Speaking state management
- `discord-types/rpc` - Rich Presence Client with:
  - 60+ RPC commands for Discord client interaction
  - Real-time event subscriptions
  - Rich Presence activity management
  - Voice settings and device control
  - OAuth2 authentication support
- `discord-types/utils` - Utility functions and helpers

## Version Compatibility

This library tracks Discord API v10. This is the initial v1.0.0 release with complete Discord API v10 support:

- v1.0.0+ - Complete Discord API v10 coverage

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Documentation

For complete Discord API documentation, see the [official Discord API docs](https://discord.com/developers/docs/intro).
