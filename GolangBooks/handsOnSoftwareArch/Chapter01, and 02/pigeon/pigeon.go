// Package pigeon is used to demonstrate visibily and scope for packages
package pigeon


type Pigeon struct {
	Name     string
	featherLength int
}

func (p *Pigeon) GetFeatherLength() int {
	return p.featherLength
}


func (p *Pigeon) SetFeatherLength(length int)  {
	p.featherLength = length
}
