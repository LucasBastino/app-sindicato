package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func getIdModel(model string, r *http.Request) int {
	IdStr := r.PathValue(fmt.Sprintf("Id%s", model))
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		fmt.Println("error converting type")
		panic(err)
	}
	return Id
}
