package obserser

import (
	"testing"
)

func TestObserver(t *testing.T) {
	subject := NewSubject()
	reader1 := NewReader("r1")
	reader2 := NewReader("r2")
	reader3 := NewReader("r3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContent("2022")
}
