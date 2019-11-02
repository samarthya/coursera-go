package trunc

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World!"

	if got := displayHelloWorld(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
