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
