package token

import "time"

type Config struct {
	SymmetricKey         string        `mapstructure:"symmetric_key"`
	AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
}
