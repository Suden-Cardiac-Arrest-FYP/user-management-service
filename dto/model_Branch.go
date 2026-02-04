package dto

type Branch struct {
	BranchId    string ` json:"BranchId" `
	Name        string ` json:"Name" `
	Address     string ` json:"Address" `
	PhoneNumber string ` json:"PhoneNumber" `
	Email       string ` json:"Email" `
	ManagerName string ` json:"ManagerName" `
	Logo        string ` json:"Logo" `
}
