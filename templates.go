package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/theme"
)

func NewTemplate(co *Company, in *CV) *Page {
	content := Wrap(
		Header(
			func() *Element {
				if co.hide {
					return Wrap()
				}
				return Div(Class("company"),
					Img(Class("logo"),
						Src(co.Logo),
					),
					Br(),
					co.Phone,
				)
			}(),
			Div(Class("picname"),
				Img(Src(in.Person.Image)),
				H1(in.Person.Name),
			),
		),

		Div(Class("left"), Class("column"),

			H2("Technical skills"),
			func() *Element {
				ul := Ul()
				for _, skill := range in.TechnicalSkills {
					ul.With(Li(skill.Item,
						func() *Element {
							// experience bar
							w := skill.E * 10.0
							widthAttr := Attr("style", fmt.Sprintf("width: %vpx", w))

							return Div(Class("bar"),
								Div(Class("exp"), widthAttr, "&nbsp;"),
							)
						}(),
					))
				}
				return ul
			}(),

			H2("Languages"),
			func() *Element {
				ul := Ul()
				for _, lang := range in.Languages {
					ul.With(Li(lang.Item))
				}
				return ul
			}(),

			H2("Education"),
			func() *Element {
				ul := Ul()
				for _, edu := range in.Educations {
					ul.With(Li(edu.Grade, " in ", edu.Subject), periodSpan(edu.Period))
				}
				return ul
			}(),
		),

		Div(Class("right"), Class("column"),
			H2("Personal"),
			P(in.Person.Description),

			H2("Experience",
				Span(" (", experienceYears(in.Experience), " years)"),
			),
			func() *Element {

				s := Section()

				var (
					oneLiners      = make([]Project, 0)
					fullExperience = make([]Project, 0)
				)
				for _, p := range in.Experience {
					if p.oneLiner {
						oneLiners = append(oneLiners, p)
						continue
					}
					fullExperience = append(fullExperience, p)
				}
				for _, p := range fullExperience {
					if p.hide {
						continue
					}
					h := H3(p.Title,
						Span(Class("customer"), p.Customer),
						periodSpan(p.Period),
					)

					content := P()
					if p.showShort {
						content.With(p.Short)
					}
					if p.showMore {
						content.With(p.More)
					}
					roles := Wrap()
					if len(p.Roles) > 0 {
						roles = Span(Class("roles"),
							strings.Join(p.Roles, ", "), ": ",
						)
					}

					tags := Wrap()
					if len(p.Tags) > 0 {
						tags = Span(Class("tags"),
							roles, strings.Join(p.Tags, ", "),
						)
					}
					s.With(
						h,
						content,
						tags,
					)
				}

				if len(oneLiners) == 0 {
					return s
				}

				rest := Section(Class("rest"))
				for _, p := range oneLiners {
					if p.hide {
						continue
					}
					h := H3(Class("oneliner"), p.Title,
						Span(Class("customer"), p.Customer),
						periodSpan(p.Period),
					)

					rest.With(h)
				}
				return Wrap(s, rest)

			}(),
		),
	)
	return NewPage(
		Html(
			Head(
				Lang("en-US"),
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1"),
				),
				Style(onePage()),
			),
			Body(
				Nav(
					A(Href("short.html"), "Short"),
					" | ",
					A(Href("full.html"), "Full"),
				),
				content,
				Br(Attr("clear", "all")),
				Div(Class("footer"),
					"Updated: ", time.Now().Format("2006-01-02"),
				),
			),
		),
	)
}

func experienceYears(projects []Project) int {
	// calculate number of hears
	var from int
	var to int
	for _, p := range projects {
		if from == 0 || p.Period.FromYear < from {
			from = p.Period.FromYear
		}
		if p.Period.ToYear > to {
			to = p.Period.ToYear
		}
	}
	return to - from
}

func periodSpan(p Period) *Element {
	var content interface{} = Wrap(
		p.FromYear,
		" - ",
		p.ToYear,
	)
	if p.FromYear == p.ToYear {
		content = p.FromYear
	}
	return Span(Class("period"), content)
}

func onePage() *CSS {
	css := theme.GoldenSpace()
	css.Style("nav",

		"text-align: right",
		"margin-bottom: 2em",
		"padding-bottom: 1em",
		"margin-left: -1.63em",
		"margin-right: -1.63em",
	)
	css.Style("nav a",
		"text-decoration: none",
	)
	css.Style("body",
		"width: 18.5cm",
	)
	blue := "#00a4e4"
	//
	css.Style("*",
		"font-family: Arial, Sans-serif",
	)

	css.Style("h1,h2, h2>*",
		"color: "+blue,
		"font-family: Helveticat, Serif",
	)
	css.Style("header",
		"height: 140px",
		"border-bottom: 1px solid "+blue,
		"text-align: right",
	)
	css.Style(".footer",
		"font-size: 0.8em",
		"text-align: center",
	)

	css.Style(".company",
		"float: left",
		"text-align: center",
	)

	css.Style(".picname",
		"text-align: center",
		"float: right",
	)
	css.Style(".picname img",
		"width: 100px",
	)

	css.Style("header h1",
		"font-size: 1em",
		"margin-top: 0",
		"text-align: right",
		"color: black",
	)

	css.Style("h2",
		"font-size: 1.2em",
	)

	css.Style(".left",
		"float: left",
		"margin-right: 1em",
		"width: 190px",
	)

	css.Style(".left ul",
		"font-size: 0.9em",
		"padding-left: 16px",
	)

	css.Style(".right",
		"margin-left: 230px",
	)

	css.Style(".tags",
		"display: block",
		"font-size: 0.8em",
		"margin-top: -10px",
	)
	css.Style("h2>span",
		"font-weight: normal",
	)
	css.Style(".period",
		"font-weight: normal",
		"font-size: 0.7em",
		"font-style: italic",
	)
	css.Style("h3 .customer",
		"font-weight: normal",
		"font-size: 0.8em",
		"margin-left: 1em",
	)
	css.Style(".rest",
		"margin-top: 2em",
	)
	css.Style(".rest h3",
		"margin-top: 3px",
		"margin-bottom: 3px",
		"font-size: 1em",
	)
	css.Style("h3 .period",
		"margin-top: 5px",
		"margin-left: 20px",
		"float: right",
	)

	css.Style(".bar",
		"margin-top: 3px",
		"border: 1px solid #727272",
		"width: 52px",
		"float: right",
		"height: 10px",
	)

	css.Style(".exp",
		"background-color: "+blue,
		"height: 8px",
		"margin-top: 1px",
		"margin-left: 1px",
	)

	p := css.Media("print")
	p.Style("nav", "display: none")
	return css
}
