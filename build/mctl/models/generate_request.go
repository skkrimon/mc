package models

type GenerateRequest struct {
	Username string `json:"username" binding:"required"`
}
