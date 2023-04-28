package db

import (
	"github.com/coderitx/webServerTools/config"
	"testing"
)

func DbTest(t *testing.T){
	cfg := config.GetConfig().DB
	NewConnection(cfg)

}
