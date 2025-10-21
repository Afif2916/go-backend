package test

import (
    "testing"

    "github.com/Afif2916/go-backend/internal/models"
    "github.com/Afif2916/go-backend/internal/services"
    "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    user := models.User{Name: "Test User", Email: "test@example.com"}
    err := services.CreateUser(&user) // ✅ panggil lewat package services
    assert.Nil(t, err)
}
