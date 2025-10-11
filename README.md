# discord-types

[![Go Reference](https://pkg.go.dev/badge/github.com/kolosys/discord-types.svg)](https://pkg.go.dev/github.com/kolosys/discord-types)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolosys/discord-types)](https://goreportcard.com/report/github.com/kolosys/discord-types)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://go.dev)

A comprehensive Go library providing type definitions for Discord API v10. This library offers strongly-typed structs, constants, and utilities for building Discord bots and applications in Go.

## ✨ Features

- **Complete Discord API v10 Coverage** - All Discord API types, structures, and constants
- **Strongly Typed** - Full type safety with comprehensive struct definitions
- **Zero Dependencies** - Pure Go implementation with no external dependencies
- **Well Documented** - Extensive documentation with Discord API references
- **Permission System** - Complete permission flags and bitwise operations
- **Message Formatting** - Built-in regex patterns for Discord message parsing
- **Interaction Support** - Full support for slash commands, components, and modals
- **Gateway Events** - Complete gateway event structures and dispatch data
- **REST API Types** - All REST endpoint request/response structures

## 🚀 Quick Start

### Installation

```bash
go get github.com/kolosys/discord-types/v10
```

### Basic Usage

```go
package main

import (
    "fmt"
    discord "github.com/kolosys/discord-types/v10"
)

func main() {
    // Create a message embed
    embed := discord.Embed{
        Title:       "Hello, Discord!",
        Description: "This is a sample embed",
        Color:       discord.ColorBlurple,
        Fields: []discord.EmbedField{
            {
                Name:   "Field 1",
                Value:  "Some value",
                Inline: true,
            },
        },
    }

    // Create a message with components
    message := discord.Message{
        Content: "Hello from Go!",
        Embeds:  []discord.Embed{embed},
        Components: []discord.MessageComponent{
            discord.ActionRow{
                Components: []discord.MessageComponent{
                    discord.Button{
                        Type:     discord.ComponentTypeButton,
                        Style:    discord.ButtonStylePrimary,
                        Label:    "Click me!",
                        CustomID: "button_1",
                    },
                },
            },
        },
    }

    fmt.Printf("Message: %+v\n", message)
}
```

## 📚 Key Components

### Core Types

- **Snowflake** - Discord's unique identifier system
- **User, Guild, Channel** - Core Discord entities
- **Message, Embed** - Message structures and rich embeds
- **Interaction** - Slash commands, buttons, select menus, modals
- **Permissions** - Complete permission system with bitwise operations

### Message Components

```go
// Create interactive components
button := discord.Button{
    Type:     discord.ComponentTypeButton,
    Style:    discord.ButtonStyleSecondary,
    Label:    "Click me",
    CustomID: "my_button",
    Emoji: &discord.PartialEmoji{
        Name: "👋",
    },
}

selectMenu := discord.StringSelectMenu{
    Type:        discord.ComponentTypeStringSelect,
    CustomID:    "my_select",
    Placeholder: "Choose an option",
    Options: []discord.SelectOption{
        {
            Label:       "Option 1",
            Value:       "opt1",
            Description: "First option",
        },
    },
}
```

### Slash Commands

```go
// Define a slash command
command := discord.ApplicationCommand{
    Name:        "hello",
    Description: "Say hello to someone",
    Options: []discord.ApplicationCommandOption{
        {
            Type:        discord.ApplicationCommandOptionTypeUser,
            Name:        "user",
            Description: "User to greet",
            Required:    true,
        },
    },
}
```

### Permission Handling

```go
// Check and manipulate permissions
perms := discord.PermissionFlagsBits.SendMessages |
          discord.PermissionFlagsBits.EmbedLinks

// Check if user has specific permission
hasPermission := (userPermissions & discord.PermissionFlagsBits.Administrator) != 0

// Create permission overwrite
overwrite := discord.PermissionOverwrite{
    ID:   userID,
    Type: discord.PermissionOverwriteTypeMember,
    Allow: discord.PermissionFlagsBits.SendMessages.String(),
    Deny:  discord.PermissionFlagsBits.MentionEveryone.String(),
}
```

### Message Formatting

```go
// Parse Discord message formatting
userMention := discord.FormattingPatterns.User.FindString("<@123456789>")
channelMention := discord.FormattingPatterns.Channel.FindString("<#987654321>")
timestamp := discord.FormattingPatterns.Timestamp.FindString("<t:1640995200:F>")
```

## 🎯 Common Use Cases

### Building a Discord Bot

```go
// Handle slash command interaction
func handleSlashCommand(interaction discord.Interaction) {
    if interaction.Data.Name == "ping" {
        response := discord.InteractionResponse{
            Type: discord.InteractionResponseTypeChannelMessageWithSource,
            Data: &discord.InteractionResponseData{
                Content: "Pong! 🏓",
                Flags:   discord.MessageFlagEphemeral,
            },
        }
        // Send response...
    }
}

// Handle button interaction
func handleButton(interaction discord.Interaction) {
    if interaction.Data.CustomID == "my_button" {
        response := discord.InteractionResponse{
            Type: discord.InteractionResponseTypeUpdateMessage,
            Data: &discord.InteractionResponseData{
                Content: "Button clicked!",
                Components: []discord.MessageComponent{}, // Remove components
            },
        }
        // Send response...
    }
}
```

### Webhook Integration

```go
// Create webhook message
webhookMessage := discord.ExecuteWebhookParams{
    Content:   "Message from webhook",
    Username:  "My Bot",
    AvatarURL: "https://example.com/avatar.png",
    Embeds: []discord.Embed{
        {
            Title: "Webhook Message",
            Color: discord.ColorGreen,
        },
    },
}
```

## 📖 Documentation

### Type Categories

- **`application.go`** - Application and OAuth2 structures
- **`channel.go`** - Channel types and properties
- **`components.go`** - Message components (buttons, select menus, modals)
- **`embed.go`** - Rich embed structures
- **`guild.go`** - Guild (server) related types
- **`interactions.go`** - Slash commands and interaction handling
- **`message.go`** - Message structures and attachments
- **`permission.go`** - Permission flags and calculations
- **`user.go`** - User and member structures
- **`webhook.go`** - Webhook types and execution

### Constants and Limits

The library includes all Discord API limits and constants:

```go
discord.MaxMessageLength          // 2000 characters
discord.MaxEmbedLength           // 6000 characters total
discord.MaxEmbedFields           // 25 fields per embed
discord.MaxEmbedsPerMessage      // 10 embeds per message
discord.MaxButtonsPerActionRow   // 5 buttons per row
discord.MaxActionRowsPerMessage  // 5 action rows per message
```

### Color Constants

```go
discord.ColorBlurple         // Discord's signature color
discord.ColorGreen          // Success green
discord.ColorRed            // Error red
discord.ColorYellow         // Warning yellow
// ... and many more
```

## 🔗 Related Projects

This library is designed to work seamlessly with Discord bot frameworks and HTTP clients:

- Use with [discordgo](https://github.com/bwmarrin/discordgo)
- Integrate with custom HTTP clients
- Perfect for webhook implementations
- Ideal for Discord application development

## 🤝 Contributing

We welcome contributions! Please feel free to submit issues, feature requests, or pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Discord Developer Documentation](https://discord.com/developers/docs)
- [Go Package Documentation](https://pkg.go.dev/github.com/kolosys/discord-types/v10)
- [Discord API v10 Reference](https://discord.com/developers/docs/reference)

---

Built with ❤️ for the Discord developer community
