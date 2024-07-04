package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func parseMember(r *http.Request) models.Member {
	var member models.Member
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}

func getIdModel(model string, r *http.Request) int {
	IdStr := r.PathValue(fmt.Sprintf("Id%s", model))
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		fmt.Println("error converting type")
		panic(err)
	}
	return Id
}
