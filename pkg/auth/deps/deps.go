package deps

import (
	"statuarius/internal/config"
	"statuarius/pkg/auth/dal"
)

type AuthDep struct {
	DAL dal.AuthDAL
}

func New(cfg *config.AuthConfig) *AuthDep {
	return &AuthDep{}
}
