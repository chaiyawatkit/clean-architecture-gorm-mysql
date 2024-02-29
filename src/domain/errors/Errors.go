// Package errors defines the domain errors used in the application.
package errors

import "errors"

const (
	// NotFound error indicates a missing / not found record
	NotFound        = "NotFound"
	notFoundMessage = "record not found"

	// ValidationError indicates an error in input validation
	ValidationError        = "ValidationError"
	validationErrorMessage = "validation error"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists     = "ResourceAlreadyExists"
	alreadyExistsErrorMessage = "resource already exists"

	// RepositoryError indicates a repository (e.g database) error
	RepositoryError        = "RepositoryError"
	repositoryErrorMessage = "error in repository operation"

	// NotAuthenticated indicates an authentication error
	NotAuthenticated             = "NotAuthenticated"
	notAuthenticatedErrorMessage = "not Authenticated"

	// TokenGeneratorError indicates an token generation error
	TokenGeneratorError        = "TokenGeneratorError"
	tokenGeneratorErrorMessage = "error in token generation"

	// NotAuthorized indicates an authorization error
	NotAuthorized             = "NotAuthorized"
	notAuthorizedErrorMessage = "not authorized"

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError        = "UnknownError"
	unknownErrorMessage = "something went wrong"

	UrlDecodeError        = "UrlDecodeError"
	urlDecodeErrorMessage = "url decode error"

	RequestError        = "RequestError"
	requestErrorMessage = "request error"

	UnmarshalError        = "UnmarshalError"
	unmarshalErrorMessage = "unmarshal Error"
)

type AppError struct {
	Err  error
	Type string
}

func NewAppError(err error, errType string) *AppError {
	return &AppError{
		Err:  err,
		Type: errType,
	}
}

func NewAppErrorWithType(errType string) *AppError {
	var err error

	switch errType {
	case NotFound:
		err = errors.New(notFoundMessage)
	case ValidationError:
		err = errors.New(validationErrorMessage)
	case ResourceAlreadyExists:
		err = errors.New(alreadyExistsErrorMessage)
	case RepositoryError:
		err = errors.New(repositoryErrorMessage)
	case NotAuthenticated:
		err = errors.New(notAuthenticatedErrorMessage)
	case NotAuthorized:
		err = errors.New(notAuthorizedErrorMessage)
	case TokenGeneratorError:
		err = errors.New(tokenGeneratorErrorMessage)
	case UrlDecodeError:
		err = errors.New(urlDecodeErrorMessage)
	case RequestError:
		err = errors.New(requestErrorMessage)
	case UnmarshalError:
		err = errors.New(unmarshalErrorMessage)

	default:
		err = errors.New(unknownErrorMessage)
	}

	return &AppError{
		Err:  err,
		Type: errType,
	}
}

func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}
