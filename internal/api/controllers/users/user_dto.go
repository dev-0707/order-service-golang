// Package dto provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package users

const (
	Authorization_TokenScopes = "Authorization_Token.Scopes"
)

// UsersUser defines model for users.User.
type UsersUser struct {
	Firstname *string        `json:"firstname,omitempty"`
	Id        *int           `json:"id,omitempty"`
	Lastname  *string        `json:"lastname,omitempty"`
	Role      *UsersUserRole `json:"role,omitempty"`
	Username  *string        `json:"username,omitempty"`
}

// UsersUserRole defines model for users.UserRole.
type UsersUserRole struct {
	RoleName *string `json:"role_name,omitempty"`
}

// GetApiUsersParams defines parameters for GetApiUsers.
type GetApiUsersParams struct {
	// Username Username
	Username *string `form:"username,omitempty" json:"username,omitempty"`

	// Firstname Firstname
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty"`

	// Lastname Lastname
	Lastname *string `form:"lastname,omitempty" json:"lastname,omitempty"`
}