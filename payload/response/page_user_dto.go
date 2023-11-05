package response

type PageUserDto struct {
	Items        []UserDto `json:"items"`
	Page         int64       `json:"page"`
	Size         int64       `json:"size"`
	MaxPage      int64       `json:"max_page"`
	TotalPages   int64       `json:"total_pages"`
	Total        int64       `json:"total"`
	Last         bool        `json:"last"`
	First        bool        `json:"first"`
}