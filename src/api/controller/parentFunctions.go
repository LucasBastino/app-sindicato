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
	errorMap := validateFieldsCaller(parent, r)
	parser := i.ParentParser{}
	parent = parserCaller(parser, r)
	if len(errorMap) > 0 {
		templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", errorMap)
		c.RenderHTML(w, templateData)
	} else {
		insertModelCaller(parent, c.DB)
		templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", errorMap)
		c.RenderHTML(w, templateData)
	}
}

func (c *Controller) deleteParent(w http.ResponseWriter, r *http.Request) {
	parent = searchOneModelByIdCaller(parent, r, c.DB)
	deleteModelCaller(parent, c.DB)
	// renderiza de nuevo la tabla
	// falla aca abajo, fijarse. Tambien probar agregar familiares
	member = searchOneModelByIdCaller(member, r, c.DB)
	c.renderMemberParents(w, r)
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	parentEdited := parserCaller(parentParser, r)
	IdParent := getIdModelCaller(parent, r)
	parentEdited.IdParent = IdParent
	editModelCaller(parentEdited, c.DB)
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderParentTable(w http.ResponseWriter, r *http.Request) {
	parents := searchModelsCaller(parent, r, c.DB)
	templateData := createTemplateDataCaller(parent, parent, parents, "src/views/tables/parentTable.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderParentFile(w http.ResponseWriter, r *http.Request) {
	parent := searchOneModelByIdCaller(parent, r, c.DB)
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/files/parentFile.html", nil)
	c.RenderHTML(w, templateData)
}

func (c *Controller) renderCreateParentForm(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModelCaller(member, r)
	// creo un parent nuevo para que el form aparezca con campos vacios
	parent = models.Parent{}
	parent.IdMember = IdMember
	templateData := createTemplateDataCaller(parent, parent, nil, "src/views/forms/createParentForm.html", nil)
	c.RenderHTML(w, templateData)
}
