package drivens

type ForTokenManager interface {
	GenerateToken(userID string, role string) (string, error)

	GenerateRefreshToken(userID string) (string, error)

	ValidateToken(tokenString string) (string, error)

	InvalidateToken(token string) error
}
