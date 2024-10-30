package secrets

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	"golang.org/x/crypto/nacl/box"
)

// this matches the secrets config file by the users, see the secretsConfig.yaml file
type SecretsConfig struct {
	SecretsNames map[string][]string `yaml:"secretsNames"`
}

// this is the payload that will be encrypted
type SecretPayloadToEncrypt struct {
	WorkflowOwner string            `json:"workflowOwner"`
	Secrets       map[string]string `json:"secrets"`
}

// this holds the mapping of secret name (e.g. API_KEY) to the local environment variable name which points to the raw secret
type AssignedSecrets struct {
	WorkflowSecretName string `json:"workflowSecretName"`
	LocalEnvVarName    string `json:"localEnvVarName"`
}

// this is the metadata that will be stored in the encrypted secrets file
type Metadata struct {
	WorkflowOwner            string                       `json:"workflowOwner"`
	CapabilitiesRegistry     string                       `json:"capabilitiesRegistry"`
	DonId                    string                       `json:"donId"`
	DateEncrypted            string                       `json:"dateEncrypted"`
	NodePublicEncryptionKeys map[string]string            `json:"nodePublicEncryptionKeys"`
	EnvVarsAssignedToNodes   map[string][]AssignedSecrets `json:"envVarsAssignedToNodes"`
}

// this is the result of the encryption, will be used by the DON
type EncryptedSecretsResult struct {
	EncryptedSecrets map[string]string `json:"encryptedSecrets"`
	Metadata         Metadata          `json:"metadata"`
}

func ContainsP2pId(p2pId [32]byte, p2pIds [][32]byte) bool {
	for _, id := range p2pIds {
		if id == p2pId {
			return true
		}
	}
	return false
}

func EncryptSecretsForNodes(
	workflowOwner string,
	secrets map[string][]string,
	encryptionPublicKeys map[string][32]byte,
	config SecretsConfig,
) (map[string]string, map[string][]AssignedSecrets, error) {
	encryptedSecrets := make(map[string]string)
	secretsEnvVarsByNode := make(map[string][]AssignedSecrets) // Only used for metadata
	i := 0

	// Encrypt secrets for each node
	for p2pId, encryptionPublicKey := range encryptionPublicKeys {
		secretsPayload := SecretPayloadToEncrypt{
			WorkflowOwner: workflowOwner,
			Secrets:       make(map[string]string),
		}

		for secretName, secretValues := range secrets {
			// Assign secrets to nodes in a round-robin fashion
			secretValue := secretValues[i%len(secretValues)]
			secretsPayload.Secrets[secretName] = secretValue
		}

		// Marshal the secrets payload into JSON
		secretsJSON, err := json.Marshal(secretsPayload)
		if err != nil {
			return nil, nil, err
		}

		// Encrypt secrets payload
		encrypted, err := box.SealAnonymous(nil, secretsJSON, &encryptionPublicKey, rand.Reader)
		if err != nil {
			return nil, nil, err
		}
		encryptedSecrets[p2pId] = base64.StdEncoding.EncodeToString(encrypted)

		// Generate metadata showing which nodes were assigned which environment variables
		for secretName, envVarNames := range config.SecretsNames {
			secretsEnvVarsByNode[p2pId] = append(secretsEnvVarsByNode[p2pId], AssignedSecrets{
				WorkflowSecretName: secretName,
				LocalEnvVarName:    envVarNames[i%len(envVarNames)],
			})
		}

		i++
	}

	return encryptedSecrets, secretsEnvVarsByNode, nil
}
