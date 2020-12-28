package gomesh

import "testing"

func TestParse(t *testing.T) {
	res := Parse("data/testmesh.msh")
	if len(res.Nodes) != 6 {
		t.Errorf("Nodes count incorrect, got: %d, want: %d.", len(res.Nodes), 6)
	}
}