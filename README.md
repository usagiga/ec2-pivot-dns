# ec2-pivot-dns
Patch DNS records from ephemeral IP address like DDNS

## Installation

:warning: Under constructing, TBU

1. Set up Cloudflare
2. Set up AWS
    - IAM Policy
        - [Examples](./examples/iam-policy.json)
    - IAM Role
        - `AWSLambdaBasicExecutionRole`
        - the policy above
    - Lambda
        - Runtime: `Go 1.x`
        - Handler: `OnEC2StateChange`
        - Arch: `x86_64`
        - IAM Role: the role above
    - Secret Manager
    - EventBridge
    - EC2
3. Download *ec2-pivot-dns-lambda.zip* from [Releases](https://github.com/usagiga/ec2-pivot-dns/releases/latest)
    - Or, run `curl -fsSL https://github.com/usagiga/ec2-pivot-dns/releases/download/${VERSION}/ec2-pivot-dns-lambda.zip --output ec2-pivot-dns-lambda.zip`
4. Upload *ec2-pivot-dns-lambda.zip* to AWS Lambda

## Usage

:warning: TBU

## LICENSE

MIT
