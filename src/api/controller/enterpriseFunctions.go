package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createEnterprise(w http.ResponseWriter, r *http.Request) {
	EnterpriseParser := i.EnterpriseParser{}
	newEnterprise := parserCaller(EnterpriseParser, r)
	insertInDBCaller(newEnterprise, c.DB)
	renderFileTemplateCaller(newEnterprise, w, "src/views/files/enterpriseFile.html")

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderEnterpriseList(w, r) // esto tambien funciona
}

func (c *Controller) deleteEnterprise(w http.ResponseWriter, r *http.Request) {
	IdEnterprise := getIdModelCaller(models.Enterprise{}, r)
	deleteFromDBCaller(models.Enterprise{IdEnterprise: IdEnterprise}, c.DB)
	c.renderEnterpriseTable(w, r)
}

func (c *Controller) editEnterprise(w http.ResponseWriter, r *http.Request) {
	parser := i.EnterpriseParser{}
	enterpriseEdited := parserCaller(parser, r)
	IdEnterprise := getIdModelCaller(models.Enterprise{}, r)
	updateInDBCaller(enterpriseEdited, IdEnterprise, c.DB)

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func (c *Controller) searchEnterprise(w http.ResponseWriter, r *http.Request) {
	enterprises := searchInDBCaller(models.Enterprise{}, r, c.DB)
	renderTableTemplateCaller(models.Enterprise{}, w, "src/views/tables/enterpriseTable.html", enterprises)
}
