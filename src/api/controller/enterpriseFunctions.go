package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

var (
	enterprise       models.Enterprise
	enterpriseParser i.EnterpriseParser
)

func (c *Controller) createEnterprise(w http.ResponseWriter, r *http.Request) {
	newEnterprise := parserCaller(enterpriseParser, r)
	insertModelCaller(newEnterprise, c.DB)
	templateData := createTemplateDataCaller(enterprise, newEnterprise, nil, "src/views/files/enterpriseFile.html", nil)
	c.RenderHTML(w, templateData)

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderEnterpriseList(w, r) // esto tambien funciona
}

func (c *Controller) deleteEnterprise(w http.ResponseWriter, r *http.Request) {
	IdEnterprise := getIdModelCaller(enterprise, r)
	enterprise.IdEnterprise = IdEnterprise
	deleteModelCaller(enterprise, c.DB)
	c.renderEnterpriseTable(w, r)
}

func (c *Controller) editEnterprise(w http.ResponseWriter, r *http.Request) {
	enterpriseEdited := parserCaller(enterpriseParser, r)
	IdEnterprise := getIdModelCaller(enterprise, r)
	enterpriseEdited.IdEnterprise = IdEnterprise
	editModelCaller(enterpriseEdited, c.DB)
	templateData := createTemplateDataCaller(enterprise, enterpriseEdited, nil, "src/views/files/enterpriseFile.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderEnterpriseTable(w http.ResponseWriter, r *http.Request) {
	enterprises := searchModelsCaller(enterprise, r, c.DB)
	templateData := createTemplateDataCaller(enterprise, enterprise, enterprises, "src/views/tables/enterpriseTable.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderEnterpriseFile(w http.ResponseWriter, r *http.Request) {
	enterprise := searchOneModelByIdCaller(enterprise, r, c.DB)
	templateData := createTemplateDataCaller(enterprise, enterprise, nil, "src/views/files/enterpriseFile.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderCreateEnterpriseForm(w http.ResponseWriter, r *http.Request) {
	templateData := createTemplateDataCaller(enterprise, enterprise, nil, "src/views/forms/createEnterpriseForm.html", nil)
	c.RenderHTML(w, templateData)
}
