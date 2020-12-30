package gomesh

import "testing"

func TestParse22Ascii(t *testing.T) {
	res, err := Parse("data/2.2_ascii.msh")
	if err != nil {
		t.Errorf("Err during parse: %v", err)
	}
	if len(res.Nodes) != 891 {
		t.Errorf("Nodes count incorrect, got: %d, want: %d.", len(res.Nodes), 891)
	}
	if len(res.Elements) != 2748 {
		t.Errorf("Nodes count incorrect, got: %d, want: %d.", len(res.Elements), 2748)
	}
	el := res.Elements[len(res.Elements)-1] //TODO: get Elements by tags and improve next test
	if len(el.Data) == 1 && el.Data[0] == "4.28717" {
	} else {
		t.Errorf("El Data incorrect, got: %v.", el.Data)
	}
}
