package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	IdParent := r.PathValue("IdParent")
	Name := r.FormValue("name")
	Rel := r.FormValue("rel")

	update, err := c.DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%s'", Name, Rel, IdParent))
	if err != nil {
		DBError{"UPDATE PARENT"}.Error(err)
	}
	update.Close()

	c.renderParentFile(w, r)
}

func (c *Controller) searchParent(w http.ResponseWriter, r *http.Request) {
	searchKey := r.FormValue("search-key")
	var parents []models.Parent
	var parent models.Parent

	result, err := c.DB.Query(fmt.Sprintf(`SELECT * FROM ParentTable WHERE Name LIKE '%%%s%%' OR Rel LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		DBError{"SELECT PARENT"}.Error(err)
	}
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel, &parent.IdMember)
		if err != nil {
			ScanError{"PARENT"}.Error(err)
		}
		parents = append(parents, parent)
	}
	defer result.Close()

	path := "src/views/tables/allParentsTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, parents)

}
