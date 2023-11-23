package language

import (
	"fmt"
)

type Language string

const (
	CPP  = Language("cpp")
	Go   = Language("go")
	Rust = Language("rs")
)

func New(language string) (Language, error) {
	switch language {
	case string(CPP):
		return "", fmt.Errorf("cpp is not supported yet")
	case string(Go):
		return Go, nil
	case string(Rust):
		return "", fmt.Errorf("rust is not supported yet")
	default:
		return "", fmt.Errorf("invalid language: %s", language)
	}
}
