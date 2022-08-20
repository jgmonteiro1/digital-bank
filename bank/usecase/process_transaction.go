package usecase

import (
	"time"

	"com.jgmonteiro.bank/domain"
	"com.jgmonteiro.bank/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func newUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := u.builderNewCreditCard(transactionDto)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.newTransaction(transactionDto, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.SaveTransaction(*t, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

}

func (UseCaseTransaction) builderNewCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.NumberCreditCard = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV
	return creditCard
}

func (UseCaseTransaction) newTransaction(transactionDto dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.newTransaction()
	t.CreditCardId = cc.ID
	t.Amount = transactionDto.Amount
	t.Store = transactionDto.Store
	t.Description = transactionDto.Description
	t.CreatedAt = time.Now()
	return t
}
