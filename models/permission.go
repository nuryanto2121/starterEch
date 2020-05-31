package models

// Permission :
type Permission struct {
	CompanyID   int                `json:"company_id,omitempty"`
	CompanyName string             `json:"company_name,omitempty"`
	IsAccess    bool               `json:"is_access,omitempty"`
	IsDefault   bool               `json:"is_default,omitempty"`
	DataBranch  []PermissionBranch `json:"data_branch"`
}

// PermissionBranch :
type PermissionBranch struct {
	BranchID   int    `json:"branch_id,omitempty"`
	BranchName string `json:"branch_name,omitempty"`
	IsAccess   bool   `json:"is_access,omitempty"`
}
