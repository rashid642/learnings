package Database

import (
	"context"
	"database/sql"
	"fmt"
)

// store function provide all the function to execute db and transaction
type Store struct {
	*Queries 
	db *sql.DB 
}

// Create a new store
func NewStore(db *sql.DB) *Store{
	return &Store{
		db : db,
		Queries: New(db),
	}
}

// execTx execute a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) 
	if err != nil {
		return err 
	}
	q := New(tx) 

	err = fn(q)
	if err != nil {
		if rollBackError := tx.Rollback(); rollBackError != nil {
			return fmt.Errorf("Transaction Error: %v, RollbackError: %v", err, rollBackError)
		}
		return err 
	}

	return tx.Commit()
}

// this contains all the input paramters to transfer money from to account
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

// Transafer transaction result 
type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

// Transfer perform money from one account to other 
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult 

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 
		
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID: arg.ToAccountId,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})	

	return result, err
}