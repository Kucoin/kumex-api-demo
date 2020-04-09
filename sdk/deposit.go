package sdk

import (
	"net/http"
)

// A DepositAddressModel represents a deposit address of currency for deposit.
type DepositAddressModel struct {
	Address string `json:"address"`
	Memo    string `json:"memo"`
}

// A DepositAddressesModel is the set of *DepositAddressModel.
type DepositAddressesModel []*DepositAddressModel

// DepositAddresses returns the deposit address of currency for deposit.
// If return data is empty, you may need create a deposit address first.
func (as *ApiService) DepositAddresses(currency string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	req := NewRequest(http.MethodGet, "/api/v1/deposit-addresses", params)
	return as.Call(req)
}

// A DepositModel represents a deposit record.
type DepositModel struct {
	Currency   string `json:"currency"`
	Status     string `json:"status"`
	Address    string `json:"address"`
	IsInner    bool   `json:"isInner"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	WalletTxId string `json:"walletTxId"`
	CreatedAt  int64  `json:"createdAt"`
}

// A DepositsModel represents a deposit list.
type DepositsModel []*DepositModel

// Deposits returns a list of deposit.
func (as *ApiService) Deposits() (*ApiResponse, error) {
	params := map[string]string{}
	pagination := &PaginationParam{CurrentPage: 1, PageSize: 10}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/deposit-list", params)
	return as.Call(req)
}
