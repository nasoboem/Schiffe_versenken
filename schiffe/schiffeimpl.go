package schiffe 



import (. "gfx2"
		)


type data struct {
	x,y uint16 //Koordinaten der linken oberen Ecke
	block uint16
	groesse uint16
	orientierung bool //orientierung true = waagerecht; false = senkrecht
	highlighting bool
	dummy bool
}

func New () *data {
	var s *data
	s = new(data)
	s.block = 50
	return s
}

func (s *data) SetzeKoordinaten (x,y uint16) {
	s.x = x
	s.y = y
}
func (s *data) SetzteGroesse (groesse uint16) {
	s.groesse = groesse 
	
}
func (s *data) GehoertPunktzuSchiff (x,y uint16) bool {
		var hoehe,breite uint16
	if s.orientierung {
		breite = s.block
		hoehe = s.block * s.groesse
	} else {
		breite = s.block * s.groesse
		hoehe = s.block
	} 
	return x>=s.x && x<=s.x+breite && y>=s.y && y<=s.y+hoehe
}

func (s *data) GibOrientierung() bool {
	return s.orientierung
}
func (s *data) GibGroesse() uint16{
	return s.groesse 
}

func (s *data) SwitchHighlighting() {
	if s.highlighting {
		s.highlighting = false
	}else{
		s.highlighting = true
	}
}

func (s *data) Gibhighlight() bool {
	return s.highlighting 
}

func (s *data) SwitchOrientierung() {
	if s.orientierung {
		s.orientierung = false
	}else{
		s.orientierung = true	
	}
}

func (s *data) SetzeDummy (dummy bool) {
	s.dummy = dummy
}

func (s *data) GibDummy () bool {
	return s.dummy
}

//koordinaten, treffer, groesse, orientierung
//wenn groesse = anzahl treffer, schiff versenkt
//orientierung mit rechtsklick wechselbar

//koordinate linker oberer punkt des schiffes 

func (s *data) Draw() {
	if !s.dummy {
		var hoehe,breite uint16
		if s.orientierung {
			breite = s.block
			hoehe = s.block * s.groesse
		} else {
			breite = s.block * s.groesse
			hoehe = s.block
		} 
		if s.highlighting{
			Stiftfarbe(255,0,35)
			Vollrechteck(s.x-5,s.y-5,breite+10,hoehe+10)
		}
		Stiftfarbe(113,113,113)
		Vollrechteck(s.x,s.y,breite,hoehe)
	}
}
