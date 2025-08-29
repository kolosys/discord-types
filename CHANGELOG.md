# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),  
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [1.0.0] - 2025-08-29

### Added

#### Discord API v10 Coverage

- 100% coverage of all Discord API v10 types, endpoints, and events.

#### REST Package (`rest`)

- **8 specialized modules (~4,200 LOC)** with comprehensive Discord API coverage:
  - **Routes**: Complete REST routes, CDN builders, OAuth2 helpers, image validation.
  - **Channel Management**: Messages, reactions, threads, permissions, pins, typing, follows, group DMs.
  - **Guild Management**: Members, roles, settings, integrations, templates, audit logs, discovery.
  - **Interactions**: Commands, permissions, callbacks, modals, autocomplete.
  - **Specialized Features**: Emojis, stickers, soundboard, polls, users, invites, apps.
  - **Query Builders**: Type-safe query construction, pagination, filtering, encoding.
  - **Types**: Request/response definitions, error handling, type aliases.
  - **Error Codes**: 230+ categorized Discord error codes with descriptions.

#### Payloads

- **Auto Moderation**: Rules, triggers, actions, full lifecycle.
- **Monetization**: Entitlements, SKUs, subscriptions.
- **Polls**: Creation, voting, results.
- **Soundboard**: Sound management and playback.
- **Guild Scheduled Events**: Recurrence rules, statuses, metadata.
- **Audit Logs**: Full event and change key support.
- **OAuth2**: Scope definitions.
- **Voice**: States and regions.
- **Templates**: Guild templates with roles, channels, overwrites.
- **Teams**: Team and member management.
- **Stage Instances**: Full lifecycle and privacy levels.

#### Gateway

- **70+ Dispatch Events**: Coverage of all events (guild, channels, messages, polls, integrations, etc.).
- **Connection Management**: Heartbeat, resume, compression, URL queries.
- **Sendable Payloads**: Guild member requests, soundboard, voice state, presence updates, resume.
- **Type Safety**: Union interfaces with compile-time validation.

#### Voice Gateway v8

- **All opcodes, close codes, and encryption modes**.
- **DAVE Protocol support** with MLS and full E2E encryption.
- **Connection Management**: Sessions, heartbeats, UDP, SSRC, key management.

#### Utils Package (`utils`)

- Pointer helpers, snowflake utilities, permission checks.
- Mention and timestamp builders.
- CDN, color, and type conversion helpers.

#### RPC Package (`rpc`)

- **60+ RPC commands** and full event system.
- OAuth2 integration, Rich Presence, voice settings, device certification.
- Guild/channel access, relationships, invitations, error codes, client events.

#### OAuth2 RPC

- RPC scopes, application RPC origins/flags.
- RPC activities, notifications, and voice settings.

#### Type Safety

- **Interface marker methods** for compile-time validation with union-style type safety.
- **Generic payload types** (`VoiceDataPayload[T]`).

### Fixed

- Removed duplicate type declarations and circular dependencies.
- All packages compile with zero lint errors.
- REST type coverage expanded to complete Discord API v10 coverage.

### Changed

- **Package structure** redesigned for modularity.
- **Documentation** rewritten with examples.
- **Type names** standardized to Go conventions.
- **REST architecture** split into specialized modules.

### Performance

- Zero runtime cost marker methods.
- Optimized struct layouts and memory usage.
- Fast compilation with modular dependencies.

### Documentation

- Complete README with installation and usage.
- Full API documentation with Discord links.
- Usage examples for all major features.
- Detailed CHANGELOG.

### Testing & Quality

- Verified builds across all packages.
- Zero critical lint issues.
- Go module compliance with proper dependency management.

### Production-Ready

✅ Complete type safety across all APIs  
✅ Complete Discord API v10 feature coverage  
✅ 230+ categorized error codes  
✅ Developer-friendly naming and docs  
✅ Optimized performance and compile-time validation
