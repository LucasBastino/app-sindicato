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
	IdParent := getIdModel("Parent", r)
	deleteFromDBCaller(models.Parent{IdParent: IdParent}, c.DB)
	c.renderParentTable(w, r)
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	IdParent := getIdModel("Parent", r)
	parser := i.ParentParser{}
	parent := parserCaller(parser, r)
	updateInDBCaller(parent, IdParent, c.DB)
	renderFileTemplateCaller(parent, w, "src/views/files/parentFile.html")
}

func (c *Controller) searchParent(w http.ResponseWriter, r *http.Request) {
	parents := searchInDBCaller(models.Parent{}, r, c.DB)
	renderTableTemplateCaller(models.Parent{}, w, "src/views/tables/parentTable.html", parents)

}
