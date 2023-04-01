package model

import (
	"time"

	"github.com/google/uuid"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type M_Response_session struct {
	SessionID        uuid.UUID `json:"id" form:"id" query:"id"`
	Username         string    `json:"username" form:"username" query:"username"`
	AccessToken      string    `json:"access_token" form:"access_token" query:"access_token"`
	Access_ExpiresAt time.Time `json:"acces_expires_at" form:"access_expires_at" query:"access_expires_at"`
}

type ResponsesData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Keyword string      `json:"keyword"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	OrderBy string      `json:"order_by"`
	Data    interface{} `json:"data"`
}
