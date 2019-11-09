package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"testing"
)

func TestGetEnvVarsFromDirectory(t *testing.T) {
	t.Run("handle non existent folders correctly", func(t *testing.T) {
		randomPath := "./" + funk.RandomString(5) + "/" + funk.RandomString(5)
		assert.Equal(t, []EnvVar{}, GetEnvVarsFromDirectory(randomPath))
	})
	t.Run("handle existent folders correctly", func(t *testing.T) {
		assert.Equal(t, []EnvVar{{
			Name:  "A_ENV",
			Value: "123",
		}, {
			Name:  "B_VAR",
			Value: "another_val",
		}}, GetEnvVarsFromDirectory("./env-vars"))
	})
}