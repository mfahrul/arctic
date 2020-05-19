package app

import "os"

//TemplateDir const
var TemplateDir string

// const TemplateDir string = "/Users/muhammadfahrul/WORK/FAHRUL/arctic-tpl"

//TempRepo const
const TempRepo string = "https://github.com/mfahrul/arctic-tpl"

//CfgFileName const
const CfgFileName string = ".arctic"

//CfgPath var
var CfgPath string

//WorkDir var
var WorkDir string

func init() {
	CfgPath = CfgFileName + ".yaml"
	TemplateDir = os.TempDir() + "arctic"
}
