package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestLinuxVMPreChecks(t *testing.T) {
	t.Parallel()

	// Locate the correct directory inside .terragrunt-cache
	baseTerragruntDir := "/Users/snsbabu/MyData/first_tasks/azure_provisioning_template/sites/southindia/linux_vm/.terragrunt-cache"
	terraformConfigDir, err := findTerraformConfigDir(baseTerragruntDir)
	if err != nil {
		t.Fatalf("Failed to find Terraform configuration directory: %v", err)
	}

	terraformOptions := &terraform.Options{
		TerraformDir: terraformConfigDir,
	}

	// Validate Terraform Syntax
	terraform.InitAndValidate(t, terraformOptions)

	// Run Terraform Plan and Capture Output
	stdout, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "plan", "-out=tfplan.out")
	if err != nil {
		t.Fatalf("Failed to generate Terraform plan: %v", err)
	}

	// Assertions: Ensure expected resources are in the plan output
	assert.Contains(t, stdout, "azurerm_resource_group", "Resource Group should be planned.")
	assert.Contains(t, stdout, "azurerm_virtual_network", "VNet should be planned.")
	assert.Contains(t, stdout, "azurerm_subnet", "Subnet should be planned.")
	assert.Contains(t, stdout, "azurerm_network_security_group", "NSG should be planned.")
	assert.Contains(t, stdout, "azurerm_linux_virtual_machine", "Linux VM should be planned.")
}

// Helper function to locate Terraform configuration inside .terragrunt-cache
func findTerraformConfigDir(baseDir string) (string, error) {
	var terraformDir string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.Contains(path, "modules/azure/infrastructure/linux_vm") {
			terraformDir = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	if terraformDir == "" {
		return "", os.ErrNotExist
	}
	return terraformDir, nil
}

// package test

// import (
// 	"os"
// 	"path/filepath"
// 	"testing"

// 	"github.com/gruntwork-io/terratest/modules/terraform"
// )

// func TestLinuxVMPreChecks(t *testing.T) {
// 	t.Parallel()

// 	// Locate the latest .terragrunt-cache directory dynamically
// 	baseTerragruntDir := "/Users/snsbabu/MyData/first_tasks/azure_provisioning_template/sites/southindia/linux_vm/.terragrunt-cache"
// 	latestCacheDir, err := findLatestTerragruntCache(baseTerragruntDir)
// 	if err != nil {
// 		t.Fatalf("Failed to find Terragrunt cache directory: %v", err)
// 	}

// 	terraformOptions := &terraform.Options{
// 		TerraformDir: latestCacheDir,
// 	}

// 	// Validate Terraform Syntax
// 	terraform.InitAndValidate(t, terraformOptions)
// }

// // Helper function to find the latest .terragrunt-cache directory
// func findLatestTerragruntCache(baseDir string) (string, error) {
// 	files, err := os.ReadDir(baseDir)
// 	if err != nil {
// 		return "", err
// 	}

// 	for _, f := range files {
// 		if f.IsDir() {
// 			// Construct full path to cache directory
// 			cachePath := filepath.Join(baseDir, f.Name())
// 			return cachePath, nil // Return the first valid directory found
// 		}
// 	}
// 	return "", os.ErrNotExist
// }

// package test

// import (
// 	"testing"

// 	"github.com/gruntwork-io/terratest/modules/terraform"
// 	"github.com/stretchr/testify/assert"
// )

// func TestLinuxVMPreChecks(t *testing.T) {
// 	t.Parallel()

// 	terraformOptions := &terraform.Options{
// 		TerraformDir: "../../../../sites/southindia/linux_vm/",
// 		PlanFilePath: "../../../../sites/southindia/linux_vm/tfplan.out",
// 	}

// 	// Validate Terraform Syntax
// 	terraform.InitAndValidate(t, terraformOptions)

// 	// Retrieve planned outputs before deployment
// 	plannedOutputs := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

// 	// Ensure Resource Group is being created
// 	assert.Contains(t, plannedOutputs, "azurerm_resource_group", "Resource Group should be planned.")

// 	// Ensure VNet is being created
// 	assert.Contains(t, plannedOutputs, "azurerm_virtual_network", "VNet should be planned.")

// 	// Ensure Subnet is being created
// 	assert.Contains(t, plannedOutputs, "azurerm_subnet", "Subnet should be planned.")

// 	// Ensure NSG is being created
// 	assert.Contains(t, plannedOutputs, "azurerm_network_security_group", "NSG should be planned.")

// 	// Ensure Linux VM is being created
// 	assert.Contains(t, plannedOutputs, "azurerm_linux_virtual_machine", "Linux VM should be planned.")
// }
