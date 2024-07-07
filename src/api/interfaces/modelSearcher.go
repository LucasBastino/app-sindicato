package interfaces

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type ModelSearcher interface {
	SearchModel(*http.Request, *sql.DB) []*IModel
}

type MemberSearcher struct{}

func (m MemberSearcher) SearchModel(r *http.Request, DB *sql.DB) []models.Member {
	searchKey := r.FormValue("search-key")
	var members []*models.Member
	var member models.Member

	result, err := DB.Query(fmt.Sprintf(`SELECT * FROM MemberTable WHERE Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching member in DB")
	}
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning member")
		}
		members = append(members, &member)
	}
	defer result.Close()
	return members
}

// probar sacar esta interface y hacerlo con generics, igualmente primero probar asi que ya esta hecho
