package jwt

import (
	"container/list"
	"encoding/base64"

	"github.com/go-masonry/mortar/interfaces/auth/jwt"
)

// JSONDecoder json.Unmarshaler
type JSONDecoder func(data []byte, v interface{}) error

// ExtractorBuilder defines what can be configured when building JWT Token Extractor
type ExtractorBuilder interface {
	SetDecoder(dec JSONDecoder) ExtractorBuilder
	SetContextExtractor(extractor jwt.ContextExtractor) ExtractorBuilder
	SetBase64Decoder(dec *base64.Encoding) ExtractorBuilder
	Build() jwt.TokenExtractor
}

type extractorConfig struct {
	jsonDecoder      JSONDecoder
	base64Enc        *base64.Encoding
	contextExtractor jwt.ContextExtractor
}

type builder struct {
	ll *list.List
}

// Builder creates a fresh instance of Extractor Builder
func Builder() ExtractorBuilder { _ = "STUB: not implemented"; return *new(ExtractorBuilder) }

func (b *builder) SetDecoder(dec JSONDecoder) ExtractorBuilder {
	_ = "STUB: not implemented"
	return *new(ExtractorBuilder)
}

func (b *builder) SetContextExtractor(extractor jwt.ContextExtractor) ExtractorBuilder {
	_ = "STUB: not implemented"
	return *new(ExtractorBuilder)
}

func (b *builder) SetBase64Decoder(dec *base64.Encoding) ExtractorBuilder {
	_ = "STUB: not implemented"
	return *new(ExtractorBuilder)
}

func (b *builder) Build() jwt.TokenExtractor {
	_ = "STUB: not implemented"
	return *new(jwt.TokenExtractor)
}
