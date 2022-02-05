package main

import (
	"log"
	"os"
	"sort"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/web"
	"gopkg.in/yaml.v2"
)

func main() {
	var (
		cli          = cmdline.NewBasicParser()
		filename     = cli.Option("-cv, --cv-input").String("")
		companyFile  = cli.Option("-co, --co-file").String("")
		maxSkills    = cli.Option("-ms, --max-skills").Uint(1000)
		maxProjects  = cli.Option("-mp, --max-projects").Uint(1000)
		fullProjects = cli.Option("-fp, --full-projects").Uint(3)
		template     = cli.Option("-t, --template").Enum("one-page", "one-page", "full")
		saveas       = cli.Option("-s, --save-as").String("cv.html")
	)
	cli.Parse()

	// load curriculum vitae
	var in CV
	loadYaml(filename, &in)
	var co Company
	loadYaml(companyFile, &co)
	// prepare model, depending on what output you want
	sort.Sort(sort.Reverse(TechSkillByE(in.TechnicalSkills)))

	var page *web.Page
	switch template {
	case "one-page":
		// limit skills
		if max := int(maxSkills); len(in.TechnicalSkills) > max {
			in.TechnicalSkills = in.TechnicalSkills[:max]
		}
		// short descriptions for all projects
		for i, _ := range in.Projects {
			in.Projects[i].showShort = true
		}

		// trim projects to fit more on first page
		if max := int(fullProjects); len(in.Projects) > max {
			for i := max; i < len(in.Projects); i++ {
				in.Projects[i].oneLiner = true
			}
		}
		// hide projects after maxProjects
		for i := 0; i < len(in.Projects); i++ {
			if i < int(maxProjects) {
				continue
			}
			in.Projects[i].hide = true
		}

	case "full":
		for i, _ := range in.Projects {
			in.Projects[i].showShort = true
			in.Projects[i].showMore = true
		}

	default:
		log.Fatal("unknown template")
	}

	// Create and save page
	page = NewCVPage(&co, &in)
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
