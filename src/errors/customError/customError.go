package customError

type CustomError struct {
	Code      int
	Msg       string
	ClientMsg string
}

func (customError CustomError) Error() string {
	return customError.ClientMsg
}

var QueryError = CustomError{
	Code:      500,
	ClientMsg: "error interno de la base de datos",
}
var ScanError = CustomError{
	Code:      500,
	ClientMsg: "error interno",
}
var FormatError = CustomError{
	Code:      500,
	ClientMsg: "error interno",
}

var ValidationError = CustomError{
	Code:      400,
	ClientMsg: "error de validacion",
}
var UnauthorizedError = CustomError{
	Code:      500,
	ClientMsg: "error de autorizacion",
}
var InsufficientPermisionsError = CustomError{
	Code:      500,
	ClientMsg: "permisos insuficientes",
}
var InternalServerError = CustomError{
	Code:      500,
	ClientMsg: "error interno",
}
var StrConvError = CustomError{
	Code:      500,
	ClientMsg: "error interno",
}
var DatabaseError = CustomError{
	Code:      500,
	ClientMsg: "error interno de la base de datos",
}
var ParamsError = CustomError{
	Code:      500,
	ClientMsg: "error interno",
}
var FileError = CustomError{
	Code:      500,
	ClientMsg: "error interno",
}