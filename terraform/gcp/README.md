# Learn Terraform for Provisioning GCP Resources

### install terraform 
```bash
curl -o /tmp/terraform.zip -LO https://releases.hashicorp.com/terraform/1.3.6/
terraform_1.3.6_linux_amd64.zip
```

```bash
unzip /tmp/terraform.zip
```

```bash
chmod +x terraform && mv terraform /usr/local/bin/
```

```bash
terraform
```

### login to gcloud
```bash
gcloud auth application-default login --no-launch-browser
```

### several important Terraform commands
```bash
terraform init
```

```bash
terraform apply -var-name <file_name> -auto-approve
```

```bash
terraform destroy -target <resource_name>
```

### using environment variable
```bash
export TF_VAR_resource_name=my-value
```