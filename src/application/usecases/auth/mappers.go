// Package auth provides the use case for authentication
package auth

import (
	userDomain "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/user"
)

func secAuthUserMapper(domainUser *userDomain.User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			UserName:  domainUser.UserName,
			Email:     domainUser.Email,
			FirstName: domainUser.FirstName,
			LastName:  domainUser.LastName,
			ID:        domainUser.ID,
			Status:    domainUser.Status,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}

}
