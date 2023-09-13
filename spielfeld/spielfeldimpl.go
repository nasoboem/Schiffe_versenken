package spielfeld

import (
    "fmt"
    . "gfx2"
    "schiffe"
    . "zufallszahlen"
)

type data struct{
	groesse uint16
	besitzer string
	feld [][]uint16
	x_verschiebung uint16
	}

func New(groesse uint16) *data {
	var s *data
	s=new(data)
	s.groesse = groesse
	
	for y:=0;y<int(groesse);y++{
		var spalte []uint16
		for x:=0;x<int(groesse);x++{
				spalte = append (spalte, 0)
		}
		s.feld = append (s.feld,spalte)
	}
	return s
}

func (s *data) SetzeBesitzer (name string) {
	s.besitzer = name
}

func (s *data) SetzeX_Verschiebung (x uint16) {
	s.x_verschiebung = x
}

func (s *data) GibX_Verschiebung () uint16 {
	return s.x_verschiebung
}


func (s *data) SetzeFeld (x,y int, wert uint16) {
	if x<int(s.groesse) && y<int(s.groesse) {
		s.feld[x][y] = wert
	}
}

func (s *data) BombeEintragen (x,y int) {
	if x<int(s.groesse) && y<int(s.groesse) && s.feld[x][y]%2==0 {
		s.feld[x][y] = s.feld[x][y] + 1
	}
}

func (s *data) Bombardieren (x,y uint16) (kx,ky int, auffeld bool){
		kx = int(x-s.x_verschiebung)/75
		ky = int(y)/75
		if s.GehoertPunktzuSpielfeld(x,y){
			if s.feld[kx][ky]%2==0{
			//s.feld[kx][ky] = s.feld[kx][ky] + 1
			auffeld = true
			}
		} 
    return
}

func (s *data) String () string {
	var erg string
	for y:=0;y<int(s.groesse);y++{
		for x:=0;x<int(s.groesse);x++{
			erg = erg + fmt.Sprint(s.feld[x][y]) 
		}
		erg = erg + fmt.Sprintln("")
	}
	return erg
}

func (s *data) GibGroesse () uint16 {
	return s.groesse
}

func (s *data) Draw() {
	switch s.besitzer {
		case "":
				for y:=0;y<int(s.groesse);y++{
			for x:=0;x<int(s.groesse);x++{
				switch s.feld[x][y]{
					case 0:
						Stiftfarbe(124,185,232) // wasser hellblau
					case 1:
						Stiftfarbe(0,0,204)
					case 2:
						Stiftfarbe(124,185,232) // keine Schiffe
					case 3:
						Stiftfarbe(0,0,0)
					default:
						Stiftfarbe(255,255,255)
					
				}
				Vollrechteck(uint16(x)*75+s.x_verschiebung,uint16(y)*75,75,75)
				Stiftfarbe(255,255,255)
				Rechteck(uint16(x)*75+s.x_verschiebung,uint16(y)*75,75,75)
			}
		}

		default:
		for y:=0;y<int(s.groesse);y++{
			for x:=0;x<int(s.groesse);x++{
				switch s.feld[x][y]{
					case 0:
						Stiftfarbe(124,185,232) // wasser hellblau
					case 1:
						Stiftfarbe(0,0,204)
					case 2:
						Stiftfarbe(113,113,113) // schiff grau
					case 3:
						Stiftfarbe(0,0,0)
					default:
						Stiftfarbe(255,255,255)
					
				}
				Vollrechteck(uint16(x)*75+s.x_verschiebung,uint16(y)*75,75,75)
				Stiftfarbe(255,255,255)
				Rechteck(uint16(x)*75+s.x_verschiebung,uint16(y)*75,75,75)
			}
		}
	}
}


func (s *data) SchiffeSetzen (x,y uint16,c schiffe.Schiff) {
     var groesse uint16
     var orientierung bool //false=senkrecht true=waagrecht
     var kx,ky int
     kx = int(x-s.x_verschiebung)/75
     ky = int(y)/75
	orientierung = c.GibOrientierung()
    groesse = c.GibGroesse()
    if s.Schiffpasst(x,y,c) {
		if !orientierung {
			s.feld[kx][ky] = 2
			for i:=1;i<int(groesse);i++{
				s.feld[kx+i][ky] = 2
			}
		}else {
			s.feld[kx][ky] = 2
			for i:=1;i<int(groesse);i++{
				s.feld[kx][ky+i] = 2
			}
		}
	}
	s.Umranden()
}

