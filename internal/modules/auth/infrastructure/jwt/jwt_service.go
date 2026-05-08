package jwt

import (
	"errors"
	"samurenkoroma/services/internal/infrastructure/configs"
	"samurenkoroma/services/internal/modules/auth/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims структура JWT токена (с OrganizationID)
type Claims struct {
	UserID         string `json:"user_id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Role           string `json:"role"`            // глобальная роль (для админки)
	OrganizationID string `json:"organization_id"` // ТЕКУЩАЯ ОРГАНИЗАЦИЯ!
	OrgRole        string `json:"org_role"`        // роль в текущей организации
	jwt.RegisteredClaims
}

// TokenPair пара токенов
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
}

// Service JWT сервис
type Service struct {
	config configs.AuthConfig
}

func NewService(config configs.AuthConfig) *Service {
	return &Service{config: config}
}

// GenerateTokenPair генерирует пару токенов с organizationID
func (s *Service) GenerateTokenPair(userID, username, email, globalRole, organizationID, orgRole string) (*TokenPair, error) {
	accessToken, err := s.generateAccessToken(userID, username, email, globalRole, organizationID, orgRole)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(userID, username, email, globalRole, organizationID, orgRole)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.config.AccessExpiry.Seconds()),
	}, nil
}

func (s *Service) generateAccessToken(userID, username, email, globalRole, organizationID, orgRole string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:         userID,
		Username:       username,
		Email:          email,
		Role:           globalRole,
		OrganizationID: organizationID,
		OrgRole:        orgRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.config.AccessExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    s.config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.SecretKey))
}

func (s *Service) generateRefreshToken(userID, username, email, globalRole, organizationID, orgRole string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:         userID,
		Username:       username,
		Email:          email,
		Role:           globalRole,
		OrganizationID: organizationID,
		OrgRole:        orgRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.config.RefreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    s.config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.SecretKey))
}

// ValidateToken валидирует токен и возвращает claims
func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, domain.ErrTokenExpired
		}
		return nil, domain.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, domain.ErrInvalidToken
}

// RefreshToken обновляет access токен
func (s *Service) RefreshToken(refreshToken string) (*TokenPair, error) {
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	return s.GenerateTokenPair(
		claims.UserID,
		claims.Username,
		claims.Email,
		claims.Role,
		claims.OrganizationID,
		claims.OrgRole,
	)
}
