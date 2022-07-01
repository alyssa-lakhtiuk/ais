package controller

type ErrorMessage struct {
	Message string `json:"error"`
}

//var (
//	bearerTokenErr  = errors.New("could not authorize the request. make sure the request has an Authorization header with a bearer token")
//	unauthorizedErr = errors.New("unauthorized")
//	invalidTokenErr = errors.New("token has been expired or revoked")
//	internalError   = errors.New(internalServerErrorMessage)
//)
