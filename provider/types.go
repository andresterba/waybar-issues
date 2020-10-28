package provider

type Provider interface {
	Process() error
	GetFormatedOutput() string
}
