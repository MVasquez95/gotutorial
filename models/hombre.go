package models

type Hombre struct {
	edad    int
	altura  float32
	peso    float32
	respira bool
	piensa  bool
	come    bool
}

func (h *Hombre) Respirar()    { h.respira = true }
func (h *Hombre) Comer()       { h.come = true }
func (h *Hombre) Pensar()      { h.piensa = true }
func (h *Hombre) Sexo() string { return "Hombre" }
func (h *Hombre) Vive() bool   { return true }
