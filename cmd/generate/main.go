package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/evanw/esbuild/pkg/api"

	"github.com/lytric-health/js-polyfill-server/version"
)

var (
	funcs = template.FuncMap{
		"var": func(name string) string {
			name = strings.ReplaceAll(name, ".", "_")
			name = strings.ReplaceAll(name, "@", "α")
			name = strings.ReplaceAll(name, "~", "__")
			name = strings.ReplaceAll(name, "-", "_")
			return strings.Title(name)
		},
	}

	codeTmpl = template.Must(template.New("code").Funcs(funcs).Parse(`package polyfill
{{ range . }}
const {{var .Name}}_Code = {{printf "%q" .Code}}
{{ end }}
`))
	polyTmpl = template.Must(template.New("poly").Funcs(funcs).Parse(`package polyfill
import "github.com/lytic-health/js-polyfill-server/version"
import "github.com/lytic-health/js-polyfill-server/types"
{{ range . }}
var {{var .Name}} = types.NewPolyfill(
	 {{var .Name}}_Code,
	 []*types.Polyfill{ {{range .Conf.Deps}} {{var .}}, {{end}} },
	 types.SupportMatrix{ {{range $k, $v := .Conf.Versions}}
		{{$k}}Needed: {{$v}},{{end}}
	},
)
{{ end }}

var polyfillsByName = map[string]*types.Polyfill{
{{ range . }}
	"{{.Name}}": {{var .Name}},{{end}}
}
`))
)

type polyfill struct {
	Name string
	Code string
	Conf polyfillConfig
}

type polyfillConfig struct {
	Deps     []string
	Versions map[string]string
}

func intervalString(fn string, version string) string {
	switch version {
	case "":
		return "version.None"
	case "*":
		return "version.All"
	}

	return fmt.Sprintf(`%s(version.MustParse("%s"))`, fn, version)
}

func rangeString(s string) string {
	switch s {
	case "":
		return "version.None"
	case "*":
		return "version.All"
	}

	parts := strings.Split(s, "-")
	if len(parts) == 2 {
		return fmt.Sprintf("version.Range(%s).AND(%s)",
			intervalString("version.MakeGTE", strings.TrimSpace(parts[0])),
			intervalString("version.MakeLTE", strings.TrimSpace(parts[1])),
		)
	}

	v, err := version.Parse(s)
	if err != nil {
		// Assume it's a semver-ish range
		return fmt.Sprintf(`version.MustParseRange("%s")`, s)
	}
	if v.Patch != 0 {
		return fmt.Sprintf(`version.MakeEqualsFull(version.MustParse("%s"))`, s)
	}
	if v.Minor != 0 {
		return fmt.Sprintf(`version.MakeEqualsMajorMinor(version.MustParse("%s"))`, s)
	}
	return fmt.Sprintf(`version.MakeEqualsMajor(version.MustParse("%s"))`, s)
}

func fromTOML(in polyfillTOMLConfig) polyfillConfig {
	r := rangeString
	return polyfillConfig{
		Deps: in.Deps,
		Versions: map[string]string{
			"Android":       r(in.Browsers.Android),
			"BB":            r(in.Browsers.BB),
			"Chrome":        r(in.Browsers.Chrome),
			"Edge":          r(in.Browsers.Edge),
			"EdgeMobile":    r(in.Browsers.EdgeMobile),
			"Firefox":       r(in.Browsers.Firefox),
			"IOSChrome":     r(in.Browsers.IOSChrome),
			"IOSSafari":     r(in.Browsers.IOSSafari),
			"IE":            r(in.Browsers.IE),
			"IEMobile":      r(in.Browsers.IEMobile),
			"Opera":         r(in.Browsers.Opera),
			"OperaMini":     r(in.Browsers.OperaMini),
			"Safari":        r(in.Browsers.Safari),
			"FirefoxMobile": r(in.Browsers.FirefoxMobile),
			"SamsungMobile": r(in.Browsers.SamsungMobile),
		},
	}
}

type polyfillTOMLConfig struct {
	Deps     []string `toml:"dependencies"`
	Browsers struct {
		Android       string `toml:"android"`
		BB            string `toml:"bb"`
		Chrome        string `toml:"chrome"`
		Edge          string `toml:"edge"`
		EdgeMobile    string `toml:"edge_mob"`
		Firefox       string `toml:"firefox"`
		IOSChrome     string `toml:"ios_chr"`
		IOSSafari     string `toml:"ios_saf"`
		IE            string `toml:"ie"`
		IEMobile      string `toml:"ie_mob"`
		Opera         string `toml:"opera"`
		OperaMini     string `toml:"op_mini"`
		Safari        string `toml:"safari"`
		FirefoxMobile string `toml:"firefox_mob"`
		SamsungMobile string `toml:"samsung_mob"`
	} `toml:"browsers"`
}

func minify(data []byte) string {
	result := api.Transform(string(data), api.TransformOptions{
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
	})
	return string(result.Code)
}

func main() {
	workDir, err := ioutil.TempDir("", "polyfill")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "clone", "--depth", "1", "--single-branch", "--no-tags", "git@github.com:Financial-Times/polyfill-library.git", workDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	rootDir := filepath.Join(workDir, "polyfills")

	confs := map[string]polyfillTOMLConfig{}
	codes := map[string]string{}
	dirs := []string{rootDir}
	for len(dirs) > 0 {
		dir := dirs[0]
		dirs = dirs[1:]
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			absF := filepath.Join(dir, f.Name())
			if f.IsDir() {
				dirs = append(dirs, absF)
				continue
			}

			name := func() string {
				relP, err := filepath.Rel(rootDir, dir)
				if err != nil {
					panic(err)
				}
				return strings.ReplaceAll(relP, "/", ".")
			}

			if f.Name() == "config.toml" {
				file, err := os.Open(absF)
				if err != nil {
					panic(err)
				}
				var conf polyfillTOMLConfig
				if _, err := toml.DecodeReader(file, &conf); err != nil {
					panic(err)
				}
				confs[name()] = conf
			} else if f.Name() == "polyfill.js" {
				data, err := ioutil.ReadFile(absF)
				if err != nil {
					panic(err)
				}
				codes[name()] = minify(data)
			}
		}
	}

	var polyfills []polyfill
	for name, conf := range confs {
		code, ok := codes[name]
		if !ok {
			// TODO(dave): we need to build these from the node_modules, but it's more complicated :(
			// For now, just get fetch working
			data, err := ioutil.ReadFile(filepath.Join("cmd", "generate", name+".js"))
			if err != nil {
				fmt.Println(name + " code not found!")
			} else {
				code = minify(data)
			}
		}
		polyfills = append(polyfills, polyfill{
			Name: name,
			Code: code,
			Conf: fromTOML(conf),
		})
	}

	sort.Slice(polyfills, func(i, j int) bool {
		return polyfills[i].Name < polyfills[j].Name
	})

	out, err := os.Create("generated_code.go")
	if err != nil {
		panic(err)
	}
	err = codeTmpl.Execute(out, polyfills)
	if err != nil {
		panic(err)
	}

	out, err = os.Create("generated_structs.go")
	if err != nil {
		panic(err)
	}
	err = polyTmpl.Execute(out, polyfills)
	if err != nil {
		panic(err)
	}
}
