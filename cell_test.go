package mazer

import "testing"

func TestLinkBidi(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(1, 1)

	c1.Link(c2, true)

	if len(c1.GetLinks()) != 1 {
		t.Fatalf("Expected only a single link!")
	}

	if !c1.IsLinked(c2) || !c2.IsLinked(c1) {
		t.Fatalf("Expected %+v to be linked to %+v", c1, c2)
	}
}

func TestLinkUnidi(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(1, 1)

	c1.Link(c2, false)

	if len(c1.GetLinks()) != 1 {
		t.Fatalf("Expected only a single link")
	}

	if len(c2.GetLinks()) != 0 {
		t.Fatalf("Expected no links")
	}

	if !c1.IsLinked(c2) || c2.IsLinked(c1) {
		t.Fatalf("Expected %+v to be linked to %+v, but not the other way", c1, c2)
	}

}

func TestUnlinkBidi(t *testing.T) {
	c1 := NewCell(0, 0)
	c2 := NewCell(1, 1)

	c1.Link(c2, true)
	c1.Unlink(c2, true)

	if len(c1.GetLinks()) != 0 {
		t.Fatalf("Expected no links")
	}

	if c1.IsLinked(c2) || c2.IsLinked(c1) {
		t.Fatalf("Expected %+v to be unlinked from %+v, and vice-versa", c1, c2)
	}
}
