package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/Luc1808/go-task-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
var refreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")

const (
	accessTokenDuration  = 30 * time.Minute
	refreshTokenDuration = 7 * 24 * time.Hour
)

type AuthHandler struct {
	*gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	// Get data from body
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Transform data
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	user.Password = string(hashedPassword)

	// Set creation time
	user.CreatedAt = time.Now()

	// Save data into db
	if err := h.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user", "err": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	// Get data from body
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Get data from db
	var user models.User
	if err := h.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// issue a new token pair using a valid refresh token
func (h *AuthHandler) RefreshHandler(ctx *gin.Context) {
	var tokenRequest struct {
		RefreshToken string `json:"refresh+token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&tokenRequest); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// Parse refresh token
	token, err := jwt.Parse(tokenRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token claims"})
		return
	}

	userID := int(claims["user_id"].(float64))

	// Verify token in DB
	storedToken, err := h.findRefreshToken(tokenRequest.RefreshToken, userID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new token pair
	newTokenPair, err := generateTokenPair(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Change old for new tokens
	h.deleteRefreshToken(storedToken.ID)
	h.storeRefreshToken(userID, newTokenPair.RefreshToken, time.Now().Add(refreshTokenDuration))

	ctx.JSON(http.StatusOK, newTokenPair)
}

func generateTokenPair(userID int) (models.TokenPair, error) {
	var tokenPair models.TokenPair

	// Create access token
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(accessTokenDuration).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	var err error
	tokenPair.AcessToken, err = accessToken.SignedString(accessTokenSecret)
	if err != nil {
		return tokenPair, err
	}

	// Create refresh token
	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(refreshTokenDuration).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenPair.RefreshToken, err = refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		return tokenPair, err
	}

	return tokenPair, nil
}

func (h *AuthHandler) storeRefreshToken(userID int, token string, expiresAt time.Time) error {
	refreshToken := models.RefreshToken{
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}

	result := h.DB.Create(&refreshToken)
	return result.Error
}

func (h *AuthHandler) findRefreshToken(token string, userID int) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	result := h.DB.Where("token = ? AND user_id = ? AND expires_at > ?", token, userID, time.Now()).First(&refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &refreshToken, nil
}

func (h *AuthHandler) deleteRefreshToken(id int) error {
	result := h.DB.Delete(&models.RefreshToken{}, id)
	return result.Error
}

func (h *AuthHandler) clearExpiredTokens() error {
	result := h.DB.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{})
	return result.Error
}
