package response

type ErrorResp struct {
	Code    string `json:"code "`
	Message string `json:"message"`
}
type CreateResponse struct {
	Id string `json:"id"`
}
