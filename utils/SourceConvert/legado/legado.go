package legado

import (
	"SourceConvert/tools"
	"encoding/json"
	"os"
)

type SourceConvert struct {
	srcBookRule []BookSource
}

func New(fileName string) *SourceConvert {
	fileData, err := os.ReadFile(fileName)
	tools.PanicIfError(err)
	var sourceSlice []BookSource
	err = json.Unmarshal(fileData, &sourceSlice)
	tools.PanicIfError(err)
	return &SourceConvert{
		srcBookRule: sourceSlice,
	}
}
