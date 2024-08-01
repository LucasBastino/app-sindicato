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
	errorMap := validateFieldsCaller(enterprise, r)
	parser := i.EnterpriseParser{}
	enterprise = parserCaller(parser, r)
	if len(errorMap) > 0 {
		templateData := createTemplateDataCaller(enterprise, enterprise, nil, "src/views/files/enterpriseFile.html", errorMap)
		c.RenderHTML(w, templateData)
	} else {
		insertModelCaller(enterprise, c.DB)
		templateData := createTemplateDataCaller(enterprise, enterprise, nil, "src/views/files/enterpriseFile.html", errorMap)
		c.RenderHTML(w, templateData)
	}
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
	// le paso un enterprise vacio para que los campos del form aparezcan vacios
	templateData := createTemplateDataCaller(enterprise, models.Enterprise{}, nil, "src/views/forms/createEnterpriseForm.html", nil)
	c.RenderHTML(w, templateData)
}
