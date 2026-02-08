package koap

type SOAPOperation interface {
	Name() string
	SOAPAction() string
	BindingType() string
}
