package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/gregoryv/cmdline"
	"gopkg.in/yaml.v2"
)

func main() {
	log.SetFlags(0)

	var (
		cli         = cmdline.NewBasicParser()
		filename    = cli.Option("-cv, --cv-input").String("")
		companyFile = cli.Option("-co, --company-file").String("")
		maxSkills   = cli.Option("-ms, --max-skills").Uint(1000)
		sortSkills  = cli.Option("-ss, --sort-skills").Enum(
			"by-experience", "by-name", "by-experience",
		)
		maxProjects  = cli.Option("-mp, --max-projects").Uint(1000)
		fullProjects = cli.Option("-fp, --full-projects").Uint(3)
		template     = cli.Option("-t, --template").Enum("one-page", "one-page", "full")
		saveas       = cli.Option("-s, --save-as").String("cv.html")
		showVersion  = cli.Flag("-v, --version")
	)
	cli.Parse()

	if showVersion {
		fmt.Println(version())
		os.Exit(0)
	}

	switch {
	case filename == "":
		log.Fatal("missing --cv-input")
	}

	// load curriculum vitae
	var in CV
	loadYaml(filename, &in)
	var co Company
	if companyFile != "" {
		loadYaml(companyFile, &co)
	}

	// prepare model, depending on what output you want
	switch sortSkills {
	case "by-name":
		sort.Sort(TechSkillByName(in.TechnicalSkills))
	case "by-experience":
		sort.Sort(sort.Reverse(TechSkillByE(in.TechnicalSkills)))
	}

	switch template {
	case "one-page":
		// limit skills
		if max := int(maxSkills); len(in.TechnicalSkills) > max {
			in.TechnicalSkills = in.TechnicalSkills[:max]
		}
		// short descriptions for all projects
		for i, _ := range in.Experience {
			in.Experience[i].showShort = true
		}

		// trim projects to fit more on first page
		if max := int(fullProjects); len(in.Experience) > max {
			for i := max; i < len(in.Experience); i++ {
				in.Experience[i].oneLiner = true
			}
		}
		// hide projects after maxProjects
		for i := 0; i < len(in.Experience); i++ {
			if i < int(maxProjects) {
				continue
			}
			in.Experience[i].hide = true
		}

	case "full":
		for i, _ := range in.Experience {
			in.Experience[i].showMore = true
			in.Experience[i].showShort = true
		}

	default:
		log.Fatal("unknown template")
	}

	// Create and save page
	page := NewCVPage(&co, &in)
	switch saveas {
	case "":
		page.WriteTo(os.Stdout)
	default:
		page.SaveAs(saveas)
	}
}

// loadYaml reads the given yaml file and unmarshals into object.
// Fatal on errors.
func loadYaml(filename string, into interface{}) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, into); err != nil {
		log.Fatal(err)
	}
}

func version() string {
	from := bytes.Index(changelog, []byte("## ["))
	to := bytes.Index(changelog[from:], []byte("]"))
	return string(changelog[from+4 : from+to])
}

//go:embed changelog.md
var changelog []byte
