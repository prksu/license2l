package license

import "text/template"

// Data license data that will be rendered
type Data struct {
	Holder string
	Year   int
}

// License hold the license template
var License = make(map[string]*template.Template)

func init() {
	License["apache"] = template.Must(template.New("").Parse(apacheLicense))
	License["apache-2.0"] = template.Must(template.New("").Parse(apacheLicense))
	License["mit"] = template.Must(template.New("").Parse(mitLicense))
}
