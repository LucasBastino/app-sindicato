package api

import (
	"net/http"

	i "github.com/LucasBastino/app-sindicato/src/api/interfaces"
	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) createParent(w http.ResponseWriter, r *http.Request) {
	parentParser := i.ParentParser{}
	newParent := parserCaller(parentParser, r)
	insertInDBCaller(newParent, c.DB)
	renderFileTemplateCaller(newParent, w, "src/views/files/parentFile.html")
}

func (c *Controller) deleteParent(w http.ResponseWriter, r *http.Request) {
	IdParent := getIdModelCaller(models.Parent{}, r)
	deleteFromDBCaller(models.Parent{IdParent: IdParent}, c.DB)
	c.renderParentTable(w, r)
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	parser := i.ParentParser{}
	parentEdited := parserCaller(parser, r)
	IdParent := getIdModelCaller(models.Parent{}, r)
	parentEdited.IdParent = IdParent
	updateInDBCaller(parentEdited, IdParent, c.DB)
	renderFileTemplateCaller(parentEdited, w, "src/views/files/parentFile.html")
}

func (c *Controller) searchParent(w http.ResponseWriter, r *http.Request) {
	parents := searchInDBCaller(models.Parent{}, r, c.DB)
	renderTableTemplateCaller(models.Parent{}, w, "src/views/tables/parentTable.html", parents)
}
