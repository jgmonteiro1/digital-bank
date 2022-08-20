package usecase

import "com.jgmonteiro.bank/domain"

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func newUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}
