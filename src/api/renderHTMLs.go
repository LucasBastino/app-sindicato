package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/home.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	// tmpl := createTemplate("src/views/createMemberForm.html")
	// execTemplate(w, nil, tmpl, "createMemberForm.html")

	// tmpl, _ := template.New("createMemberForm.html").ParseFiles("src/views/createMemberForm.html")
	// tmpl.Execute(w, nil)

	tmpl, _ := template.ParseFiles("src/views/createMemberForm.html", "src/views/footer.html")
	// el primero siempre es el main template, los demas se usan como componentes
	tmpl.Execute(w, nil)
}

func (c *Controller) renderMemberList(w http.ResponseWriter, r *http.Request) {
	result, err := c.DB.Query("SELECT IdMember, Name, DNI FROM MemberTable")
	if err != nil {
		fmt.Println("error obtaining data from database")
		log.Panic(err)
	}

	var members []models.Member
	for result.Next() {
		member := models.Member{}
		err := result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning data")
			log.Panic(err)
		}
		members = append(members, member)
	}

	tmpl, err := template.ParseFiles("src/views/memberList.html", "src/views/footer.html")
	if err != nil {
		fmt.Println("error parsing files")
		panic(err)
	}

	tmpl.Execute(w, members)
}

func (c *Controller) renderMemberList2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/views/memberList2.html", "src/views/footer.html")
	if err != nil {
		fmt.Println("error parsing file memberList2")
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func (c *Controller) getList(w http.ResponseWriter, r *http.Request) {
	result, err := c.DB.Query("SELECT IdMember, Name, DNI FROM MemberTable")
	if err != nil {
		fmt.Println("error obtaining data from database")
		log.Panic(err)
	}

	var members []models.Member
	for result.Next() {
		member := models.Member{}
		err := result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning data")
			log.Panic(err)
		}
		members = append(members, member)
	}
	// file, err := os.ReadFile("src/views/getList.html")
	// if err != nil {
	// 	fmt.Println("error reading file getList.html")
	// 	panic(err)
	// }
	// strFile := string(file)
	// tmpl := template.Must(template.New("listTmpl").Parse(strFile))
	// tmpl.Execute(w, members)
	tmpl := returnHtmlTemplate("src/views/getList.html")
	tmpl.Execute(w, members)
}

// func returnList() string {
// 	return `
//     <table class="table ">
//         <tr>
//             <th>Afiliado</th>
//             <th>DNI</th>
//         </tr>

//         {{range $_, $member := .}}
//         <tr>
//             <td>{{$member.Name}} </td>
//             <td>{{$member.DNI}} </td>

//             <td><form action="/member/{{$member.IdMember}}/delete" method="DELETE"><button type="submit" onclick="jsFunc()" class="btn-danger">Eliminar</button></td>

//             </form>     </tr>
//         {{end}}

//     </table>`
// }
