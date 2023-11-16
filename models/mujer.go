package models

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar()    { m.respira = true }
func (m *Mujer) Comer()       { m.come = true }
func (m *Mujer) Pensar()      { m.piensa = true }
func (m *Mujer) Sexo() string { return "Mujer" }
