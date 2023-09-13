package spielfeld

import ("schiffe"
		. "zufallszahlen")

type Spielfeld interface{

	String () string
	
	SetzeFeld (x,y int, wert uint16)

	Draw()
	
	GibGroesse () uint16
	
	SchiffeSetzen (x,y uint16,c schiffe.Schiff)
	
	SetzeX_Verschiebung (x uint16)
	
	GibX_Verschiebung () uint16
	
	Feld (x,y uint16) (kx,ky int)
	
	Umranden()
	
	Removefive ()
	
	Bombardieren (x,y uint16) (kx,ky int, auffeld bool)
	
	BombeEintragen (x,y int)
	
	GehoertPunktzuSpielfeld (x,y uint16) bool
	
	GameOver() bool
	
	AufSpielfeld (x,y uint16) bool
	
	Schiffpasst (x,y uint16, c schiffe.Schiff) bool
	
	StringToFeld (sfeld string)
	
	SetzeBesitzer (name string)
	
	RandomBomb ()
	
	ProzentBomb (prozent int)
}


func RandomFeld () Spielfeld {
	var s Spielfeld
	s = New(10)
	var yard []schiffe.Schiff
	yard = schiffe.BuildShipyard()
	for i:=0;i<len(yard);i++{
		nochmal:
		orientierung:=int(Zufallszahl(0,1))
		if orientierung == 0 {
			yard = schiffe.SwitchYard(yard)
		}
		x:=uint16(Zufallszahl(int64(s.GibX_Verschiebung()),int64(s.GibX_Verschiebung()+75*s.GibGroesse())))
		y:=uint16(Zufallszahl(0,int64(75*s.GibGroesse())))
		if s.GehoertPunktzuSpielfeld(x,y) {
			if s.Schiffpasst(x,y,yard[i]) {
				s.SchiffeSetzen(x,y,yard[i])
			} else {
				goto nochmal
				}
		}
	}
	return s
}

				
