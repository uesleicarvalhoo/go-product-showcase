package config

import (
	"github.com/netflix/go-env"
)

func LoadFromEnv(s any) error {
	_, err := env.UnmarshalFromEnviron(s)

	return err
}
