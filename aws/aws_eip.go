// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEip(client *Client) {
	req := client.ec2conn.DescribeAddressesRequest(&ec2.DescribeAddressesInput{})

	resp, err := req.Send(context.Background())
	if err != nil {
		log.Printf("aws_eip: %s", err)
	} else {
		if len(resp.Addresses) > 0 {
			fmt.Println("")
			fmt.Println("aws_eip:")
			for _, r := range resp.Addresses {
				fmt.Println(*r.AllocationId)
				for _, t := range r.Tags {
					fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
				}
			}
		}
	}

}