package main

import (
	"fmt"
	"log"
	"os"

	"main/pkg/triangle"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Triangle Solver"
	app.Usage = "solve triangles"

	app.Commands = []cli.Command{
		{
			Name:    "solve",
			Aliases: []string{"s"},
			Usage:   "solve a triangle",
			Action: func(ctx *cli.Context) error {
				a := ctx.Float64("a")
				alpha := ctx.Float64("alpha")
				b := ctx.Float64("b")
				beta := ctx.Float64("beta")
				c := ctx.Float64("c")
				gamma := ctx.Float64("gamma")

				aTriangle, error := triangle.NewTriangle(a, alpha, b, beta, c, gamma)
				aTriangle.Solve()
				fmt.Println(aTriangle.Description())
				return error
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "a",
					Value: "0",
					Usage: "A side length",
				},
				cli.StringFlag{
					Name:  "alpha",
					Value: "0",
					Usage: "Alpha angle in radians",
				},
				cli.StringFlag{
					Name:  "b",
					Value: "0",
					Usage: "B side length",
				},
				cli.StringFlag{
					Name:  "beta",
					Value: "0",
					Usage: "Beta angle in radians",
				},
				cli.StringFlag{
					Name:  "c",
					Value: "0",
					Usage: "C side length",
				},
				cli.StringFlag{
					Name:  "gamma",
					Value: "0",
					Usage: "Gamma angle in radians",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
