package interfaces

import "github.com/LucasBastino/app-sindicato/src/models"

type TypeModel interface {
	models.Member | models.Parent | models.Enterprise
}
