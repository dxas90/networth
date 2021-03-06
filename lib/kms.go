package nwlib

import (
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go/aws"
)

var kmsKeyAlias = GetEnv("KMS_KEY_ALIAS", "alias/networth")

// KMSClient kms client struct
type KMSClient struct {
	*kms.KMS
}

// NewKMSClient new instance of the kms client
func NewKMSClient() *KMSClient {
	cfg := LoadAWSConfig()
	svc := kms.New(cfg)

	return &KMSClient{svc}
}

// Encrypt key
func (k KMSClient) Encrypt(token string) (string, error) {
	input := &kms.EncryptInput{
		KeyId:     aws.String(kmsKeyAlias),
		Plaintext: []byte(token),
	}

	req := k.EncryptRequest(input)
	res, err := req.Send()

	if err != nil {
		log.Printf("Problem encrypting key: %+v\n", err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(res.CiphertextBlob), nil
}

// Decrypt decrypt kms token
func (k KMSClient) Decrypt(token string) (string, error) {
	decoded, _ := base64.StdEncoding.DecodeString(token)
	input := &kms.DecryptInput{
		CiphertextBlob: []byte(decoded),
	}

	req := k.DecryptRequest(input)
	res, err := req.Send()

	if err != nil {
		log.Printf("Problem decrypting key: %+v\n", err)
		return "", err
	}

	return string(res.Plaintext), nil
}
