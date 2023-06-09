package organizationMember

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Controller struct {
	Config                    *config.Config
	OrganizationMemberService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:                    cfg,
		OrganizationMemberService: svc,
	}
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response := c.OrganizationMemberService.RetrieveById(id)
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
