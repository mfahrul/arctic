package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

//Modules func
func Modules() *Module {
	return &Module{}
}

//New Func
func (m *Module) New(project *Project) *Module {
	if !IsDirExist(TemplateDir) {
		fmt.Println("Template directory does not exit. Cloning it from repository")
		// Clone the given repository to the given directory
		// _, cloneError := git.PlainClone(TemplateDir, false, &git.CloneOptions{
		// 	URL:      TempRepo,
		// 	Progress: os.Stdout,
		// })
		Execute(Curdir(), "git", "clone", TempRepo, TemplateDir)
		fmt.Println("Using " + TemplateDir + " as template directory")
		fmt.Println("=================================")
	}

	fmt.Println("Creating new module")
	m.readInput()
	nemu := false
	for _, module := range project.Modules {
		if module.Name == strings.ToLower(m.Name) {
			m = &module
			nemu = true
			break
		}
	}

	if nemu {
		Amsyong("Module " + strcase.ToCamel(m.Name) + " is already exist.")
	}
	return m
}

func (m *Module) readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	m.Name = inflection.Singular(strings.ToLower(GetInput("Enter module name : ", *scanner, true)))
}

//CopyModule func
func (m *Module) CopyModule() {
	err := CopyDir(TemplateDir+string(filepath.Separator)+"App"+string(filepath.Separator)+"src"+string(filepath.Separator)+"core", WorkDir+"App"+string(filepath.Separator)+"src"+string(filepath.Separator)+m.Name)
	if err != nil {
		panic(err)
	}
}
