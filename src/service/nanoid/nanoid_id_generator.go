package nanoid

import (
	"net/http"
	"strings"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
	"github.com/aidarkhanov/nanoid"
)

type nanoidIDGenerator struct {
}

func NewNanoidIDGenerator() application.IDGenerator {
	return &nanoidIDGenerator{}
}

func (n *nanoidIDGenerator) Generate() (string, int, error) {
	id, err := nanoid.Generate(nanoid.DefaultAlphabet, 5)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return strings.ToLower(id), http.StatusOK, nil
}
