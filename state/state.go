package main

/*
func ViewState() string {
	return "[ Kylling Rev Korn ---V \\ \\_HS_/ _____________/ Ø---]"
}
*/

var tilstand string 

func ViewState()string{
	return Tilstander()
}

func Tilstander()string{
	tilstand = "[ Kylling Rev Korn ---V \\ \\_HS_/ _____________/ Ø---]"
	return tilstand
}
