package repository

import (
	"fmt"
	"sync"
)

type AccountID string

type Accounts interface {
	PutAccount(Account) error
	GetAccount(AccountID) (Account, error)
}

type memAccounts struct {
	sync.RWMutex
	accounts map[AccountID]Account
}

func (m *memAccounts) PutAccount(a Account) error {
	m.Lock()
	defer m.Unlock()
	m.accounts[a.GetID()] = a
	return nil
}

func (m *memAccounts) GetAccount(a AccountID) (*Account, error) {
	m.RLock()
	defer m.RUnlock()
	account, ok := m.accounts[a]
	if !ok {
		return nil, fmt.Errorf("account not found: %s", a)
	}
	return &account, nil
}

type Account interface {
	GetID() AccountID
}

type account struct {
	Tokens []Token
}

func (a account) GetID() AccountID {
	return AccountID("123")
}

func (a account) checkAccountLimits() error {
	return nil
}

type Token struct {
}
