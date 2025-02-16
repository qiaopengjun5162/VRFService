package tx_manager

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

// 每发一笔交易的时候会创建一个状态
type SendState struct {
	minedTxs         map[common.Hash]struct{}
	nonceTooLowCount uint64
	mu               sync.RWMutex

	safeAbortNonceTooLowCount uint64
}

func NewSendState(safeAbortNonceTooLowCount uint64) *SendState {
	if safeAbortNonceTooLowCount == 0 {
		panic("tx_manager: safeAbortNonceTooLowCount cannot be zero")
	}

	return &SendState{
		minedTxs:                  make(map[common.Hash]struct{}),
		nonceTooLowCount:          0,
		safeAbortNonceTooLowCount: safeAbortNonceTooLowCount,
	}
}

// 当交易发送完成后会检测是否是nonce太低错误，如果是则计数器加1
func (s *SendState) ProcessSendError(err error) {
	if err == nil {
		return
	}

	if !strings.Contains(err.Error(), core.ErrNonceTooLow.Error()) {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.nonceTooLowCount++
}

// TxMined函数用于将交易哈希添加到已挖矿交易集合中
func (s *SendState) TxMined(txHash common.Hash) {
	// 加锁，防止并发访问
	s.mu.Lock()
	defer s.mu.Unlock()

	// 将交易哈希添加到已挖矿交易集合中
	s.minedTxs[txHash] = struct{}{}
}

// TxNotMined函数用于处理交易未挖矿的情况
func (s *SendState) TxNotMined(txHash common.Hash) {
	// 加锁，防止并发访问
	s.mu.Lock()
	defer s.mu.Unlock()

	// 查找交易是否已经被挖矿
	_, wasMined := s.minedTxs[txHash]
	// 删除交易
	delete(s.minedTxs, txHash)

	// 如果交易已经被挖矿，并且当前没有其他交易被挖矿，则将nonceTooLowCount重置为0
	if len(s.minedTxs) == 0 && wasMined {
		s.nonceTooLowCount = 0
	}
}

// ShouldAbortImmediately函数用于判断是否需要立即中止发送状态
func (s *SendState) ShouldAbortImmediately() bool {
	// 读取锁
	s.mu.RLock()
	// 在函数结束时释放锁
	defer s.mu.RUnlock()

	// 如果已经挖到了交易，则不需要立即中止
	if len(s.minedTxs) > 0 {
		return false
	}
	// 如果nonceTooLowCount大于等于safeAbortNonceTooLowCount，则需要立即中止
	return s.nonceTooLowCount >= s.safeAbortNonceTooLowCount
}

func (s *SendState) IsWaitingForConfirmation() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.minedTxs) > 0
}
