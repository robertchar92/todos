package models

type Order struct {
	Type           string `json:"type"`
	OrderNumber    string `json:"order_number"`
	OrderVerb      string `json:"order_verb"`
	Quantity       string `json:"quantity"`
	OrderBook      string `json:"order_book"`
	Price          string `json:"price"`
	StockCode      string `json:"stock_code"`
	ExecutedQty    string `json:"executed_quantity"`
	ExecutionPrice string `json:"execution_price"`
}
