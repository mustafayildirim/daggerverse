// A generated module for AzureKeyVault functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type AzureKeyVault struct{}

// Example: dagger call get-secret --tenant-id=env:AZURE_TENANT_ID --client-id=env:AZURE_CLIENT_ID --client-secret=env:AZURE_CLIENT_SECRET --key-vault-name quickstart-kv --secret-name test-secret1
func (m *AzureKeyVault) GetSecret(ctx context.Context, keyVaultName string, secretName string, tenantId *Secret, clientId *Secret, clientSecret *Secret) (string, error) {
	keyVaultUrl := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)

	plaintextTenantId, err := tenantId.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	plaintextClientId, err := clientId.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	plaintextSecretId, err := clientSecret.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	os.Setenv("AZURE_TENANT_ID", plaintextTenantId)
	os.Setenv("AZURE_CLIENT_ID", plaintextClientId)
	os.Setenv("AZURE_CLIENT_SECRET", plaintextSecretId)

	//Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	secretClientOptions := azsecrets.ClientOptions{
		DisableChallengeResourceVerification: true,
	}

	//Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(keyVaultUrl, cred, &secretClientOptions)
	if err != nil {
		log.Fatalf("failed to connect to client: %v", err)
	}

	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	resp, err := client.GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
	}

	return *resp.Value, nil
}
