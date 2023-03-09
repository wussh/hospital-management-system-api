package application

type IDGenerator interface {
	Generate() (string, int, error)
}
