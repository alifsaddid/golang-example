package request

type RoleRequest struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}
