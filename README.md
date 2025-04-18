### 🌩️ Cloud Abstractor

**Cloud Abstractor** is a lightweight Go application that reads a `deploy.yaml` configuration file and generates a `main.tf` Terraform script to provision AWS infrastructure such as:
- ECS Clusters
- S3 Buckets

### 🛠️ Requirements

- [Go](https://golang.org/doc/install)
- [Terraform](https://developer.hashicorp.com/terraform/downloads)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) (configured via `aws configure`)
- AWS IAM user with permissions for ECS and S3

### 🧾 Sample YAML (`deploy.yaml`)

```yaml
app:
  name: myapp2
  provider: aws
  compute:
    type: container
    cpu: "256"
    memory: "512"
    image: nginx:latest
  storage:
    type: s3
    bucket_name: my-app-bucket-unique1234
