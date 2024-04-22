package wallet

import "testing"

import (
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"net/http"
	"reflect"
	"time"
	"encoding/json"
)

type StubHandler struct {
	handler Handler
	err     error
}


func (s StubHandler) Wallets() ([]Wallet, error) {
	mockTime := time.Date(2024, time.April, 3, 10, 0, 0, 0, time.UTC)
	wallets := []Wallet{
		{ID: 1, UserID: 1, UserName: "user1", WalletName: "wallet1", WalletType: "type1", Balance: 1000, CreatedAt: mockTime},
	}
	return wallets, s.err
}



func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := StubHandler{err: echo.ErrInternalServerError}
		h := New(stubError)

		h.WalletHandler(c)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := StubHandler{}

		h := New(stubError)
		mockTime := time.Date(2024, time.April, 3, 10, 0, 0, 0, time.UTC)
		want := []Wallet{	{ID: 1, UserID: 1, UserName: "user1", WalletName: "wallet1", WalletType: "type1", Balance: 1000, CreatedAt:  mockTime}}
		h.WalletHandler(c)
		gotjson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotjson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
