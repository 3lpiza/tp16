package main

import (
	"testing"
)

func TestSuma(t *testing.T) {
	resultado := suma(5, 3)
	if resultado != 8 {
		t.Errorf("La suma esperada era 8, pero se obtuvo %f", resultado)
	}
}

func TestResta(t *testing.T) {
	resultado := resta(10, 4)
	if resultado != 6 {
		t.Errorf("La resta esperada era 6, pero se obtuvo %f", resultado)
	}
}

func TestMultiplicacion(t *testing.T) {
	resultado := multiplicacion(6, 5)
	if resultado != 30 {
		t.Errorf("La multiplicación esperada era 30, pero se obtuvo %f", resultado)
	}
}

func TestDivision(t *testing.T) {
	resultado := division(15, 3)
	if resultado != 5 {
		t.Errorf("La división esperada era 5, pero se obtuvo %f", resultado)
	}
}
