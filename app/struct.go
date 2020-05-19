package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

//Structs func
func Structs() *Struct {
	return &Struct{}
}

//New struct func
func (s *Struct) New(module *Module) *Struct {
	var structType string = "additional"
	if module.Model.Name == "" {
		structType = "model"
	}
	fmt.Println("Creating new", structType, "structure")
	s.readInput(module)
	return s
}

func (s *Struct) readInput(module *Module) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		s.Name = module.Name
		if s.Name == "" || module.Model.Name != "" {
			s.Name = GetInput("Enter struct name : ", *scanner, true)
		}

		nemu := false
		if module.Model.Name == s.Name {
			nemu = true
		}
		for _, strct := range module.AddStructs {
			if strct.Name == s.Name {
				nemu = true
				break
			}
		}
		if nemu {
			fmt.Println("Struct " + strcase.ToCamel(s.Name) + " is already exist.")
			continue
		} else {
			break
		}
	}
	for {
		var structure = &Structure{}
		// fmt.Print("Enter structure name (empty for finish): ")
		// fmt.Scan(&structure.Name)
		// if strings.TrimSpace(structure.Name) == "" {
		// 	break
		// }

		structure.Name = GetInput("Enter structure variable name (empty for finish): ", *scanner, false)

		if len(structure.Name) == 0 {
			break
		}

		// fmt.Print("Enter structure type : ")
		// fmt.Scanln(&structure.Type)

		structure.Type = GetInput("Enter structure variable type : ", *scanner, true)
		// structure := make(map[string]string)
		// structure[sname] = stype
		s.Structures = append(s.Structures, *structure)
	}
}
