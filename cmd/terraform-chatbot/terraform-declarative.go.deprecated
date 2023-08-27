package main

import (
	builder "github.com/hashicorp/terraform/configs"
)

func main_declarative() {
	config := builder.NewConfig(
		builder.Provider("aws", map[string]string{
			"region": "us-west-2",
		}),
		builder.Resource("aws_instance", "example", map[string]interface{}{
			"ami":           "ami-0c55b159cbfafe1f0",
			"instance_type": "t2.micro",
			"tags": map[string]string{
				"Name": "example-instance",
			},
		}),
	)

	// Serialize the Terraform configuration to HCL
	hclBytes, err := config.ToHCL()
	if err != nil {
		panic(err)
	}
}
