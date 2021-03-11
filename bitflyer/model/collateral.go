package model

type (
// change: 証拠金の変動額です。
// amount: 変動後の証拠金の残高です。
	Collateral struct{
		ID int64 `json:"id"`
		Code string `json:"currency_code"`
		Change float64 `json:"change"`
		Amount float64 `json:"amount"`
		Date string `json:"date"`
	}
)
