package app

import (
	"os/user"
	"path/filepath"
)

//ArcticVersion const
const ArcticVersion string = "V1.0.5"

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
	usr, err := user.Current()
	if err != nil {
		Amsyong(err)
	}
	TemplateDir = filepath.Join(usr.HomeDir, CfgFileName)
}
