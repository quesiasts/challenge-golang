package main

type Storage interface {
	SaveAccount(newAccounts ...Account) error
	UpdateAccount(changedAccounts ...Account) error
	SaveTransfer(newTransfers ...Transfer) error
	FindAccount(id int) Account
	FindAccountByCpf(cpf string) Account
	FindTransfers(accountId int) Transfers
	FindAccounts() Accounts
	SaveToken(newTokens ...Token) error
	FindTokens() Tokens
}

type StorageMemory struct {
	accounts  Accounts
	transfers map[int]Transfers
	token     Tokens
}

func CreateNewStorage() Storage {
	return &StorageMemory{
		make(Accounts, 0),
		make(map[int]Transfers),
		make(Tokens, 0),
	}
}

func (s *StorageMemory) SaveAccount
