package schiffe 

//import "fmt"

type Schiff interface {
	SetzeKoordinaten (x,y uint16)
	
	GehoertPunktzuSchiff (x,y uint16) bool 
	
	Draw()
	
	GibOrientierung() bool 

	GibGroesse() uint16
	
	SetzteGroesse (groesse uint16)
	
	SwitchOrientierung()
	
	SwitchHighlighting()
	
	Gibhighlight() bool
	
	SetzeDummy(dummy bool)
	
	GibDummy () bool
	
}

func BuildShipyard () []Schiff {
	var s1,s2,s3,s4,s5,s6,s7,s8,s9,s10 Schiff
	

		//1x 4er, 2x 3er, 3x 2er, 4x 1er
		
		s1 = New()
		s1.SwitchOrientierung()
		s1.SetzteGroesse(4)
		s1.SetzeKoordinaten(50,50)
		s2 = New()
		s2.SwitchOrientierung()
		s2.SetzteGroesse(3)
		s2.SetzeKoordinaten(145,50)
		s3 = New()
		s3.SwitchOrientierung()
		s3.SetzteGroesse(3)
		s3.SetzeKoordinaten(145,245)
		s4 = New()
		s4.SwitchOrientierung()
		s4.SetzteGroesse(2)
		s4.SetzeKoordinaten(240,50)
		s5 = New()
		s5.SwitchOrientierung()
		s5.SetzteGroesse(2)
		s5.SetzeKoordinaten(240,200)
		s6 = New()
		s6.SwitchOrientierung()
		s6.SetzteGroesse(2)
		s6.SetzeKoordinaten(240,350)
		s7 = New()
		s7.SwitchOrientierung()
		s7.SetzteGroesse(1)
		s7.SetzeKoordinaten(335,50)
		s8 = New()
		s8.SwitchOrientierung()
		s8.SetzteGroesse(1)
		s8.SetzeKoordinaten(335,150)
		s9 = New()
		s9.SwitchOrientierung()
		s9.SetzteGroesse(1)
		s9.SetzeKoordinaten(335,250)
		s10 = New()
		s10.SwitchOrientierung()
		s10.SetzteGroesse(1)
		s10.SetzeKoordinaten(335,350)
	var yard []Schiff
	yard = append(yard,s1,s2,s3,s4,s5,s6,s7,s8,s9,s10)
	return yard
}

func SwitchYard (yard []Schiff) []Schiff{
	for i:=0;i<len(yard);i++{
		yard[i].SwitchOrientierung()
	}
	for i:=0;i<len(yard);i++{
		if yard[0].GibOrientierung() {
			switch i {
				case 0:
					yard[i].SetzeKoordinaten(50,50)
				case 1:
					yard[i].SetzeKoordinaten(145,50)
				case 2:
					yard[i].SetzeKoordinaten(145,245)
				case 3:
					yard[i].SetzeKoordinaten(240,50)
				case 4:
					yard[i].SetzeKoordinaten(240,200)
				case 5:
					yard[i].SetzeKoordinaten(240,350)
				case 6:
					yard[i].SetzeKoordinaten(335,50)
				case 7:
					yard[i].SetzeKoordinaten(335,150)
				case 8:
					yard[i].SetzeKoordinaten(335,250)
				case 9:
					yard[i].SetzeKoordinaten(335,350)
				}
		} else {
			switch i {
				case 0:
					yard[i].SetzeKoordinaten(50,50)
				case 1:
					yard[i].SetzeKoordinaten(50,150)
				case 2:
					yard[i].SetzeKoordinaten(50,250)
				case 3:
					yard[i].SetzeKoordinaten(300,50)
				case 4:
					yard[i].SetzeKoordinaten(300,150)
				case 5:
					yard[i].SetzeKoordinaten(300,250)
				case 6:
					yard[i].SetzeKoordinaten(50,350)
				case 7:
					yard[i].SetzeKoordinaten(150,350)
				case 8:
					yard[i].SetzeKoordinaten(250,350)
				case 9:
					yard[i].SetzeKoordinaten(350,350)
			}
		}
	}
	return yard
}
				

func DrawShipyard (yard []Schiff){
	for i:=0;i<len(yard);i++{
		yard[i].Draw()
	}
}

func Highlight (yard []Schiff) bool {
	var erg bool
	for i:=0;i<len(yard);i++{
		if yard[i].Gibhighlight() {
			erg = true
		}
	}
	return erg
} 

func SchiffGeben (yard []Schiff) ([]Schiff) {
	if Highlight(yard) {
		for i:=0;i<len(yard);i++ {
			if yard[i].Gibhighlight() {
				yard[i].SetzeDummy(true)
				yard[i].SwitchHighlighting()
			}
		}
	}
	return yard
}

func YardLeer (yard []Schiff) bool {
	for i:=0;i<len(yard);i++{
		if !yard[i].GibDummy() {
			return false
		}
	}
	return true
}
	
	
	
