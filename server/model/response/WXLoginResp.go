package response
type WXloginResp struct {
	OpenId string 		`json:"open_id"`
	SessionKey string	`json:"session_key"`
	UnionId string		`json:"union_id"`
	ErrCode int			`json:"err_code"`
	ErrMsg string		`json:"err_msg"`
}
