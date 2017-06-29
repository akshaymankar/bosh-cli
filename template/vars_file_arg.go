package template

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"gopkg.in/yaml.v2"
)

type VarsFileArg struct {
	FS boshsys.FileSystem

	Vars StaticVariables
}

func (a *VarsFileArg) UnmarshalFlag(filePath string) error {
	//TOOD: Inject FS somehow
	if a.FS == nil {
		logger := boshlog.NewLogger(boshlog.LevelInfo)
		a.FS = boshsys.NewOsFileSystemWithStrictTempRoot(logger)
	}

	if len(filePath) == 0 {
		return bosherr.Errorf("Expected file path to be non-empty")
	}

	bytes, err := a.FS.ReadFile(filePath)
	if err != nil {
		return bosherr.WrapErrorf(err, "Reading variables file '%s'", filePath)
	}

	var vars StaticVariables

	err = yaml.Unmarshal(bytes, &vars)
	if err != nil {
		return bosherr.WrapErrorf(err, "Deserializing variables file '%s'", filePath)
	}

	(*a).Vars = vars

	return nil
}
