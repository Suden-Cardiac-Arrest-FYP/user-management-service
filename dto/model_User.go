package dto

type User struct {
	UserId         string ` json:"UserId" `
	FirstName      string ` json:"FirstName" `
	LastName       string ` json:"LastName" `
	Email          string ` json:"Email" `
	RoleId         string ` json:"RoleId" `
	RoleName       string ` json:"RoleName" `
	OrganizationId string ` json:"organizationId" `
	Organization   string ` json:"Organization" `
	Deleted        bool   `json:"deleted"`
}
