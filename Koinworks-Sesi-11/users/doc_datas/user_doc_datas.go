package doc_datas

import "time"

type UserRegisterResponse struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginResponse struct {
	Token string `json:"token" example:"kjnaoreknvrogndklanoek"`
}
