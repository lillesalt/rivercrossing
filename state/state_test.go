package main

import ("testing")

func TestViewState(t *testing.T) {
	wanted := "[ Kylling Rev Korn ---V \\ \\_HS_/ _____________/ Ø---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
