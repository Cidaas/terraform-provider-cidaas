// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cidaas": providerserver.NewProtocol6WithError(Provider("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	t.Helper()

	if os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID") == "" {
		t.Fatal("TERRAFORM_PROVIDER_CIDAAS_CLIENT_ID must be set for acceptance tests")
	}

	if os.Getenv("TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET") == "" {
		t.Fatal("TERRAFORM_PROVIDER_CIDAAS_CLIENT_SECRET must be set for acceptance tests")
	}
}

// RandString generates a random string with the given length.
func RandString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
