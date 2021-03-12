package application

type (
	SfdPosition struct {
		HavePosition *HavePosition
		//PositionInfo
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

func (p *SfdPosition) SetSellPosition() {
	sfdPosition.HavePosition = &HavePosition{
		Buy:  false,
		Sell: true,
	}
}
