package main

import ("spielfeld"
		"fmt"
		. "gfx2"
		"schiffe"
		"zufallszahlen"
		)
		
func main () {
	zufallszahlen.Randomisieren()
	var yard []schiffe.Schiff
	yard = schiffe.BuildShipyard()
	var s spielfeld.Spielfeld //Feldgroesse 600x600
	s = spielfeld.New(10)
	s.SetzeBesitzer("test")
	Fenster(((2*(s.GibGroesse())+1)*75),s.GibGroesse()*75)
	s.SetzeX_Verschiebung((s.GibGroesse()+1)*75)
	var c schiffe.Schiff
	var x, y uint16
	var taste uint8
	var status int8
	for  {
		taste,status,x,y = MausLesen1 ()
		for i:=0;i<len(yard);i++{
			if yard[i].GehoertPunktzuSchiff(x,y) {
				if taste==1 && status ==1 {
					for j:=0;j<len(yard);j++{
							if yard[j].Gibhighlight()  {
								yard[j].SwitchHighlighting()
							}
					}
					yard[i].SwitchHighlighting()
					c = yard[i]
				}
				
			}
			
			
		}
		if taste ==3 && status ==1{
					yard = schiffe.SwitchYard(yard)
		}
		
		if schiffe.Highlight(yard) {
			if s.AufSpielfeld(x,y) && status==1 && taste == 1 && s.Schiffpasst(x,y,c) {
			//Spielfeldtest - einfügenüberarbeiten 
				yard = schiffe.SchiffGeben(yard)
					if taste==1 && status == 1 {
						s.SchiffeSetzen(x,y,c)
						s.Umranden()
					}
			}
		}
		UpdateAus()
		Stiftfarbe(255,255,255)
		Cls()
		schiffe.DrawShipyard(yard)
		s.Draw()
		UpdateAn()
		if schiffe.YardLeer(yard){
			break
		}
	}
	s.Removefive()
	s.Draw()
	
	
	//fmt.Println(s)
	var s2 spielfeld.Spielfeld
	s2 = spielfeld.New(10)
	s2 = spielfeld.RandomFeld()
	s2.Removefive()
	fmt.Println("")
	//fmt.Println(s2)
	s2.Draw()
	for {
		taste,status,mx,my:=MausLesen1()
		if taste==1 && status==1 {
			fx,fy,auffeld:=s2.Bombardieren(mx,my)
			if auffeld {
				s2.BombeEintragen(fx,fy)
				s2.Draw()
				if s2.GameOver(){
					break
				}
				s.ProzentBomb (25)
				s.Draw()
				if s.GameOver(){
					break
				}
			}
		}	
	}
	fmt.Println("Game Over")
	if s.GameOver() {
		fmt.Println("Der Computer hat gewonnen!")
	} else {
		fmt.Println("Du hast gewonnen!")
	}
	TastaturLesen1()
}
