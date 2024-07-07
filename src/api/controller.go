package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/LucasBastino/app-sindicato/src/models"
	// "syscall/js"
)

type Controller struct {
	DB *sql.DB
}

// ------------------------------------

func (c *Controller) createMember(w http.ResponseWriter, r *http.Request) {
	memberParser := MemberParser{}
	newMember := parserCaller(memberParser, r)
	insertInDBCaller(newMember, c.DB)
	renderTemplateCaller(newMember, w, "src/views/files/memberFile.html")

	// http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderMemberList(w, r) // esto tambien funciona
}

func (c *Controller) deleteMember(w http.ResponseWriter, r *http.Request) {
	IdMember := getIdModel("Member", r)
	deleteFromDBCaller(models.Member{IdMember: IdMember}, c.DB)
	c.renderMemberTable(w, r)
}

func (c *Controller) editMember(w http.ResponseWriter, r *http.Request) {
	memberParser := MemberParser{}
	memberEdited := parserCaller(memberParser, r)
	IdMember := getIdModel("Member", r)
	updateInDBCaller(memberEdited, IdMember, c.DB)

	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

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

func (c *Controller) createEnterprise(w http.ResponseWriter, r *http.Request) {
	var enterprise models.Enterprise
	enterprise.Name = r.FormValue("name")
	enterprise.Address = r.FormValue("address")
	// parseEnterprise()

	insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s', '%s')", enterprise.Name, enterprise.Address))
	fmt.Println("error inserting data to database")
	if err != nil {
		DBError{"INSERT ENTERPRISE"}.Error(err)
	}
	defer insert.Close()

	path := "src/views/files/enterpriseFile.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, enterprise)

}

func (c *Controller) searchMember(w http.ResponseWriter, r *http.Request) {
	emptyMember := models.Member{}
	members := searchInDBCaller(emptyMember, r, c.DB)

}

func (c *Controller) searchEnterprise(w http.ResponseWriter, r *http.Request) {
	searchKey := r.FormValue("search-key")
	var enterprises []models.Enterprise
	var enterprise models.Enterprise

	result, err := c.DB.Query(fmt.Sprintf(`SELECT * FROM EnterpriseTable WHERE Name LIKE '%%%s%%' OR Address LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		DBError{"SELECT ENTERPRISE"}.Error(err)
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			ScanError{"ENTERPRISE"}.Error(err)
		}
		enterprises = append(enterprises, enterprise)
	}
	defer result.Close()

	path := "src/views/tables/enterpriseTable.html"
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		TmplError{path}.Error(err)
	}
	tmpl.Execute(w, enterprises)
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
