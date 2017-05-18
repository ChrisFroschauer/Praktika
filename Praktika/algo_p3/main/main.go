package main

/**
Christian Froschauer

Algorithmen und Datenstrukturen Praktikum 3
Karazuba Algorithmus
 */
import (
	"math"
	"strconv"
)

//import "fmt"
func main(){
	u := "100111" //100111 = 39
	v := "-110" //-110 = -6

	//println(subtractBinaryString(u,v)) //teste subtractBinaryString
	//println(subtractBinaryString(v,u))
	println(u + " * " + v + " = ")
	println(multiplyKarazuba(u,v,2)) // Karazuba mit Strings
	println(strconv.Itoa(baseToDec(u,2)) + " * " + strconv.Itoa(baseToDec(v, 2)) + " = ")
	println(baseToDec(u,2)*baseToDec(v,2)) // Multiplikation Integer
}

//derzeit gehts nur für binär wegen subtract
func multiplyKarazuba(u string, v string, base int) int {

	// temporär nur für base=2 gelöst, wegen subtractBinaryString
	if(base !=2){
		println("Derzeit nur Basis 2 erlaubt (subtraktionsmethode gehört noch verallgemeinert)")
		return math.MinInt32
	}

	// Festhalten des Vorzeichens
	sign := 1
	if u[0] == '-'{
		sign = sign*(-1)
		u = u[1:]
	}
	if v[0] == '-'{
		sign = sign*(-1)
		v = v[1:]
	}

	//Ausstieg aus der rekursion
	if ( (len(u)==2 && u[0]=='-') || len(u)==1){
		uInt := int(u[len(u)-1]) - 48
		vInt := int(v[len(v)-1]) - 48
		return sign*uInt*vInt
	}

	//zweierpotenz länge, sonst geht das mittig splitten nicht
	if !(len(u)==len(v) && isBasepotenz(len(u), base)){
		for !isBasepotenz(len(u), base) {
			u = "0" + u
		}

		for !isBasepotenz(len(v), base){
			v = "0" + v
		}
		// gleiche länge:
		for len(v)<len(u){
			v = "0" + v
		}
		for len(u)<len(v){
			u = "0" + u
		}
	}



	// uL uR vL vR init
	uL := u[:len(u)/2]
	uR := u[len(u)/2:]
	vL := v[:len(v)/2]
	vR := v[len(v)/2:]

	//println(uL + uR + " " + vL + vR)

	// Karazuba formel:
	return sign * ( (1<<uint(len(u)) + 1<<uint((len(u)/2)))*multiplyKarazuba(uL, vL, base) +
		(1<<uint(len(u)/2))*multiplyKarazuba(subtractBinaryString(uL, uR), subtractBinaryString(vR,vL), base) +
		(1<<uint((len(u)/2))+1)*multiplyKarazuba(uR,vR, base))

}

//derzeit gehts nur für binär... wegen subtract
func subtractBinaryString(a string, b string) string{
	ergeb := ""
	// a und b garantiert gleich lang

	// subtraktions routine
	merk := 0
	for i:=1; i<=len(a); i++{
		aIntAtI := int(a[len(a)-i])-48
		bIntAtI := int(b[len(b)-i])-48
		zwischen := aIntAtI - bIntAtI - merk
		if zwischen <0{
			merk = 1;
			if zwischen == -1{
				ergeb = "1" + ergeb
			}else if zwischen == -2{
				ergeb = "0" + ergeb
			}
		}else{
			merk = 0;
			if zwischen == 0 {
				ergeb = "0" + ergeb
			}else if zwischen == 1 {
				ergeb = "1" + ergeb
			}
		}
	}

	// finish up negativ ergebnis:
	if merk == 1{ // übertrag am ende noch
		//println(ergeb)
		//ergeb ist negativ, das heißt alle einsen vorn vom Ergebnis weg

		for len(ergeb)>0 && ergeb[0]=='1'{
			ergeb = ergeb[1:]
		}



		//println(ergeb)
		// jetzt das komplement vom ergebnis
		ergComp := "";
		for i:=0; i<len(ergeb); i++{
			if ergeb[i] == '1'{
				ergComp = ergComp+"0"
			}else{
				ergComp = ergComp+"1"
			}
		}
		//println(ergComp)
		// jetzt ergComp + 1
		done := false
		for i:=len(ergComp)-1; i>=0; i--{
			if ergComp[i] == '0' && done == false{
				ergComp = ergComp[:i] + "1" + ergComp[i+1:] // die 0 bei i mit 1 ersetzen
				done = true
			}else if ergComp[i] == '1' && done == false{
				ergComp = ergComp[:i] + "0" + ergComp[i+1:] // 1 mit 0 ersetzen
			}
		}
		ergeb = ergComp
		//println(ergeb)
		// wenn durchgelaufen ohne 0 gefunden, dann vorne ne 1 anhängen
		if (done == false){
			ergeb = "1" + ergComp
		}

		// so jetzt auf länge auffüllen mit 0
		for len(ergeb)<len(a){
			ergeb = "0" + ergeb
		}

		// vorne minus "-" anghängen weil karazuba-funktion das interpretieren kann
		ergeb = "-" + ergeb

	}
	return ergeb
}

func isBasepotenz(n int, base int) bool{
	for n!=1 {
		if n%base == 0 {
			n = n/base
		}else{
			return false
		}
	}
	return true
}

// übersetzung eines alten Java Codes von mir, in Go erstaunlich schlank
func baseToDec(in string, base int) int {

	// Vorzeichen festhalten
	sign := 1
	if in[0] == '-'{
		sign = -1
		in = in[1:]
	}

	// Umwandlungsroutine
	sum := 0;
	times := 1;
	for i := 0; i < len(in); i++ {
		times = 1;
		for j := 1; j < len(in)-i; j++ {
			times *= base;
		}
		if(in[i]>='a'&& in[i]<='z'){
			sum += (int(in[i]-'a')+10)*times;
		}else{
			sum += (int(in[i])-48)*times;
		}
	}
	return sign*sum;
}
/* Noch großteil java code
decToBase(numb int, base int) string{
	numbStr := "";

	// ab hier noch java:
	while (numb>0){
		if(numb%base>=10){
			numbStr = (char)((numb%base)-10+'A') + numbStr;
		}else {
			numbStr = (numb % base) + numbStr;
		}
		numb = numb/base;
	}
	return numbStr;

}*/
