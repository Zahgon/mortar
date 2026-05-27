package jwt

import (
	"context"

	"github.com/go-masonry/mortar/interfaces/auth/jwt"
)

type tokenExtractorImpl struct {
	cfg *extractorConfig
}

func newTokenExtractor(cfg *extractorConfig) jwt.TokenExtractor {
	_ = "STUB: not implemented"
	return *new(jwt.TokenExtractor)
}

func (t *tokenExtractorImpl) FromContext(ctx context.Context) (jwt.Token, error) {
	_ = "STUB: not implemented"
	return *new(jwt.Token), nil
}

func (t *tokenExtractorImpl) FromString(str string) (token jwt.Token, err error) {
	_ = "STUB: not implemented"
	return *new(jwt.Token), nil
}

type tokenInstance struct {
	raw         string
	payload     []byte
	jsonDecoder JSONDecoder
}

func newToken(jwtAsString string, justPayload []byte, decoder JSONDecoder) jwt.Token {
	_ = "STUB: not implemented"
	return *new(jwt.Token)
}

func (t *tokenInstance) Raw() string { _ = "STUB: not implemented"; return "" }

func (t *tokenInstance) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (t *tokenInstance) Map() (output map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tokenInstance) Decode(target interface{}) error { _ = "STUB: not implemented"; return nil }
