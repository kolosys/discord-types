package v10

// Team represents a Discord team.
// Reference: https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	// ID is the team's id.
	ID Snowflake `json:"id"`

	// Name is the team's name.
	Name string `json:"name"`

	// Icon is the team's icon hash.
	Icon *string `json:"icon"`

	// OwnerUserID is the user id of the team owner.
	OwnerUserID Snowflake `json:"owner_user_id"`

	// Members are the team members.
	Members []TeamMember `json:"members"`
}

// TeamMember represents a member of a Discord team.
// Reference: https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	// MembershipState is the user's membership state on the team.
	MembershipState TeamMembershipState `json:"membership_state"`

	// Role is the user's role on the team.
	Role TeamMemberRole `json:"role"`

	// TeamID is the id of the parent team of which they are a member.
	TeamID Snowflake `json:"team_id"`

	// User is the avatar, discriminator, id, and username of the user.
	User PartialUser `json:"user"`
}

// TeamMembershipState represents the membership state of a team member.
// Reference: https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
type TeamMembershipState int

// Team membership state constants
const (
	// TeamMembershipStateInvited indicates the user was invited to the team.
	TeamMembershipStateInvited TeamMembershipState = iota + 1

	// TeamMembershipStateAccepted indicates the user has accepted the invitation to the team.
	TeamMembershipStateAccepted
)

// TeamMemberRole represents the role of a team member.
// Reference: https://discord.com/developers/docs/topics/teams#data-models-team-member-role-types
type TeamMemberRole string

// Team member role constants
const (
	// TeamMemberRoleAdmin indicates the team member is an admin.
	TeamMemberRoleAdmin TeamMemberRole = "admin"

	// TeamMemberRoleDeveloper indicates the team member is a developer.
	TeamMemberRoleDeveloper TeamMemberRole = "developer"

	// TeamMemberRoleReadOnly indicates the team member has read-only access.
	TeamMemberRoleReadOnly TeamMemberRole = "read_only"
)
