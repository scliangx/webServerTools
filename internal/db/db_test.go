package db

import (
	"github.com/scliang-strive/webServerTools/config"
	"testing"
)

func DbTest(t *testing.T){
	cfg := config.GetConfig().DB
	NewConnection(cfg)

}
