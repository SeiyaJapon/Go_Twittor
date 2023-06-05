package models

type Woman struct {
	Man
}

func (woman *Woman) Breath() {
	woman.breathing = true
}

func (woman *Woman) Eat() {
	woman.eating = true
}

func (woman *Woman) Think() {
	woman.thinking = true
}

func (woman *Woman) Sex() string {
	return "Mujer"
}
