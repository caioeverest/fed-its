package model

type MethodProvider struct {
	MethodID   uint `gorm:"not null;index" json:"method_id"`
	Method     Method
	ProviderID uint `gorm:"not null;index" json:"provider_id"`
	Provider   Provider
}
