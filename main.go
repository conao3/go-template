package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/goccy/go-yaml"
	flags "github.com/jessevdk/go-flags"
)

func main_() (int, error) {
	var opts struct {
		Input string `short:"i" long:"input" description:"Input file" default:"-"`
		Param string `short:"p" long:"param" description:"Parameter file"`
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		return 1, err
	}

	if len(args) > 0 {
		return 1, fmt.Errorf("unexpected argument: %q", args[0])
	}

	var r io.Reader
	switch opts.Input {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(opts.Input)
		if err != nil {
			return 1, err
		}
		defer f.Close()
		r = f
	}

	t, err := io.ReadAll(r)
	if err != nil {
		return 1, err
	}

	tmpl, err := template.New("base").Funcs(sprig.FuncMap()).Parse(string(t))
	if err != nil {
		return 1, err
	}

	var param any
	if opts.Param != "" {
		f, err := os.Open(opts.Param)
		if err != nil {
			return 1, err
		}
		defer f.Close()

		t, err := io.ReadAll(f)
		if err != nil {
			return 1, err
		}
		err = yaml.Unmarshal([]byte(t), &param)
		if err != nil {
			return 1, err
		}
	}

	err = tmpl.Execute(os.Stdout, param)
	if err != nil {
		return 1, err
	}

	return 0, nil
}

func main() {
	ret, err := main_()
	if err != nil {
		log.Println(err)
		os.Exit(ret)
	}
}
