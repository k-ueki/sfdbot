package application

type (
	SfdPosition struct {
		HavePosition *HavePosition
		AcceptanceID *string
	}

	HavePosition struct {
		Buy  bool
		Sell bool
	}
)

var sfdPosition SfdPosition

func (p *SfdPosition) Reset() {
	p.HavePosition = nil
}

func (p *SfdPosition) SetSellPosition(orderAcceptanceID string) {
	sfdPosition.HavePosition = &HavePosition{
		Buy:  false,
		Sell: true,
	}
	sfdPosition.AcceptanceID = &orderAcceptanceID
}
