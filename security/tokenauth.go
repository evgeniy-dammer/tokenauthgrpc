package security

import "context"

type TokenAuth struct {
	Token string
}

func (tokenAuth TokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Rearer " + tokenAuth.Token,
	}, nil
}

func (TokenAuth) RequireTransportSecurity() bool {
	return true
}
