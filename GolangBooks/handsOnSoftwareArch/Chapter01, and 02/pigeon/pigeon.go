// Package pigeon is used to demonstrate visibily and scope for packages
package pigeon

//Name is a public attribute, and the code outside of the package can reference it.
//featherLength is a private attribute, and the code outside of the package
//cannot reference it.

type Pigeon struct {
	Name          string
	featherLength int
}

func (p *Pigeon) GetFeatherLength() int {
	return p.featherLength
}

func (p *Pigeon) SetFeatherLength(length int) {
	p.featherLength = length
}
