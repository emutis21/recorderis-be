package drivens

type ForPasswordManager interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) error
}
