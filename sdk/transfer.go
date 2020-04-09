package sdk

import (
	"net/http"
	"time"
)

// A TransferOutModel represents a transfer out record.
type TransferOutModel struct {
	ApplyId string `json:"applyId"`
}

// TransferOut Transfer Funds to KuCoin-Main Account.
func (as *ApiService) TransferOut(amount string) (*ApiResponse, error) {
	p := map[string]string{
		"bizNo":  IntToString(time.Now().UnixNano()),
		"amount": amount,
	}
	req := NewRequest(http.MethodPost, "/api/v1/transfer-out", p)
	return as.Call(req)
}

// A TransferOutModel represents a transfer out record.
type TransferOutV2Model struct {
	ApplyId string `json:"applyId"`
}

// TransferOut Transfer Funds to KuCoin-Main Account.
func (as *ApiService) TransferOutV2(amount, currency string) (*ApiResponse, error) {
	p := map[string]string{
		"bizNo":    IntToString(time.Now().UnixNano()),
		"amount":   amount,
		"currency": currency,
	}
	req := NewRequest(http.MethodPost, "/api/v2/transfer-out", p)
	return as.Call(req)
}

// A TransferModel represents a transfer record.
type TransferModel struct {
	ApplyId   string `json:"applyId"`
	Currency  string `json:"currency"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	Reason    string `json:"reason"`
	Offset    int64  `json:"offset"`
	CreatedAt int64  `json:"createdAt"`
}

// A TransfersModel  represents a transfer list.
type TransfersModel []*TransferModel

// TransferList returns a list of deposit.
func (as *ApiService) TransferList() (*ApiResponse, error) {
	params := map[string]string{}
	pagination := &PaginationParam{CurrentPage: 1, PageSize: 10}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/transfer-list", params)
	return as.Call(req)
}

// CancelTransferModel represents the result of CancelWithdrawal().
type CancelTransferModel struct {
	ApplyId string `json:"applyId"`
}

// CancelTransfer Cancel Transfer-Out Request.
func (as *ApiService) CancelTransfer(applyId string) (*ApiResponse, error) {
	p := map[string]string{
		"applyId": applyId,
	}
	req := NewRequest(http.MethodDelete, "/api/v1/cancel/transfer-out", p)
	return as.Call(req)
}
