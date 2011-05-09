package bayesian

type Data interface {
	GetWords() []string
	Is() bool
}
