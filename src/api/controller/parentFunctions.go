package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

var (
	parent       models.Parent
	parentParser i.ParentParser
)

func (c *Controller) createParent(w http.ResponseWriter, r *http.Request) {
	newParent := parserCaller(parentParser, r)
	insertModelCaller(newParent, c.DB)
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	renderFileTemplateCaller(newParent, w, templateData)
}

func (c *Controller) deleteParent(w http.ResponseWriter, r *http.Request) {
	IdParent := getIdModelCaller(parent, r)
	parent.IdParent = IdParent
	deleteModelCaller(parent, c.DB)
	c.renderParentTable(w, r)
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	parentEdited := parserCaller(parentParser, r)
	IdParent := getIdModelCaller(parent, r)
	parentEdited.IdParent = IdParent
	editModelCaller(parentEdited, IdParent, c.DB)
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	renderFileTemplateCaller(parentEdited, w, templateData)
}

func (c *Controller) renderParentTable(w http.ResponseWriter, r *http.Request) {
	parents := searchModelsCaller(parent, r, c.DB)
	templateData := createTemplateDataCaller(parent, parent, parents, "src/views/tables/parentTable.html", nil)
	renderTableTemplateCaller(parent, w, templateData)
}

func (c *Controller) renderParentFile(w http.ResponseWriter, r *http.Request) {
	parent := searchOneModelByIdCaller(parent, r, c.DB)
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	renderFileTemplateCaller(parent, w, templateData)
}

func (c *Controller) renderCreateParentForm(w http.ResponseWriter, r *http.Request) {
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/forms/createParentForm.html", nil)
	renderCreateModelFormCaller(parent, w, templateData)
}
