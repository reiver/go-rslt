package rslt


type Result interface {
	Result() (interface{}, error, Warning)
}
