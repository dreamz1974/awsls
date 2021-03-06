// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdaFunction(client *Client) ([]Resource, error) {
	req := client.Lambdaconn.ListFunctionsRequest(&lambda.ListFunctionsInput{})

	var result []Resource

	p := lambda.NewListFunctionsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Functions {

			result = append(result, Resource{
				Type:      "aws_lambda_function",
				ID:        *r.FunctionName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
