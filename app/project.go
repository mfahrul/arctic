package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/spf13/viper"
)

//Projects func
func Projects() *Project {
	return &Project{}
}

//New func
func (p *Project) New() {
	fmt.Println("Creating new project with name " + p.Projectname)
	fmt.Println("Checking template directory")

	if !IsDirExist(TemplateDir) {
		fmt.Println("Template directory does not exit. Cloning it from repository")
		// Clone the given repository to the given directory
		// _, cloneError := git.PlainClone(TemplateDir, false, &git.CloneOptions{
		// 	URL:      TempRepo,
		// 	Progress: os.Stdout,
		// })
		Execute(Curdir(), "git", "clone", TempRepo, TemplateDir)
	}

	fmt.Println("Using " + TemplateDir + " as template directory")
	fmt.Println("=================================")
	p.readInput()
}

//Save func
func (p *Project) Save() {
	defer viper.WriteConfigAs(CfgPath)
	defer viper.Set("project", p)
}

//ParseProject func
func (p *Project) ParseProject() {
	err := CopyDir(TemplateDir+string(filepath.Separator)+"App", WorkDir+"App")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = filepath.Walk(WorkDir+"App", p.walkFunc)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", TemplateDir, err)
		return
	}
	Execute(WorkDir+"App", "swag", "init")
}

func (p *Project) walkFunc(path string, info os.FileInfo, err error) error {
	subDirToSkip := "tmp"
	if err != nil {
		log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	if info.IsDir() && info.Name() == subDirToSkip {
		fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
		return filepath.SkipDir
	}
	if !info.IsDir() {
		fmt.Printf("Parsing file : %q\n", path)
		funcMap := template.FuncMap{
			"ToCamel":    strcase.ToCamel,
			"ToSnake":    strcase.ToSnake,
			"ToLower":    strings.ToLower,
			"ToPlural":   inflection.Plural,
			"ToSingular": inflection.Singular,
		}
		t := template.Must(template.New(info.Name()).Funcs(funcMap).Delims("[[", "]]").ParseFiles(path))

		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer f.Close()

		err = t.Execute(f, p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return nil
}

func (p *Project) readInput() {
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Enter database name [giftanoDb]: ")
	// fmt.Scanln(&p.Dbname)
	p.Dbname = GetInput("Enter database name [giftanoDb] : ", *scanner, false)
	if p.Dbname == "" {
		p.Dbname = "giftanoDb"
	}
	// fmt.Print("Enter database username [root]: ")
	// fmt.Scanln(&p.Dbusername)
	p.Dbusername = GetInput("Enter database username [root] : ", *scanner, false)
	if p.Dbusername == "" {
		p.Dbusername = "root"
	}
	// fmt.Print("Enter database password [Giftano_Id]: ")
	// fmt.Scanln(&p.Dbpassword)
	p.Dbpassword = GetInput("Enter database password [Giftano_Id] : ", *scanner, false)
	if p.Dbpassword == "" {
		p.Dbpassword = "Giftano_Id"
	}
	// fmt.Print("Enter database host [localhost]: ")
	// fmt.Scanln(&p.Dbhost)
	p.Dbhost = GetInput("Enter database host [localhost] : ", *scanner, false)
	if p.Dbhost == "" {
		p.Dbhost = "localhost"
	}
	// fmt.Print("Enter database port number [27017]: ")
	// fmt.Scanln(&p.Dbport)
	p.Dbport = GetInput("Enter database port number [27017] : ", *scanner, false)
	if p.Dbport == "" {
		p.Dbport = "27017"
	}
}
