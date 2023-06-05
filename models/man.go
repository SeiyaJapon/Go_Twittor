package models

type Man struct {
	age       int
	high      float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	alife     bool
}

func (man *Man) Breath() {
	man.breathing = true
}

func (man *Man) Eat() {
	man.eating = true
}

func (man *Man) Think() {
	man.thinking = true
}

func (man *Man) Sex() string {
	return "Hombre"
}
