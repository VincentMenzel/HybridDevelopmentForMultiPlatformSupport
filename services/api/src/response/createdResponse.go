package response

type CreatedResponse struct {
	Id int64 `json:"id"`
}

func NewCreatedResponse(id int64) *CreatedResponse {
	return &CreatedResponse{Id: id}
}
