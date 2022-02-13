package inventory

type (
	PaginatedResponse struct {
		PaginatedRequest
		Items Responses `json:"items"`
	}
)
