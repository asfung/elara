package models

// Request DTOs
type AddRoleRequest struct {
	Name        string `validate:"required,name"`
	Description string `validate:"omitempty,description"`
}
type UpdateRoleRequest struct {
	ID          uint   `validate:"required"`
	Name        string `validate:"required,name"`
	Description string `validate:"omitempty,description"`
}

// Response DTOs
type RoleResponse struct {
}

// Entity -> Response
