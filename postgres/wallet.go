package postgres

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"time"
	"fmt"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets() ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) WalletType(wallet_type string) ([]wallet.Wallet, error) {
	fmt.Printf("wallet_type: %s type: %T\n", wallet_type, wallet_type)
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type=$1", wallet_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) GetWalletSpecificByUserId(user_id string) ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE user_id=$1", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) CreateWallet(w wallet.Wallet) error  {

    query := `
        INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := p.Db.Exec(query, w.UserID, w.UserName, w.WalletName, w.WalletType, w.Balance, w.CreatedAt)
	if err != nil {	
		return err
	}
	return nil
}

func (p *Postgres) UpdateWallet(w wallet.Wallet) error {
	_, err := p.Db.Exec("UPDATE user_wallet SET user_id=$1, user_name=$2, wallet_name=$3, wallet_type=$4, balance=$5, created_at=$6 WHERE id=$7",
		w.UserID, w.UserName, w.WalletName, w.WalletType, w.Balance, w.CreatedAt, w.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteWallet(id string) error {
	_, err := p.Db.Exec("DELETE FROM user_wallet WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}