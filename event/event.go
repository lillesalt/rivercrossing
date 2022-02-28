package event

/*
func PutInn(item string) string {
	return "[ Rev Korn ---V \\ \\_HS+Kylling_/ _____________/ Ø---]"
}

func CrossRiver(item string) string {
	return "[ Rev Korn ---V \\ _____________\\_HS+Kylling_/ / Ø---]"
}

func TaUt(item string) string {
	return "[ Rev Korn ---V \\ _____________\\_HS_/ / Ø--- Kylling]"
}
*/

var tilstand string 
var items = [] string {
	"Kylling",	"Rev",	"Korn",
}

func Tilstander(tilstanden string) string{
	tilstand = tilstanden
	tilstand = "[ Kylling Rev Korn HS ---V \\____/___________/ Ø---]"
	return tilstanden
}

func FirstPut(item string) string {
	tilstand = "[ "
	for i:=0; i < len(items); i++{
		if item != items[i]{
			tilstand += items[i] + " "
		}
	}

	if (item == items[1]) {
		tilstand = ""+items[0]+ " spiser " + items[2]
	}

	if (item == items[2]) {
		tilstand = ""+items[1]+ " spiser " + items[0]
	}

	if (item == items[0]) {
		items = append(items[:0])
		tilstand += " ---V \\_HS+"+item+"_/___________/ Ø---]"
	}else{
		tilstand = "Error! Skriv Kylling"
	}

	return Tilstander(tilstand)

}


func GetInn(item string)string {
	tilstand = "[ Rev Korn ---V \\_HS+"+item+"_/___________/ Ø---]"
	return Tilstander(tilstand)
}

func CrossRiver(item string)string {
	tilstand = "[ Rev Korn ---V \\___________\\_HS+"+item+"___/ Ø---]"
	return Tilstander(tilstand)
}

func TaUt(item string)string {
	tilstand = "[ Rev Korn ---V \\___________\\_HS___/ Ø--- "+item+"]"
	return Tilstander(tilstand)
}

func GetOut(item string)string {
	tilstand = "[ Rev Korn ---V \\___________\\___/ Ø--- HS "+item+"]"
	return Tilstander(tilstand)
}


