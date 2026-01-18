package auth

type AccessRequest struct {
	Access string `param:"access" validate:"required"`
}

type AccessResponse struct {
	Access string `json:"access"`
}
