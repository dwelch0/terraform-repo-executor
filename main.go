package main

import (
	"log"
	"os"

	"github.com/app-sre/terraform-repo-executor/pkg"
)

const (
	CONFIG_FILE     = "CONFIG_FILE"
	GL_USERNAME     = "GITLAB_USERNAME"
	GL_TOKEN        = "GITLAB_TOKEN"
	VAULT_ADDR      = "VAULT_ADDR"
	VAULT_ROLE_ID   = "VAULT_ROLE_ID"
	VAULT_SECRET_ID = "VAULT_SECRET_ID"
	WORKDIR         = "WORKDIR"
)

func main() {
	cfgPath := getEnvOrDefault(CONFIG_FILE, "/config.yaml")
	workdir := getEnvOrDefault(WORKDIR, "/tf-repo")
	glUsername := getEnvOrError(GL_USERNAME)
	glToken := getEnvOrError(GL_TOKEN)
	vaultAddr := getEnvOrError(VAULT_ADDR)
	roleId := getEnvOrError(VAULT_ROLE_ID)
	secretId := getEnvOrError(VAULT_SECRET_ID)

	err := pkg.Run(cfgPath,
		workdir,
		glUsername,
		glToken,
		vaultAddr,
		roleId,
		secretId,
	)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("%s not set. Using default value: `%s`", key, defaultValue)
		return defaultValue
	}
	return value
}

func getEnvOrError(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is required", key)
	}
	return value
}
