// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

func ListDevicefarmProject(client *Client) ([]Resource, error) {
	req := client.Devicefarmconn.ListProjectsRequest(&devicefarm.ListProjectsInput{})

	var result []Resource

	p := devicefarm.NewListProjectsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Projects {

			result = append(result, Resource{
				Type:      "aws_devicefarm_project",
				ID:        *r.Arn,
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
