package model

type Rule struct {
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	RoleId         int64    `gorm:"column:role_id" json:"role_id"`
	PermissionId         int64    `gorm:"column:permission_id" json:"permission_id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
}

// TableName sets the insert table name for this struct type
func (r *Rule) TableName() string {
	return "rule"
}

