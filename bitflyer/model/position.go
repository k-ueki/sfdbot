package model

type (
	Position struct{
		ProductCode         string
		Side                string
		Price               float64
		size                float64
		Commission          float64
		SwapPointAccumulate float64
		RequireCollateral   float64
		OpenDate            string
		Leverage            float64
		Pnl                 float64
		Std                 float64
	}
)
