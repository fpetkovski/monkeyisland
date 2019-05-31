package ghosts

import (
	"encoding/json"
	"fpetkovski/monkeyisland/pkg/domain/ghosts"
	"net/http"
)

type Handler struct {
}

func NewGhostsHandler() Handler {
	return Handler{}
}

func (handler Handler) GetGhosts(w http.ResponseWriter, r *http.Request) {
	items := ghosts.GenerateGhosts()
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		panic(err)
	}
}
