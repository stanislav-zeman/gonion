package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/stanislav-zeman/gonion/internal/config"
	"github.com/stanislav-zeman/gonion/internal/initor"
	processor "github.com/stanislav-zeman/gonion/internal/procesor"
	"github.com/stanislav-zeman/gonion/internal/templator"
	"github.com/stanislav-zeman/gonion/internal/writer"
	yaml "gopkg.in/yaml.v3"
)

var (
	configPath      = flag.String("config", "gonion.yaml", "project structure configuration")
	outputDirectory = flag.String("out", ".", "project structure output directory")
)

func main() {
	flag.Parse()

	log.Println("running gonion...")
	err := runGonion()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func runGonion() error {
	f, err := os.ReadFile(*configPath)
	if err != nil {
		return fmt.Errorf("failed reading config file: %w", err)
	}

	var conf config.Config
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		return fmt.Errorf("failed unmarshalling config file: %w", err)
	}

	i := initor.New(conf, *outputDirectory)

	err = i.Run()
	if err != nil {
		return fmt.Errorf("failed running initor: %w", err)
	}

	t, err := templator.New("assets")
	if err != nil {
		return fmt.Errorf("failed creating templator: %w", err)
	}

	w := writer.NewWriter(*outputDirectory)
	p := processor.New(conf, t, w)

	err = p.Run()
	if err != nil {
		return fmt.Errorf("failed running processor: %w", err)
	}

	return nil
}
