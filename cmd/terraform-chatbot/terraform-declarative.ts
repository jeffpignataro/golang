import { Construct } from 'constructs';
import { TerraformAwsProvider, TerraformAwsInstance } from 'cdk8s-plus-terraform';

class MyStack extends Construct {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    // Define the AWS provider
    new TerraformAwsProvider(this, 'aws', {
      region: 'us-west-2',
    });

    // Define the EC2 instance
    new TerraformAwsInstance(this, 'example', {
      ami: 'ami-0c55b159cbfafe1f0',
      instanceType: 't2.micro',
      tags: {
        Name: 'example-instance',
      },
    });
  }
}
