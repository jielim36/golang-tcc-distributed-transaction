package models

type User struct {
	UserId             uint   `json:"user_id" gorm:"primaryKey"`
	UserName           string `json:"username"`
	SubscriptionPlanId uint   `json:"subscription_plan_id"`
	AccountId          uint   `json:"account_id"`
}
