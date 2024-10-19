package model

type GetUserInfoByIDReq struct {
	UserID string `json:"user_id"`
}

type GetUserInfoByIDResp struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
