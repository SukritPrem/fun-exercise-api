package wallet

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
	WalletType(wallet_type string) ([]Wallet, error)
	GetWalletSpecificByUserId(id string) ([]Wallet, error)
	CreateWallet(wallet Wallet) error
	UpdateWallet(wallet Wallet) error
	DeleteWallet(id string) error
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {
	var (
        wallets []Wallet // Assuming Wallet is the type of your wallet objects
        err     error
    )

	wallet_type := c.QueryParam("wallet_type");
	fmt.Printf("wallet_type: %s\n", wallet_type)
	if wallet_type != "" {
		wallets, err = h.store.WalletType(wallet_type);
	} else {
		wallets, err = h.store.Wallets()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// GetWalletSpecificByUserId
//	@Summary		Get wallet specific by user id
//	@Description	Get wallet specific by user id
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"user id"
//	@Success		200	{object}	Wallet
//	@Router			/users/{id}/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) GetWalletSpecificByUserIdHandler(c echo.Context) error {
	id := c.Param("id")
	wallets, err := h.store.GetWalletSpecificByUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// CreateWallet
//	@Summary		Create wallet
//	@Description	Create wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet	body	Wallet	true	"wallet"
//	@Success		201	{object}	Wallet
//	@Router			/api/v1/wallets [post]
//	@Failure		400	{object}	Err
//	@Failure		500	{object}	Err
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	wallet := Wallet{}
	err := c.Bind(&wallet)
	fmt.Printf("wallet: %v\n", wallet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	// fmt.Println(wallet)
	err = h.store.CreateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, wallet)
}

// UpdateWallet
//	@Summary		Update wallet
//	@Description	Update wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet	body	Wallet	true	"wallet"
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [put]
//	@Failure		400	{object}	Err
//	@Failure		500	{object}	Err
func (h *Handler) UpdateWalletHandler(c echo.Context) error {
	wallet := Wallet{}
	err := c.Bind(&wallet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err = h.store.UpdateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}

// DeleteWallet
//	@Summary		Delete wallet
//	@Description	Delete wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"wallet id"
//	@Success		200	{string}	string
//	@Router			/api/v1/wallets/{id} [delete]
//	@Failure		500	{object}	Err
func (h *Handler) DeleteWalletHandler(c echo.Context) error {
	id := c.Param("id")
	err := h.store.DeleteWallet(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, id)
}