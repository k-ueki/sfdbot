package application

type (
	SfdPosition struct {
		HavePosition *HavePosition
		HaveOrder    *HavePosition
		AcceptanceID string
	}

	HavePosition struct {
		Buy  bool
		Sell bool
	}
)

var sfdPosition SfdPosition

func init() {
	sfdPosition = SfdPosition{
		HaveOrder: &HavePosition{
			Buy:  false,
			Sell: false,
		},
		AcceptanceID: "",
	}
}

func (p *SfdPosition) Reset() {
	p.HavePosition = nil
	p.HaveOrder = &HavePosition{
		Buy:  false,
		Sell: false,
	}
	p.AcceptanceID = ""
}

func (p *SfdPosition) SetSellPosition(orderAcceptanceID string) {
	sfdPosition.HavePosition = &HavePosition{
		Buy:  false,
		Sell: true,
	}
	//sfdPosition.AcceptanceID = orderAcceptanceID
}

func (p *SfdPosition) SetSellOrder(orderAcceptanceID string) {
	sfdPosition.HaveOrder = &HavePosition{
		Buy:  false,
		Sell: true,
	}
	sfdPosition.AcceptanceID = orderAcceptanceID
}

func (p *SfdPosition) SetBuyOrder(orderAcceptanceID string) {
	sfdPosition.HaveOrder = &HavePosition{
		Buy:  true,
		Sell: false,
	}
	sfdPosition.AcceptanceID = orderAcceptanceID
}
