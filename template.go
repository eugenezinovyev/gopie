package gopie

import (
	"path"
	"text/template"

	"github.com/eugenezinoviev/gopie/assets"
)

const rootTemplateFile = "assets/template.root.svg"

var childTemplates = []string{
	"assets/template.background.svg",
	"assets/template.circle.svg",
	"assets/template.circle.mask.svg",
	"assets/template.slice.svg",
	"assets/template.slice.mask.svg",
	"assets/template.slice.label.svg",
	"assets/template.font.svg",
}

func createSvgTemplate() (tpl *template.Template, err error) {
	tpl, err = createRootTemplate()

	for _, filename := range childTemplates {
		if err = appendTemplate(tpl, filename); err != nil {
			return
		}
	}
	return
}

func createRootTemplate() (tpl *template.Template, err error) {
	rootTemplate := assets.GetFileContent(rootTemplateFile)
	funcs := template.FuncMap{
		"derefCircle": func(c *circle) circle { return *c },
	}
	tpl, err = template.New("root").Funcs(funcs).Parse(rootTemplate)
	return
}

func appendTemplate(t *template.Template, filename string) error {
	_, p := path.Split(filename)
	_, err := t.New(p).Parse(assets.GetFileContent(filename))
	return err
}
