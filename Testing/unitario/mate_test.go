package unitario

import "testing"

/*
func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d y se esperaba %d", total, 10)
	}
}
*/

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{5, 5, 10},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d y se esperaba %d", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{5, 5, 5},
		{25, 24, 25},
	}

	for _, item := range tabla {
		total := GetMax(item.a, item.b)

		if total != item.c {
			t.Errorf("Máximo incorrecto, tiene %d y se esperaba %d", total, item.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{45, 1134903170},
	}

	for _, item := range tabla {
		total := Fibonacci(item.a)

		if total != item.b {
			t.Errorf("Fibonacci incorrecto, tiene %d y se esperaba %d", total, item.b)
		}
	}
}
