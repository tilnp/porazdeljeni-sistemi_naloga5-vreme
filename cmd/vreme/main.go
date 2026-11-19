package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tilnp/porazdeljeni-sistemi_naloga5-vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "weather station measurements and report",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "lowest allowed temperature (°C)",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "highest allowed temperature (°C)",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "minimum number of measurements for statistics",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			run(cmd.Float64("minTemp"), cmd.Float64("maxTemp"), cmd.Int("stMeritev"))
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]vreme.Postaja)

	podatki := []struct {
		oznaka  string
		ime     string
		kraj    string
		meritve []float64
	}{
		{"LJ01", "FRI", "Ljubljana", []float64{32.3, 32.8, 33.8, 34.7, 35.2, 36.7, 38.5}},
		{"LJ02", "Ljubljana Center", "Ljubljana", []float64{-1.2, -0.1, 0.2, 1.3, 2.5, 3.2, 4.0}},
		{"ŠL01", "Škofjeloški Grad", "Škofja Loka", []float64{14.8, 15.4, 16.9, 17.7, 18.5, 20.8, 21.8, 23.2}},
		{"KR01", "GFP", "Kranj", []float64{21.2, 21.9}},
		{"MD01", "Spar Medvode", "Medvode", []float64{1.0, 1.9, 2.6, 3.2, 3.9, 4.7, 6.1, 8.0, 9.9, 10.8}},
	}
	for _, p := range podatki {
		vreme.DodajPostajo(postaje, p.oznaka, p.ime, p.kraj)
		for _, m := range p.meritve {
			vreme.DodajMeritev(postaje, p.oznaka, m, minTemp, maxTemp)
		}
	}

	vreme.IzpisMeritev(postaje)
	fmt.Println()

	vreme.DodajMeritev(postaje, "MB01", 11.2, minTemp, maxTemp) // ne obstaja
	fmt.Println()
	vreme.DodajMeritev(postaje, "ŠL01", 51.3, minTemp, maxTemp) // previsoka temperatura
	fmt.Println()
	vreme.DodajMeritev(postaje, "LJ01", 39.2, minTemp, maxTemp) // ok
	fmt.Println()

	vreme.IzpisMeritev(postaje)
	fmt.Println()
	vreme.IzpisPorocila(postaje, stMeritev)
}
