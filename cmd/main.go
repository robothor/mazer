package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/peterbourgon/ff"
	"github.com/robothor/mazer"
	"github.com/robothor/mazer/algorithms"
)

func main() {
	fmt.Println("Starting Mazer")
	fs := flag.NewFlagSet("mazer", flag.ExitOnError)
	var (
		rows = fs.Int("rows", 5, "Number of rows")
		cols = fs.Int("cols", 5, "Number of columns")
		//debug = fs.Bool("debug", false, "Enable debug output")
		_ = fs.String("config", "", "config file (optional)")
	)

	ff.Parse(fs, os.Args[1:],
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
		ff.WithEnvVarPrefix("MAZER"),
	)

	rng := rand.New(rand.NewSource(time.Now().Unix()))
	g := mazer.NewGrid(*rows, *cols)
	alg := &algorithms.BinaryTreeMazer{Rng: rng}
	alg.On(g)
	fmt.Println("Binary Tree")
	g.PrintASCII()
}
