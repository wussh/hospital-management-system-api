package application

type PasswordHash interface {
	Hash(password string) (string, int, error)
	ComparePassword(plain string, encrypted string) (int, error)
}
