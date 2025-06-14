package ro

type JwtPair struct {
	AccessToken      string  `json:"accessToken"`
	RefreshToken     *string `json:"refreshToken"`
	AccessExpiresIn  string  `json:"accessExpiresIn"`
	RefreshExpiresIn *string `json:"refreshExpiresIn"`
}
