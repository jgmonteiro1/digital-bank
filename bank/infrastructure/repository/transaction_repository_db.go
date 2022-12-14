package repository

import (
	"database/sql"

	"com.jgmonteiro.bank/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(`insert into transactions(id, credit_card_id, amount, status, description, store, created_at)
	values($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)

	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = t.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDb) updateBalance(creditCard domain.CreditCard) error {
	_, err := t.db.Exec("update credit_cards set balance = $1 where id = $2",
		creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}
