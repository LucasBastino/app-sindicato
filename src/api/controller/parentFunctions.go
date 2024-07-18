package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createParent(w http.ResponseWriter, r *http.Request) {
	parentParser := i.ParentParser{}
	newParent := parserCaller(parentParser, r)
	insertModelCaller(newParent, c.DB)
	renderFileTemplateCaller(newParent, w, "src/views/files/parentFile.html")
}

func (c *Controller) deleteParent(w http.ResponseWriter, r *http.Request) {
	IdParent := getIdModelCaller(models.Parent{}, r)
	deleteModelCaller(models.Parent{IdParent: IdParent}, c.DB)
	c.renderParentTable(w, r)
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	parser := i.ParentParser{}
	parentEdited := parserCaller(parser, r)
	IdParent := getIdModelCaller(models.Parent{}, r)
	parentEdited.IdParent = IdParent
	editModelCaller(parentEdited, IdParent, c.DB)
	renderFileTemplateCaller(parentEdited, w, "src/views/files/parentFile.html")
}

func (c *Controller) renderParentTable(w http.ResponseWriter, r *http.Request) {
	parents := searchModelsCaller(models.Parent{}, r, c.DB)
	renderTableTemplateCaller(models.Parent{}, w, "src/views/tables/parentTable.html", parents)
}

func (c *Controller) renderParentFile(w http.ResponseWriter, r *http.Request) {
	parent := searchOneModelByIdCaller(models.Parent{}, r, c.DB)
	renderFileTemplateCaller(parent, w, "src/views/files/parentFile.html")
}

func (c *Controller) renderCreateParentForm(w http.ResponseWriter, r *http.Request) {
	renderCreateModelFormCaller(models.Parent{}, w, "src/views/forms/createParentForm.html")
}
