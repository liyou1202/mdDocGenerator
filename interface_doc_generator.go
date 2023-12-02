package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type pkg struct {
	packagePath string
	goFiles     []goFile
}

type goFile struct {
	goFilePath string
	pubFuncs   []pubFunc
}

type pubFunc struct {
	annotation   string
	name         string
	fullName     string
	isTestPrefix bool
}

const rootPath = "../"
const searchPath = "pkg/"
const exportPath = "docs/interface.md"

var pkgPath string
var allPackages []pkg

func main() {
	pkgPath = rootPath + searchPath

	var findAllPkgs = func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != pkgPath {
			allPackages = append(allPackages, pkg{packagePath: path})
		}
		return nil
	}

	err := filepath.Walk(pkgPath, findAllPkgs)
	if err != nil {
		return
	}

	//Scan all Packages
	for i, p := range allPackages {
		var findGoFile = func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(path) == ".go" {
				p.goFiles = append(p.goFiles, goFile{
					goFilePath: path,
				})
			}
			return nil
		}

		err = filepath.Walk(p.packagePath, findGoFile)
		if err != nil {
			return
		}

		//find all .go files
		for j, f := range p.goFiles {
			p.goFiles[j].pubFuncs = findPublicFunc(f.goFilePath)
		}
		allPackages[i] = p
	}

	writeBook()
}

func writeBook() {
	book := NewDoc()

	// Title
	book.AddTitle("接口文件", 1).
		AddContent("### 此文件敘述各個 Package 的開放接口 (public function) 和內容說明").
		AddBlankLines(2)

	book.AddInterval().
		AddTitle("Packages 索引", 2).
		AddBlankLines(1)

	// Index
	for _, p := range allPackages {
		book.AddContent("###  - ").
			AddLink(filepath.Base(p.packagePath), "#"+filepath.Base(p.packagePath)).
			AddBlankLines(2)
	}

	allPackages = excludePrivatePackage(allPackages)
	// Content
	for _, p := range allPackages {
		book.AddInterval().
			AddTitle(filepath.Base(p.packagePath), 2)

		for _, f := range p.goFiles {

			for _, pfunc := range f.pubFuncs {
				if pfunc.isTestPrefix {
					continue
				}

				annotation := fmt.Sprintf("* %s", pfunc.annotation)
				if len(pfunc.annotation) == 0 {
					annotation = fmt.Sprintf("* %s", "這個 Func 沒有註解")
				}

				book.AddCodeBlock(pfunc.fullName, "go").
					AddContent(annotation).
					AddBlankLines(2)
			}
		}
	}

	err := book.Export(rootPath + exportPath)
	fmt.Println("procedure done...")
	if err != nil {
		log.Fatal(err)
	}

}

func excludePrivatePackage(pkgs []pkg) []pkg {
	var excluded []pkg
	havePubFunc := false
	for _, p := range pkgs {
		for _, f := range p.goFiles {
			if len(f.pubFuncs) != 0 {
				havePubFunc = true
				break
			}
		}

		if havePubFunc {
			excluded = append(excluded, p)
		}
		havePubFunc = false
	}
	return excluded
}

func findPublicFunc(path string) (funcs []pubFunc) {
	//逐行取得檔案內 function
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(`(^func(\s|\s\(.+\)\s)([A-Z]\w+)\(.+){`)

	for scanner.Scan() {
		matches := reg.FindStringSubmatch(scanner.Text())
		if len(matches) > 3 {

			funcs = append(funcs, pubFunc{
				fullName:     matches[1],
				name:         matches[3],
				isTestPrefix: checkTestPrefix(matches[3]),
			})
		}
	}

	//  Retrieve annotate
	//  Make sure to annotate at the beginning with functionName followed by a space.
	//  Use line breaks above and below /* as shown in the example below:

	//	/*
	//	functionName someText....
	//	first line
	//	n lines...
	//	*/
	b, _ := os.ReadFile(path)

	for i, v := range funcs {
		rule := fmt.Sprintf(`\/\*\n%s\s+([^(\*\/)]+|\s|\n)+\*\/`, v.name)
		regex := regexp.MustCompile(rule)
		matches := regex.FindStringSubmatch(string(b))
		if len(matches) >= 2 {
			funcs[i].annotation = matches[1]
		}
	}

	return
}

func checkTestPrefix(name string) bool {
	reg := regexp.MustCompile(`^Test.*`)
	matches := reg.FindStringSubmatch(name)
	if len(matches) > 0 {
		return true
	}
	return false
}

type MarkdownDoc struct {
	content *strings.Builder
}

func NewDoc() *MarkdownDoc {
	return &MarkdownDoc{
		content: &strings.Builder{},
	}
}

func (doc *MarkdownDoc) WriteLine(content string) {
	doc.content.WriteString(content + "\n")
}

func (doc *MarkdownDoc) AddTitle(t string, lv int) *MarkdownDoc {
	if lv > 6 || lv < 1 {
		fmt.Printf("failed to add Title %s in level: %d", t, lv)
		return doc
	}
	mdSyntax := strings.Repeat("#", lv) + " " + t
	doc.WriteLine(mdSyntax)
	return doc
}

func (doc *MarkdownDoc) AddInterval() *MarkdownDoc {
	mdSyntax := strings.Repeat("-", 3) + " "
	doc.WriteLine(mdSyntax)
	return doc
}

func (doc *MarkdownDoc) AddImage(placeholder, path, title string) *MarkdownDoc {
	mdSyntax := fmt.Sprintf("![%s](%s) %s", placeholder, path, title)
	doc.WriteLine(mdSyntax)
	return doc
}

func (doc *MarkdownDoc) AddBlankLines(lv int) *MarkdownDoc {
	if lv > 0 {
		for i := 1; i <= lv; i++ {
			doc.WriteLine("")
		}
	}

	return doc
}

func (doc *MarkdownDoc) AddCodeBlock(code, language string) *MarkdownDoc {
	mdSyntax := fmt.Sprintf("``` %s\n%s\n```", language, code)
	doc.WriteLine(mdSyntax)

	return doc
}

func (doc *MarkdownDoc) AddLink(text, path string) *MarkdownDoc {
	mdSyntax := fmt.Sprintf("[%s](%s)", text, path)
	doc.content.WriteString(mdSyntax)

	return doc
}

func (doc *MarkdownDoc) AddContent(text string) *MarkdownDoc {
	doc.content.WriteString(text)

	return doc
}

func (doc *MarkdownDoc) Export(filename string) error {
	return os.WriteFile(filename, []byte(doc.content.String()), os.ModePerm)
}