func (s *data) RandomBomb () {
	nochmal:
	x:=uint16(Zufallszahl(int64(s.x_verschiebung),int64(s.x_verschiebung+75*s.groesse)))
	y:=uint16(Zufallszahl(0,int64(s.x_verschiebung+75*s.groesse)))
	if s.AufSpielfeld(x,y) {
		kx,ky,auffeld:=s.Bombardieren(x,y)
		if auffeld {
			s.BombeEintragen(kx,ky)
		} else {
			goto nochmal
		}
	}
}


 
func (s *data) ProzentBomb (prozent int) {
	var zz int
	zz = int (Zufallszahl(1,100))
	if zz <= prozent {
		nochmal2:
		x:=uint16(Zufallszahl(int64(s.x_verschiebung),int64(s.x_verschiebung+75*s.groesse)))
		y:=uint16(Zufallszahl(0,int64(s.x_verschiebung+75*s.groesse)))
		if s.AufSpielfeld(x,y) {
			kx,ky,auffeld:=s.Bombardieren(x,y)
			if auffeld  && s.gibfeld(kx,ky) == 2{
				s.BombeEintragen(kx,ky)
			} else {
				goto nochmal2
			}
		}
	} else {
		nochmal0:
		x:=uint16(Zufallszahl(int64(s.x_verschiebung),int64(s.x_verschiebung+75*s.groesse)))
		y:=uint16(Zufallszahl(0,int64(s.x_verschiebung+75*s.groesse)))
		if s.AufSpielfeld(x,y) {
			kx,ky,auffeld:=s.Bombardieren(x,y)
			if auffeld  && s.gibfeld(kx,ky) == 0{
				s.BombeEintragen(kx,ky)
			} else {
				goto nochmal0
			}
		}
	}
}
		

func (s *data) Schiffpasst (x,y uint16, c schiffe.Schiff) bool {
	 var erg bool
	 erg = true
	 var groesse uint16
     var orientierung bool
     orientierung = c.GibOrientierung()
     groesse = c.GibGroesse()
     var kx,ky int
     kx = int(x-s.x_verschiebung)/75
     ky = int(y)/75
     if !orientierung {
		 if kx + int(groesse) > int(s.groesse) {
			 erg = false
		 } else {
			for i:=kx;i<kx+int(groesse);i++{
				if s.feld[i][ky] > 0 {
					erg = false
				}
			}
		}
	 } else {
		 if ky + int(groesse) > int(s.groesse) {
			 erg = false
		 }else {
			for i:=ky;i<ky+int(groesse);i++{
				if s.feld[kx][i] > 0 {
					erg = false
				}
			}
		}
	 }
	 return erg
 }

func (s *data) StringToFeld (sfeld string) {
	var zeile []uint16
	var feld [][]uint16
	for _,w:=range(sfeld) {
		if w =='\n' {
			feld = append(feld,zeile)
			var nzeile []uint16
			zeile = nzeile
		} else {
			switch w {
				case '0':
					zeile = append(zeile,0)
				case '1':
					zeile = append(zeile,1)
				case '2':
					zeile = append(zeile,2)
				case '3':
					zeile = append(zeile,3)
				default:
					zeile = append(zeile,5)
			}
		}
	}
	
	s.feld = transpose(feld)
	s.besitzer = ""
	s.x_verschiebung = 0
	s.groesse = uint16(len(s.feld))
}

func transpose(slice [][]uint16) [][]uint16 {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]uint16, xl)
    for i := range result {
        result[i] = make([]uint16, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = slice[j][i]
        }
    }
    return result
}
		 


func (s *data) Umranden() {
	for i:=0;i<len(s.feld);i++{
		for j:=0;j<len(s.feld);j++{
			if s.feld[i][j] == 2 {
				for x:=i-1;x<i+2;x++{  					//feld[i-1][j-1] - feld[i+1][j+1]
					for y:=j-1;y<j+2;y++{
						if x>=0&&y>=0&&x<len(s.feld)&&y<len(s.feld) && s.feld[x][y]==0 {
							s.feld[x][y] = 5
						}
					}
				}
			}
		}
	}
}

func (s* data) Removefive(){
	for i:=0;i<len(s.feld);i++{
		for j:=0;j<len(s.feld);j++{
			if s.feld[i][j] == 5{
				s.feld[i][j] = 0
			}
		
	    }
	}
}

func (s *data) GehoertPunktzuSpielfeld (x,y uint16) bool{
	var kx,ky int
	kx = int(x-s.x_verschiebung)/75
	ky = int(y)/75
	return kx<int(s.groesse) && ky<int(s.groesse)
}

func (s *data) AufSpielfeld (x,y uint16) bool {
	return x>=s.x_verschiebung && x<=s.x_verschiebung + s.groesse*75
}

func (s *data) Feld (x,y uint16) (kx,ky int) {
	kx = int(x-s.x_verschiebung)/75
    ky = int(y)/75
    return
}

func (s *data) gibfeld (kx,ky int) uint16 {
	return s.feld[kx][ky]
}

func (s *data) GameOver() bool {  //felder nach zweien überprüfen
	for i:=0;i<len(s.feld);i++{
		for j:=0;j<len(s.feld);j++{
			if s.feld[i][j] == 2{
				return false
			}
		
	    }
	}
	return true
}
