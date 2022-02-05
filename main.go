package main

import (
	"flag"
	"log"
	"os"
	"sort"

	"github.com/gregoryv/web"
	"gopkg.in/yaml.v2"
)

func main() {
	var (
		filename    string
		maxSkills   uint = 1000
		maxProjects uint = 1000
		template         = "one-page"
		saveas           = "cv.html"
	)
	flag.StringVar(&filename, "cv", "", "CV in yaml format")
	flag.UintVar(&maxSkills, "max-skills", maxSkills, "Number of skills to show")
	flag.UintVar(&maxProjects, "max-projects", maxProjects, "Number of projects to show")
	flag.StringVar(&template, "template", template, "Output template, one-page or full")
	flag.StringVar(&saveas, "save-as", saveas, "Html file to save")
	flag.Parse()

	// load curriculum vitae
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var in CV
	if err := yaml.Unmarshal(data, &in); err != nil {
		log.Fatal(err)
	}

	co := Company{
		Logo:  "https://preferit.se/Smith/uploads/2018/01/preferit-logo-1.png",
		Phone: "+46 (0) 76 122 93 40",
	}
	// prepare model, depending on what output you want
	sort.Sort(sort.Reverse(TechSkillByE(in.TechnicalSkills)))

	var page *web.Page
	switch template {
	case "one-page":
		// limit skills
		if max := int(maxSkills); len(in.TechnicalSkills) > max {
			in.TechnicalSkills = in.TechnicalSkills[:max]
		}

		for i, _ := range in.Projects {
			in.Projects[i].showShort = true
			//in.Projects[i].showMore = true
		}

		// trim projects to fit more on first page
		if max := 3; len(in.Projects) > max {
			for i := max; i < len(in.Projects); i++ {
				in.Projects[i].oneLiner = true
			}
		}

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

	page = NewOnePage(&co, &in)
	switch saveas {
	case "":
		page.WriteTo(os.Stdout)
	default:
		page.SaveAs(saveas)
	}
}
