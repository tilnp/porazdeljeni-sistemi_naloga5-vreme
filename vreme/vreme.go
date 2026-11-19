// Package vreme nudi vremenske postaje in funkcije za belezenje
// temperature ter izpis statistik
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo in njene meritve
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo med postaje doda novo postajo z danim imenom in lokacijo pod kljucem oznaka
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// DodajMeritev doda meritev temperature postaji z oznako oznaka
// meritev je zavrnjena, ce postaja ne obstaja ali ce lezi izven meje [minTemp, maxTemp]
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	key, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Printf("Postaja z oznako %s ne obstaja.\n", oznaka)
		return
	}
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Neveljavna temperatura: %.1f °C\n", temperatura)
		return
	}
	key.meritve = append(key.meritve, temperatura)
	postaje[oznaka] = key
}

// statistika vrne najnizjo, najvisjo in povprecno temperaturo postaje z oznako oznaka
// ok je false ce postaja ne obstaja ali ima manj kot stMeritev meritev
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	key, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Printf("Postaja z oznako %s ne obstaja.\n", oznaka)
		return
	}
	if len(key.meritve) < stMeritev {
		ok = false
		return
	}
	ok = true
	najnizja = key.meritve[0]
	najvisja = key.meritve[0]
	var sum float64
	for _, m := range key.meritve {
		sum += m
		if m < najnizja {
			najnizja = m
		}
		if m > najvisja {
			najvisja = m
		}
	}
	povprecna = sum / float64(len(key.meritve))
	return
}

// IzpisMeritev izpise meritve vseh postaj.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for key, postaja := range postaje {
		fmt.Printf("%s - %s (%s): ", key, postaja.ime, postaja.kraj)
		fmt.Println(postaja.meritve)
	}
}

// IzpisPorocila izpise porocilo o najnizji, najvisji in povprecni
// temperaturi vsake postaje, ki ima vsaj stMeritev meritev
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for key, postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, key, stMeritev)
		if ok {
			var a string
			switch {
			case povprecna >= 25:
				a = "Vroče"
			case povprecna >= 15:
				a = "Toplo"
			case povprecna >= 5:
				a = "Hladno"
			default:
				a = "Mrzlo"
			}
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", postaja.ime, najnizja, najvisja, povprecna, a)
		} else {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
		}
	}
}
