package main

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

type ValidateUserScopeResponse struct {
	Result struct {
		Valid bool
	}
}

const ValidateUserScopesQuery = `
query ($userID: ID!, $scope: String!) {
	result: validateUserScope(user: $userID, scope: $scope) {
		valid
	}
}
`

func validateScopeForUser(ctx context.Context, scope, userID string) (err error) {
	URL := os.Getenv("SCOPE_VALIDATOR_URL")
	if URL != "" {
		client := graphql.NewClient(URL)
		client.Log = func(s string) {
			fmt.Println(s)
		}
		req := graphql.NewRequest(ValidateUserScopesQuery)
		req.Var("userID", userID)
		req.Var("scope", scope)

		var res ValidateUserScopeResponse
		err = client.Run(ctx, req, &res)
		if err != nil {
			return
		}

		if !res.Result.Valid {
			err = fmt.Errorf("invalid scopes")
		}
	}
	return
}
