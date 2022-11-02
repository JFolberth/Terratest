package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"

)

func TestTerraformStorageAccount(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../infrastructure",
	})

	//defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)


}