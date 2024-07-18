package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type ModelParser[tM TypeModel] interface {
	ParseModel(*http.Request) tM
}

type MemberParser struct{}

func (m MemberParser) ParseModel(r *http.Request) models.Member {
	member := models.Member{}
	member.Name = r.FormValue("name")
	member.DNI = r.FormValue("dni")
	return member
}

type ParentParser struct{}

func (p ParentParser) ParseModel(r *http.Request) models.Parent {
	parent := models.Parent{}
	parent.Name = r.FormValue("name")
	parent.Rel = r.FormValue("rel")
	IdMemberStr := r.FormValue("idmember")
	IdMember, err := strconv.Atoi(IdMemberStr)
	if err != nil {
		fmt.Println("error converting IdMemberStr to int")
		panic(err)
	}
	parent.IdMember = IdMember

	return parent
}

type EnterpriseParser struct{}

func (p EnterpriseParser) ParseModel(r *http.Request) models.Enterprise {
	enterprise := models.Enterprise{}
	enterprise.Name = r.FormValue("name")
	enterprise.Address = r.FormValue("address")
	return enterprise
}
