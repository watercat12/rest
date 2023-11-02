package get

import (
	"net/http"

	"github.com/watercat12/rest/examples/stdlib/models"
)

type TopicsGetResponse struct {
	Topics []models.Topic `json:"topics"`
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
