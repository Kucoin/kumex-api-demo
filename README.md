

# Test Tool for KuMex API

## Choose environment

| Environment  | BaseUri                                                  |
| ------------ | -------------------------------------------------------- |
| *Production* | `https://api.kumex.com(DEFAULT)` `https://api.kumex.top` |
| *Sandbox*    | `https://sandbox-api.kumex.com`                          |

## Example

* First,you should wirte your api information into config.json.

  ```json
  {
    "ApiBaseURI":"",
    "ApiKey": "",
    "ApiSecret": "",
    "ApiPassphrase": "",
    "ApiSkipVerifyTls": false
  }
  ```
  
* If not, you can run this script and follow the guide,then input your api information, such as:

  ```shell
  # please input api base URI,such as:https://api.kumex.com
  % https://api.kumex.com
  # please input your api key...
  % input your api key
  # please input your api secret...
  % input your api secret
  # please input your api passphrase...
  % input your api passphrase
  # please input the api name which you want to test
  ```
  
 *  input the api name and parameter information to choose api which you want to test.
  
  ```shell
  % AccountOverview
  # please input the api parameters,use ',' to spilt parameters.
  # if it is no parameter api,just enter to skip...
  % XBT
  ```

* Then,it will show the request and response body to you.

## Api map list

<details>
<summary>Account</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| AccountOverview | GET /api/v1/account-overview | currency | XBT | https://docs.kumex.com/#get-account-overview |
| TransactionHistory | GET /api/v1/transaction-history | |  | https://docs.kumex.com/#get-transaction-history |

</details>

<details>
<summary>Deposit</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| #DepositAddresses | GET /api/v1/deposit-address | currency | XBT | https://docs.kumex.com/#get-deposit-address |
| Deposits | GET /api/v1/deposit-list | |  | https://docs.kumex.com/#get-deposits-list |

</details>

<details>
<summary>Withdrawal</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| WithdrawalQuotas | GET /api/v1/withdrawals/quotas | currency | XBT | https://docs.kumex.com/#get-withdrawal-limit |
| ApplyWithdrawal | POST /api/v1/withdrawals | currency,address,amount | XBT,{address},0.01 | https://docs.kumex.com/#apply-withdraw |
| Withdrawals | GET /api/v1/withdrawal-list | |  | https://docs.kumex.com/#get-withdrawal-list |
| CancelWithdrawal | DELETE /api/v1/withdrawals/{withdrawalId} | withdrawalId | {withdrawal ID} | https://docs.kumex.com/#cancel-withdrawal |

</details>

<details>
<summary>Transfer</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| TransferOut | POST /api/v1/transfer-out | amount | 0.01 | https://docs.kumex.com/#transfer-funds-to-kucoin-main-account |
| TransferOutV2 | POST /api/v2/transfer-out | amount, currency | 0.01,BTC | https://docs.kumex.com/#transfer-funds-to-kucoin-main-account-2 |
| TransferList | GET /api/v1/transfer-list | |  | https://docs.kumex.com/#get-transfer-out-request-records-2 |
| CancelTransfer | DELETE /api/v1/cancel/transfer-out | applyId | {Transfer ID} | https://docs.kumex.com/#cancel-transfer-out-request |

</details>

<details>
<summary>Fill</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| Fills | GET /api/v1/fills | |  | https://docs.kumex.com/#get-fills |
| RecentFills | GET /api/v1/recentFills | |  | https://docs.kumex.com/#recent-fills |
| #OpenOrderStatistics | GET /api/v1/openOrderStatistics | symbol | XBTUSDM | https://docs.kumex.com/#active-order-value-calculation |

</details>

<details>
<summary>Order</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| CreateOrder | POST /api/v1/orders | side, symbol,leverage,price,size | buy,XBTUSDM,2,1,1 | https://docs.kumex.com/#place-a-new-order |
| CancelOrder | DELETE /api/v1/orders/{order-id} | orderId | {order ID} | https://docs.kumex.com/#cancel-an-order |
| CancelOrders | DELETE /api/v1/orders | symbol | XBTUSDM | https://docs.kumex.com/#cancel-all-orders |
| StopOrders | DELETE /api/v1/stopOrders | symbol | XBTUSDM | https://docs.kumex.com/#stop-order-mass-cancelation |
| Orders | GET /api/v1/orders | |  | https://docs.kumex.com/#get-order-list |
| Order | GET /api/v1/orders/{order-id} | orderId | {order ID} | https://docs.kumex.com/#get-an-order |
| RecentOrders | GET /api/v1/recentDoneOrders | |  | https://docs.kumex.com/#recent-orders |
| GetStopOrders | GET /api/v1/stopOrders | symbol | XBTUSDM | https://docs.kumex.com/#get-untriggered-stop-order-list |

</details>

<details>
<summary>Market</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| Ticker | GET /api/v1/ticker | symbol | XBTUSDM | https://docs.kumex.com/#get-real-time-ticker |
| Level2Snapshot | GET /api/v1/level2/snapshot | symbol | XBTUSDM | https://docs.kumex.com/#get-full-order-book-level-2 |
| Level2MessageQuery | GET /api/v1/level2/message/query | symbol,start,end | XBTUSDM,100,200 | https://docs.kumex.com/#level-2-pulling-messages |
| Level3Snapshot | GET /api/v1/level3/snapshot | symbol | XBTUSDM | https://docs.kumex.com/#get-full-order-book-level-3 |
| Level3MessageQuery | GET /api/v1/level3/message/query | symbol,start,end | XBTUSDM,1,20 | https://docs.kumex.com/#level-3-pulling-messages|
| TradeHistory | GET /api/v1/trade/history | symbol | XBTUSDM | https://docs.kumex.com/#transaction-history |
| InterestQuery | GET /api/v1/interest/query | symbol | .XBTINT | https://docs.kumex.com/#get-interest-rate-list |
| IndexQuery | GET /api/v1/index/query | symbol | XBTUSDM | https://docs.kumex.com/#get-index-list |
| MarkPrice | GET /api/v1/mark-price/{symbol}/current | symbol | XBTUSDM | https://docs.kumex.com/#get-current-mark-price |
| PremiumQuery | GET /api/v1/premium/query | symbol | XBTUSDM | https://docs.kumex.com/#get-premium-index |
| FundingRate | GET /api/v1/funding-rate/{symbol}/current | symbol | .XBTUSDMFPI8H | https://docs.kumex.com/#get-current-funding-rate |

</details>

<details>
<summary>Symbol</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| ActiveContracts | GET /api/v1/contracts/active | |  | https://docs.kumex.com/#get-open-contract-list |
| #Contracts | GET /api/v1/contracts/{symbol} | symbol | XBTUSDM | https://docs.kumex.com/#get-order-info-of-the-contract |

</details>

<details>
<summary>WebSocket Feed</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| WebSocketPublicToken | POST /api/v1/bullet-public | |  | https://docs.kumex.com/#apply-for-connection-token |
| WebSocketPrivateToken | POST /api/v1/bullet-private | |  | https://docs.kumex.com/#apply-for-connection-token |
| NewWebSocketClient | | |  | https://docs.kumex.com/#websocket-2 |

</details>

<details>
<summary>Time</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| ServerTime | GET /api/v1/timestamp | |  | https://docs.kumex.com/#server-time |

</details>