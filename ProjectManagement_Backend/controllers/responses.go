package controllers

// jwtTokenResponse represents the response with a JWT token
// swagger:response jwtTokenResponse
type jwtTokenResponse struct {
	// in: body
	Body struct {
		JwtToken string `json:"jwtToken"`
	}
}

// errorResponse represents an error response
// swagger:response errorResponse
type errorResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// emptyResponse represents an empty response
// swagger:response emptyResponse
type emptyResponse struct{}
