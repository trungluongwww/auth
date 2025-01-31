package crypto

import "golang.org/x/crypto/bcrypt"

const (
	Cost = 8
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), Cost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CompareHashAndPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
