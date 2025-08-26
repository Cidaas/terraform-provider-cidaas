## Developer Guide

We welcome contributions to the Terraform Provider! This document outlines the process for contributing and provides guidelines to ensure consistency and quality.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Environment Setup](#development-environment-setup)
- [Testing](#testing)
- [Code Standards](#code-standards)
- [Release Process](#release-process)
- [Getting Help](#getting-help)

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.24+
- [Terraform](https://developer.hashicorp.com/terraform/install) 1.0+
- [make](https://www.gnu.org/software/make/) (recommended) or run Go commands directly
    - **Linux/macOS**: Usually pre-installed
    - **Windows**: Install via [Chocolatey](https://chocolatey.org/packages/make) (`choco install make`) or use [WSL](https://docs.microsoft.com/en-us/windows/wsl/)

> **Note**: While `make` is recommended for convenience, you can also run the Go commands directly if make is not available on your system.

## Development Environment Setup

### Install Dependencies

First, install the required development tools:
```bash
make deps
```
This will install `golangci-lint` and other necessary development dependencies.

### Build the Provider

Build the provider binary:
```bash
make build
```
This creates the `terraform-provider-cidaas` binary in the current directory.

### Install the Provider Locally

Install the provider to your local Terraform plugins directory:
```bash
make install
```
This will:
1. Build the provider binary
2. Create the appropriate plugin directory structure
3. Install the binary to `~/.terraform.d/plugins/hashicorp.com/cidaas/cidaas/1.0.0/linux_amd64/`

### Code Formatting

Format your code before committing:
```bash
make fmt
```
This runs `gofmt` to automatically fix code formatting issues.

### Code Quality Checks

Run linting and format checks:
```bash
make fmtcheck
```
This will:
- Check code formatting with `gofmtcheck.sh`
- Run `golangci-lint` with the project's lint configuration

### Use the Provider Locally

After installation, use this configuration in your Terraform files to use the local provider:

```hcl
terraform {
  required_providers {
    cidaas = {
      source  = "hashicorp.com/cidaas/cidaas"
      version = "1.0.0"
    }
  }
}

provider "cidaas" {
  base_url = "YOUR CIDAAS BASE URL"
}
```

## Testing

### Unit Tests

Run all unit tests:
```bash
make test
```
This will:
- Run format and lint checks (`make fmtcheck`)
- Execute unit tests with a 5-minute timeout
- Run tests in parallel for faster execution

Run unit tests without linting:
```bash
go test ./...
```

### Acceptance Tests

**Important**: Acceptance tests create real resources in your Cidaas environment. Ensure you have:
- Proper credentials configured
- A dedicated test environment (never use production)
- Sufficient permissions to create/modify resources

#### Environment Setup for Acceptance Tests

Set up the required environment variables:
```bash
export CIDAAS_BASE_URL="your-cidaas-base-url"
export CIDAAS_CLIENT_ID="your-test-client-id"
export CIDAAS_CLIENT_SECRET="your-test-client-secret"
```

#### Run All Acceptance Tests

```bash
make testacc
```
This runs all acceptance tests with:
- `TF_ACC=1` environment variable set
- 120-minute timeout
- Verbose output

#### Run Specific Acceptance Tests

Run a single specific test case:
```bash
TF_ACC=1 go test ./internal/provider -run TestApp_Basic -v
```

### Test Coverage

Generate comprehensive test coverage report:
```bash
make test-ci
```
This will:
- Run all acceptance tests with coverage profiling
- Generate coverage statistics in the terminal
- Create an HTML coverage report (`coverage.html`)
- Run tests in parallel for CI/CD environments

Open the HTML report in your browser

### Test Guidelines

#### Writing Tests
- All new resources must have acceptance tests
- Include both positive and negative test cases
- Test resource import functionality
- Test resource updates and lifecycle management
- Use meaningful test names: `TestAccResourceName_scenario`

### Troubleshooting Tests

**Common Issues:**
- **Authentication errors**: Verify environment variables are set correctly
- **Timeout errors**: Increase timeout for slow API responses
- **Resource conflicts**: Ensure test resources have unique names
- **Permission errors**: Verify your test credentials have sufficient permissions

**Debug failing tests:**
```bash
# Run with extra verbose output
TF_ACC=1 TF_LOG=DEBUG go test ./internal/provider -run TestApp_Basic -v
```

## Code Standards

### Adding New Resources and Data Sources

When adding new resources or data sources, follow the [Terraform Plugin Framework documentation](https://developer.hashicorp.com/terraform/plugin/framework):

- **Resources**: Follow the [Resource implementation guide](https://developer.hashicorp.com/terraform/plugin/framework/resources)
- **Data Sources**: Follow the [Data Source implementation guide](https://developer.hashicorp.com/terraform/plugin/framework/data-sources)
- **Schema Design**: Use the [Schema concepts](https://developer.hashicorp.com/terraform/plugin/framework/schemas) for proper attribute definitions
- **Validation**: Implement [Validators](https://developer.hashicorp.com/terraform/plugin/framework/validation) for input validation
- **Testing**: Follow [Testing patterns](https://developer.hashicorp.com/terraform/plugin/testing) for comprehensive test coverage

### Go Code Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Handle errors appropriately
- Use context for API calls

### Terraform Schema

- Use appropriate attribute types (`schema.StringAttribute`, `schema.Int64Attribute`, etc.)
- Add comprehensive descriptions for all attributes
- Use validators where appropriate
- Implement plan modifiers for computed attributes
- Follow Terraform naming conventions (snake_case)

### Documentation

Documentation is automatically generated using `go generate` with the `terraform-plugin-docs` tool. Developers need to ensure:

#### Resource/Data Source Documentation
- Add comprehensive `MarkdownDescription` fields to all schema attributes
- Include clear descriptions for resources, data sources, and their attributes
- Use proper markdown formatting in descriptions

Example:
```go
"client_name": schema.StringAttribute{
    Required:            true,
    MarkdownDescription: "The name of the client. Must be unique within the tenant.",
},
```

#### Examples
- Add working Terraform configuration examples in the `examples/` folder
- Organize examples by resource type: `examples/resources/cidaas_app/`
- Ensure all examples are tested and up-to-date
- Include both basic and advanced usage scenarios

#### Generate Documentation
Run documentation generation after making changes:
```bash
go generate ./...
```

This will:
- Scan all resources and data sources for schema descriptions
- Generate markdown files in the `docs/` folder
- Include examples from the `examples/` folder
- Create provider documentation automatically

## Release Process

> **Note**: The release process is handled by internal maintainers. This section is for reference and contribution context.

### Versioning

We follow [Semantic Versioning](https://semver.org/):
- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Contributor Guidelines

As a contributor, ensure your changes include:

1. **Update Changelog**
   
   Update `CHANGELOG.md` with your changes:
   - New features
   - Bug fixes
   - Breaking changes
   - Deprecations
   
   Format example:
   ```markdown
   ## 1.0.0
   
   ### Added
   - New resource `cidaas_user_group`
   - Support for custom fields in custom_provider resource
   
   ### Fixed
   - Fixed issue with app update operations
   ```

2. **Quality Checks**
   
   Ensure all checks pass before submitting:
   ```bash
   make fmtcheck
   make test
   make testacc
   ```

### Internal Release Workflow

The release process is automated:

1. **Internal maintainers** create release tags in the main GitLab repository
2. **GitHub mirror** automatically syncs the release tags
3. **GitHub CI** builds cross-platform binaries automatically
4. **Terraform Registry** syncs and publishes the new version

### For Contributors

- Focus on code quality and comprehensive testing
- Update documentation and changelog entries
- All release management is handled by the maintainer team
- No manual binary building required

## Getting Help

#### GitHub Issues
- [Create an issue](https://github.com/Cidaas/terraform-provider-cidaas/issues/new) for:
  - Bug reports
  - Feature requests  
  - General questions
  - Documentation improvements

### Before Asking for Help

1. **Check Existing Issues**: Search [existing issues](https://github.com/Cidaas/terraform-provider-cidaas/issues) for similar problems
2. **Review Documentation**: Check the [provider documentation](docs/) and [examples](examples/)
3. **Test with Latest Version**: Ensure you're using the latest provider version
4. **Check Terraform Logs**: Enable debug logging with `TF_LOG=DEBUG`

### Useful Resources

- [Terraform Plugin Framework Documentation](https://developer.hashicorp.com/terraform/plugin/framework)
- [Terraform Plugin Testing](https://developer.hashicorp.com/terraform/plugin/testing)
- [Terraform Provider Design Principles](https://developer.hashicorp.com/terraform/plugin/best-practices/hashicorp-provider-design-principles)
- [Terraform Registry Provider Guidelines](https://developer.hashicorp.com/terraform/registry/providers/publishing)
- [Go Documentation](https://golang.org/doc/)

### Getting Quality Help

To get the best assistance:

1. **Provide Context**: Include your Terraform configuration (sanitized)
2. **Include Versions**: Specify provider version, Terraform version, and OS
3. **Share Logs**: Include relevant error messages and debug logs
4. **Minimal Reproduction**: Provide the smallest configuration that reproduces the issue

---

**Note**: This provider follows Terraform Plugin Framework standards and best practices. For framework-specific questions, refer to the official Terraform documentation linked above.

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.

---

Thank you for contributing to the Terraform Provider! 🚀