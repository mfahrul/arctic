package app

//Project struct
type Project struct {
	Projectpath   string   `yaml:"projectpath"`
	Projectname   string   `yaml:"projectname"`
	Dbname        string   `yaml:"dbname"`
	Dbusername    string   `yaml:"dbusername"`
	Dbpassword    string   `yaml:"dbpassword"`
	Dbhost        string   `yaml:"dbhost"`
	Dbport        string   `yaml:"dbport"`
	Modules       []Module `yaml:"modules"`
	ModuleToParse Module   `yaml:"-"`
}

//Module struct
type Module struct {
	Name       string   `yaml:"name"`
	Model      Struct   `yaml:"model"`
	AddStructs []Struct `yaml:"addstructs"`
}

//Struct struct
type Struct struct {
	Name       string      `yaml:"name"`
	Structures []Structure `yaml:"structures"`
}

//Structure struct
type Structure struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}
