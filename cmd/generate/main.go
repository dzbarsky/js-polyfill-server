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
{{ range . }}
var {{var .Name}} = &Polyfill{
	code: {{var .Name}}_Code,
	deps: []*Polyfill{ {{range .Conf.Deps}} {{var .}}, {{end}} },
	versions: LastVersionBeforeSupport{},
}
{{ end }}

var polyfillsByName = map[string]*Polyfill{
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
	Deps              []string
	FirefoxMinVersion int
	ChromeMinVersion  int
}

func fromTOML(in polyfillTOMLConfig) polyfillConfig {
	return polyfillConfig{
		Deps: in.Deps,
	}
}

type polyfillTOMLConfig struct {
	Deps     []string `toml:"dependencies"`
	Browsers struct {
		Firefox string `toml:"firefox"`
		Chrome  string `toml:"chrome"`
		// TODO(dave): support the rest of these
		/*android = "<5"
		bb = "*"
		chrome = "<54"
		edge = "<17"
		edge_mob = "<17"
		firefox = "<49"
		ios_chr = "*"
		ios_saf = "<10"
		ie = "6 - *"
		ie_mob = "10 - *"
		opera = "<39"
		op_mini = "*"
		safari = "<10"
		firefox_mob = "<49"
		samsung_mob = "<6"*/
	} `toml:"browsers"`
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
				result := api.Transform(string(data), api.TransformOptions{
					MinifyWhitespace:  true,
					MinifyIdentifiers: true,
					MinifySyntax:      true,
				})
				codes[name()] = string(result.Code)
			}
		}
	}

	var polyfills []polyfill
	for name, conf := range confs {
		code, ok := codes[name]
		if !ok {
			// TODO(dave): we need to build these, but it's more complicated :(
			fmt.Println(name + " code not found!")
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
