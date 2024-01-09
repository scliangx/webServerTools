package db

import (
	"github.com/scliangx/webServerTools/config"
	"testing"
)

func DbTest(t *testing.T){
	cfg := config.GetConfig().DB
	NewConnection(cfg)

}
