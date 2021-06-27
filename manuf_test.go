package manuf

import "testing"

func Test_manuf(t *testing.T) {
	Init()
	t.Log(Search("00:50:56:f0:a1:95"))
}
