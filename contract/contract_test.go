package contract

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/qlcchain/go-lsobus/cmd/util"
	"github.com/qlcchain/go-lsobus/config"
)

func setupTestCase(t *testing.T) (func(t *testing.T), *ContractService) {
	cfgFile := filepath.Join(config.TestDataDir(), uuid.New().String(), config.CfgFileName)
	cs, err := NewContractService(cfgFile)
	if err != nil {
		t.Fatal(err)
	}
	setupOrchestraConfig(cs)
	if err = cs.Init(); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		err = cs.Stop()
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(filepath.Join(config.TestDataDir(), uuid.New().String()))
		if err != nil {
			t.Fatal(err)
		}
	}, cs
}

func setupOrchestraConfig(cs *ContractService) {
	cs.cfg.Partners = nil
	p1 := &config.PartnerCfg{
		Name:      "PCCWG",
		ID:        "PCCWG",
		SonataUrl: "http://127.0.0.1:7777",
		Username:  "test",
		Password:  "test",
	}
	cs.cfg.Partners = append(cs.cfg.Partners, p1)
}

func TestContractService_GetOrderInfoByInternalId(t *testing.T) {
	teardownTestCase, cs := setupTestCase(t)
	defer teardownTestCase(t)
	orderInfo, err := cs.GetOrderInfoByInternalId(uuid.New().String())
	if err == nil || err != chainNotReady {
		t.Fatal("chain is not ready")
	}
	cs.SetFakeMode(true)
	orderInfo, err = cs.GetOrderInfoByInternalId(uuid.New().String())
	if err != nil {
		t.Fatal("fake mode should return orderInfo and no error")
	}
	t.Log(util.ToIndentString(orderInfo))
}