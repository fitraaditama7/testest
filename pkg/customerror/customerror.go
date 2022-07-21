package customerror

import "test/pkg/error"

// ErrInternalServerError Error Status for Internal Server Error
var ErrInternalServerError = error.New("500", "", "ERR_INTERNAL_SERVER_ERROR", "Internal Server Error")

// ErrNoContent Error Status for No Content
var ErrNoContent = error.New("204", "", "ERR_NO_CONTENT", "No Content")

// ErrInvalidBodyRequest Error Status for Invalid Body Request
var ErrInvalidBodyRequest = error.New("400", "", "ERR_INVALID_BODY_REQUEST", "Invalid Body Request")

// ErrDataNotFound Error Status for Data Not Found
var ErrDataNotFound = error.New("404", "", "ERR_DATA_NOT_FOUND", "Data Not Found")

// ErrRequestTimeout Error Status for Request Timeout
var ErrRequestTimeout = error.New("408", "", "ERR_REQUEST_TIMEOUT", "Request Timeout")

// ErrNotAuthorize Error Status for Not Authorize
var ErrNotAuthorize = error.New("401", "", "ERR_NOT_AUHTORIZE", "Not Authorize")

// ErrToken Error Status for Token Error
var ErrToken = error.New("401", "", "ERR_TOKEN", "Not Authorized, Invalid access token")

// ErrCreatingAuth Error Status for Creating Auth
var ErrCreatingAuth = error.New("401", "", "ERR_CREATING_AUTH", "Error Creating Authentication")

// ErrForbidden Error Status for Forbidden Access
var ErrForbidden = error.New("403", "", "ERR_FORBIDDEN", "Forbidden")

// ErrConflict Error Status for Conflict Data
var ErrConflict = error.New("409", "", "ERR_CONFICT", "Data Found, Can't Add Same Data")

// ErrTokenExpired Error Status for Token Expired
var ErrTokenExpired = error.New("401", "", "ERR_TOKEN_EXPIRED", "Not Authorized, access token is expired")

// ErrUnexpectedSigning Error Status for Error Signing Method
var ErrUnexpectedSigning = error.New("401", "", "ERR_UNEXPECTED_SIGNING", "Not Authorized, unexpected signing method")

// ErrInvalidLogin Error Status for Error Login
var ErrInvalidLogin = error.New("400", "", "ERR_INVALID_LOGIN", "Invalid Login, Invalid email or phone_number and password")

// ErrUnProcessableEntity Error Status for Unprocess Entity
var ErrUnProcessableEntity = error.New("422", "", "ERR_UNPROCESSABLE_ENTITY", "Unprocessable Entity")
