package api

import (
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func parseMember(r *http.Request) models.Member {
	var member models.Member
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}
