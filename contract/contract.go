package contract

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/qlcchain/go-lsobus/mock"

	"github.com/qlcchain/go-lsobus/orchestra"

	qlcSdk "github.com/qlcchain/qlc-go-sdk"
	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"go.uber.org/zap"

	"github.com/qlcchain/go-lsobus/common"
	"github.com/qlcchain/go-lsobus/common/event"
	"github.com/qlcchain/go-lsobus/config"
	"github.com/qlcchain/go-lsobus/log"
	ct "github.com/qlcchain/go-lsobus/services/context"
)

const (
	checkNeedSignContractInterval = 15 * time.Second
	checkContractStatusInterval   = 10 * time.Second
	checkOrderStatusInterval      = 10 * time.Second
	checkProductInterval          = 10 * time.Second
	connectRpcServerInterval      = 5 * time.Second
)

var (
	chainNotReady     = errors.New("chain is not ready")
	buyerAddrNotMatch = errors.New("buyer address not match")
)

type ContractService struct {
	cfg               *config.Config
	account           *types.Account
	logger            *zap.SugaredLogger
	client            *qlcSdk.QLCClient
	ctx               context.Context
	cancel            context.CancelFunc
	handlerIds        map[common.TopicType]string
	eb                event.EventBus
	chainReady        bool
	quit              chan bool
	orderIdOnChain    *sync.Map
	orderIdFromSonata *sync.Map
	orchestra         *orchestra.Orchestra
	fakeMode          bool
}

type Product struct {
	buyerProductID string
	productID      string
}

func NewContractService(cfgFile string) (*ContractService, error) {
	cc := ct.NewServiceContext(cfgFile)
	cfg, err := cc.Config()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	or := orchestra.NewOrchestra(cfgFile)
	or.SetFakeMode(true)
	cs := &ContractService{
		cfg:               cfg,
		account:           cc.Account(),
		logger:            log.NewLogger("contract"),
		ctx:               ctx,
		cancel:            cancel,
		handlerIds:        make(map[common.TopicType]string),
		eb:                cc.EventBus(),
		quit:              make(chan bool, 1),
		orderIdOnChain:    new(sync.Map),
		orderIdFromSonata: new(sync.Map),
		orchestra:         or,
	}
	return cs, nil
}

func (cs *ContractService) SetFakeMode(mode bool) {
	cs.fakeMode = mode
}

func (cs *ContractService) GetFakeMode() bool {
	return cs.fakeMode
}

func (cs *ContractService) GetAccount() *types.Account {
	return cs.account
}

func (cs *ContractService) SetAccount(account *types.Account) {
	cs.account = account
}

func (cs *ContractService) Init() error {
	err := cs.orchestra.Init()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContractService) Start() error {
	go cs.checkDoDContract()
	go cs.connectRpcServer()
	go cs.checkContractStatus()
	go cs.checkOrderStatus()
	go cs.checkProduct()
	return nil
}

func (cs *ContractService) GetOrderInfoByInternalId(id string) (*qlcSdk.DoDSettleOrderInfo, error) {
	if cs.GetFakeMode() {
		return mock.GetOrderInfoByInternalId(id)
	}
	if cs.chainReady {
		orderInfo, err := cs.client.DoDSettlement.GetOrderInfoByInternalId(id)
		if err != nil {
			cs.logger.Error(err)
			return nil, err
		}
		return orderInfo, nil
	} else {
		return nil, chainNotReady
	}
}

func (cs *ContractService) Stop() error {
	//this must be the first step
	cs.cancel()
	err := cs.unsubscribeEvent()
	if err != nil {
		return err
	}
	if cs.client != nil {
		_ = cs.client.Close
	}
	return nil
}

func (cs *ContractService) unsubscribeEvent() error {
	for k, v := range cs.handlerIds {
		if err := cs.eb.Unsubscribe(k, v); err != nil {
			return err
		}
	}
	return nil
}
