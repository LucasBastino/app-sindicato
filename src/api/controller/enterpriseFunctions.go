package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createEnterprise(w http.ResponseWriter, r *http.Request) {
	enterpriseParser := i.EnterpriseParser{}
	newEnterprise := parserCaller(enterpriseParser, r)
	insertModelCaller(newEnterprise, c.DB)
	renderFileTemplateCaller(newEnterprise, w, "src/views/files/enterpriseFile.html")

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderEnterpriseList(w, r) // esto tambien funciona
}

func (c *Controller) deleteEnterprise(w http.ResponseWriter, r *http.Request) {
	IdEnterprise := getIdModelCaller(models.Enterprise{}, r)
	deleteModelCaller(models.Enterprise{IdEnterprise: IdEnterprise}, c.DB)
	c.renderEnterpriseTable(w, r)
}

func (c *Controller) editEnterprise(w http.ResponseWriter, r *http.Request) {
	parser := i.EnterpriseParser{}
	enterpriseEdited := parserCaller(parser, r)
	IdEnterprise := getIdModelCaller(models.Enterprise{}, r)
	enterpriseEdited.IdEnterprise = IdEnterprise
	editModelCaller(enterpriseEdited, IdEnterprise, c.DB)
	renderFileTemplateCaller(enterpriseEdited, w, "src/views/files/enterpriseFile.html")
}

func (c *Controller) renderEnterpriseTable(w http.ResponseWriter, r *http.Request) {
	enterprises := searchModelsCaller(models.Enterprise{}, r, c.DB)
	renderTableTemplateCaller(models.Enterprise{}, w, "src/views/tables/enterpriseTable.html", enterprises)
}

func (c *Controller) renderEnterpriseFile(w http.ResponseWriter, r *http.Request) {
	enterprise := searchOneModelByIdCaller(models.Enterprise{}, r, c.DB)
	renderFileTemplateCaller(enterprise, w, "src/views/files/enterpriseFile.html")
}

func (c *Controller) renderCreateEnterpriseForm(w http.ResponseWriter, r *http.Request) {
	renderCreateModelFormCaller(models.Enterprise{}, w, "src/views/forms/createEnterpriseForm.html")
}
