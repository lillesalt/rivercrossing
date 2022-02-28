
package event

import "testing"


func TestPutInn(t *testing.T) {
	wanted := "[ Rev Korn  ---V \\_HS+Kylling_/___________/ Ø---]"
	got := FirstPut("Kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestGetInn(t *testing.T) {
	wanted := "[ Rev Korn ---V \\_HS+Kylling_/___________/ Ø---]"
	got := GetInn("Kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}


func TestCrossRiver(t *testing.T) {
	wanted := "[ Rev Korn ---V \\___________\\_HS+Kylling___/ Ø---]"
	got := CrossRiver("Kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestTaUt(t *testing.T) {
	wanted := "[ Rev Korn ---V \\___________\\_HS___/ Ø--- Kylling]"
	got := TaUt("Kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestGetOut(t *testing.T) {
	wanted := "[ Rev Korn ---V \\___________\\___/ Ø--- HS Kylling]"
	got := GetOut("Kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
