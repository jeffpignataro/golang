package okta

import (
	"context"
	"fmt"

	"github.com/okta/okta-sdk-golang/v2/okta"
)

func NewClient(token string) (context.Context, okta.Client) {
	ctx, client, err := okta.NewClient(
		context.TODO(),
		okta.WithOrgUrl("https://dev-20308777.okta.com"),
		okta.WithToken(token),
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return ctx, *client
}
