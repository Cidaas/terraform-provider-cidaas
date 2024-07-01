![Logo](https://raw.githubusercontent.com/Cidaas/terraform-provider-cidaas/master/logo.jpg)

## About cidaas:
[cidaas](https://www.cidaas.com)
 is a fast and secure Cloud Identity & Access Management solution that standardises what’s important and simplifies what’s complex.

## Feature set includes:
* Single Sign On (SSO) based on OAuth 2.0, OpenID Connect, SAML 2.0 
* Multi-Factor-Authentication with more than 14 authentication methods, including TOTP and FIDO2 
* Passwordless Authentication 
* Social Login (e.g. Facebook, Google, LinkedIn and more) as well as Enterprise Identity Provider (e.g. SAML or AD) 
* Security in Machine-to-Machine (M2M) and IoT

# Terraform Provider for Cidaas

The Terraform provider for Cidaas enables interaction with Cidaas instances that allows to perform CRUD operations on applications, custom providers, registration fields and many other functionalities. From managing applications to configuring custom providers, the Terraform provider enhances the user's capacity to define, provision and manipulate their Cidaas resources.

## Prerequisites

- Ensure Terraform is installed on your local machine. Find installation instructions for different operating systems [here](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli).
- [Go](https://go.dev/doc/install) (1.21)



## Documentation

Official documentation on how to use this provider can be found on the
[Terraform Registry](https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs).

## Example Usage

Below is a step-by-step guide to help you set up the provider, configure essential environment variables and integrate the provider into your configuration:

### 1. Terraform Provider Declaration

Begin by specifying the Cidaas provider in your `terraform` block in your Terraform configuration file:

```hcl
terraform {
    required_providers {
      cidaas = {
        version = "3.0.0"
        source  = "Cidaas/cidaas"
      }
    }
}
```

Terraform pulls the version configured of the Cidaas provider for your infrastructure.

### 2. Setup Environment Variables

To authenticate and authorize Terraform operations with Cidaas, set the necessary environment variables. These variables include your Cidaas client credentials, allowing the Terraform provider to complete the client credentials flow and generate an access_token. Execute the following commands in your terminal, replacing placeholders with your actual Cidaas client ID and client secret.

### For Linux and MacOS:
```bash
export TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID="ENTER CIDAAS CLIENT ID"
export TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET="ENTER CIDAAS CLIENT SECRET"
```

### For Windows:
```bash
Set-Item -Path env:TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID -Value “ENTER CIDAAS CLIENT ID“
Set-Item -Path env:TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET -Value “ENTER CIDAAS CLIENT SECRET“
```

You can get a set of client credentials from the Cidaas Admin UI by creating a new client. Simply go to the `Apps` > `App Settings` > `Create New App`. It's important to note that when creating the client, you must select the app type as **Non-Interactive**.

### 3. Add Cidaas Provider Configuration

Next, add the Cidaas provider configuration to your Terraform configuration file. Specify the `base_url` parameter to point to your Cidaas instance. For reference, check the example folder.

```hcl
provider "cidaas" {
  base_url = "https://cidaas.de"
}
```

**Note:** Starting from version 2.5.1, the `redirect_url` is no longer supported in the provider configuration. Ensure that you adjust your configuration accordingly.

By following these steps, you integrate the Cidaas Terraform provider, enabling you to manage your Cidaas resources with Terraform.

## Supported Resources

The Terraform provider for Cidaas supports a variety of resources that enables you to manage and configure different aspects of your Cidaas environment. These resources are designed to integrate with Terraform workflows, allowing you to define, provision and manage your Cidaas resources as code.


Explore the resources in [Official Documentation](https://registry.terraform.io/providers/Cidaas/cidaas/latest/docs) to understand their attributes, functionalities and how to use them in your Terraform configurations.
