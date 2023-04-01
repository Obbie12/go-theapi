package model

type M_login_input struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type M_login_user struct {
	Displayname string `json:"nama" form:"nama" query:"nama"`
	Password    string `json:"password" form:"password" query:"password"`
	Username    string `json:"username" form:"username" query:"username"`
}

type M_login_result struct {
	Displayname string `json:"nama" form:"nama" query:"nama"`
	Username    string `json:"username" form:"username" query:"username"`
	AccessToken string `json:"access_token" form:"access_token" query:"access_token"`
}
