package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MockHandler struct {
	db *gorm.DB
}

func NewMockHandler(db *gorm.DB) *MockHandler {
	return &MockHandler{db: db}
}

func (m *MockHandler) GenerateRandomData(c *gin.Context) {

}
