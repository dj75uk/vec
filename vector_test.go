package vec

import "testing"

func TestAdd(t *testing.T) {
	a := Vector3{6, 7, 8}
	b := Vector3{4, 3, -3}
	expected := Vector3{10, 10, 5}
	actual := Add(a, b)
	if actual != expected {
		t.Errorf("Add(%v, %v) = %v, want %v", a, b, actual, expected)
	}
}

func TestSub(t *testing.T) {
	a := Vector3{6, 7, 8}
	b := Vector3{-4, 7, 10}
	expected := Vector3{10, 0, -2}
	actual := Sub(a, b)
	if actual != expected {
		t.Errorf("Sub(%v, %v) = %v, want %v", a, b, actual, expected)
	}
}

func TestMul(t *testing.T) {
	a := Vector3{0, 7, -8}
	b := 2.0
	expected := Vector3{0, 14, -16}
	actual := Mul(a, b)
	if actual != expected {
		t.Errorf("Mul(%v, %v) = %v, want %v", a, b, actual, expected)
	}
}

func TestDiv(t *testing.T) {
	a := Vector3{0, 7, -8}
	b := 2.0
	expected := Vector3{0, 3.5, -4}
	actual := Div(a, b)
	if actual != expected {
		t.Errorf("Div(%v, %v) = %v, want %v", a, b, actual, expected)
	}
}

func TestMag(t *testing.T) {
	a := Vector3{4, 5, 6}
	expected := 8.774964387392123
	actual := Mag(a)
	if actual != expected {
		t.Errorf("Mag(%v) = %v, want %v", a, actual, expected)
	}
}

func TestZero(t *testing.T) {
	expected := Vector3{}
	actual := Zero()
	if actual != expected {
		t.Errorf("Zero(%v) = %v, want %v", expected, actual, expected)
	}
}

func TestFromArray(t *testing.T) {
	a := []float64{1, 2, 3}
	expected := Vector3{1, 2, 3}
	actual := FromArray(a)
	if actual != expected {
		t.Errorf("FromArray(%v) = %v, want %v", a, actual, expected)
	}
}

func TestToArray(t *testing.T) {
	a := Vector3{1, 2, 3}
	expected := []float64{1, 2, 3}
	actual := ToArray(a)
	if actual[0] != expected[0] || actual[1] != expected[1] || actual[2] != expected[2] {
		t.Errorf("ToArray(%v) = %v, want %v", a, actual, expected)
	}
}

func TestCpy(t *testing.T) {
	a := Vector3{1, 2, 3}
	expected := Vector3{1, 2, 3}
	actual := Cpy(a)
	if actual != expected {
		t.Errorf("Cpy(%v) = %v, want %v", a, actual, expected)
	}
}

func TestInv(t *testing.T) {
	a := Vector3{1, 2, 3}
	expected := Vector3{-1, -2, -3}
	actual := Inv(a)
	if actual != expected {
		t.Errorf("Inv(%v) = %v, want %v", a, actual, expected)
	}
}

func TestNrm(t *testing.T) {
	a := Vector3{1, 2, 4}
	expected := Vector3{0.2182178902359924, 0.4364357804719848, 0.8728715609439696}
	actual := Nrm(a)
	if actual != expected {
		t.Errorf("Nrm(%v) = %v, want %v", a, actual, expected)
	}
}
