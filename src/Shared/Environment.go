package Shared

type Environment string

const (
	Local      Environment = "local"
	Develop    Environment = "dev"
	Staging    Environment = "stg"
	Production Environment = "prd"
)

func NewEnv(e string) Environment {
	return Environment(e)
}

func (e Environment) IsProduction() bool {
	return e == Production
}
