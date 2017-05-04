package authentication

import (
	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/client"
)

const (
	//ErrCodeInvalidToken  ...
	ErrCodeInvalidToken = "InvalidToken"
)

//PopulateTokens populate the relevant tokens in the bluemix Config using the token provider
func PopulateTokens(tokenProvider client.TokenProvider, c *bluemix.Config) error {
	if c.IBMIDPassword != "" {
		err := tokenProvider.AuthenticatePassword(c.IBMID, c.IBMIDPassword)
		return err
	}
	if c.BluemixAPIKey != "" {
		err := tokenProvider.AuthenticateAPIKey(c.BluemixAPIKey)
		return err
	}
	return nil
}