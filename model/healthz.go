package model

// A HealthzResponse expresses health check message.
// 大文字で構造体変数宣言
type HealthzResponse struct{
	Message string `json:"message"`

}
