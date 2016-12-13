package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	output := NewDefaultLogger()

	for missingEnvVar := range checkEnvironment() {
		output.LogError(fmt.Sprintf("Missing Environment Vaiable\"%s\".", missingEnvVar))
	}

	if !checkAzureCLI() {
		output.LogError("Azure CLI not found on this machine. This will prevent logging into Azure.")
	}

	errorCount := output.GetErrorCount()

	fmt.Println()
	if 0 == errorCount {
		if 0 != output.GetWarningCount() {
			fmt.Printf("Warnings were found. You MAY run into trouble using the Azure-SDK-for-Go samples.\n")
		} else {

			fmt.Printf("No errors were found. You are READY to run the Azure-SDK-for-Go samples.\n")
		}
	} else {
		fmt.Printf("Errors found: %d\nYou are NOT ready to run the Azure-SDK-for-Go samples.\n", errorCount)
	}
}

// checkEnvironment scans the current working environment for known values that must be present.
// For each unset environment variable it finds, it will write the name of that variable to the
// returned channel. When it has scanned all variables, the channel will be closed.
func checkEnvironment() <-chan string {
	retval := make(chan string)

	keys := []string{
		"AZURE_TENANT_ID",
		"AZURE_CLIENT_ID",
		"AZURE_CLIENT_SECRET",
		"AZURE_SUBSCRIPTION_ID"}

	go func() {
		for _, envVar := range keys {
			if value := os.Getenv(envVar); "" == value {
				retval <- envVar
			}
		}
		close(retval)
	}()

	return retval
}

func checkAzureCLI() bool {
	if _, err := exec.LookPath("azure"); nil == err {
		return true
	}
	return false
}
