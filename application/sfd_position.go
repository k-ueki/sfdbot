package application

type (
	SfdPosition struct{
		HavePosition *HavePosition
	//PositionInfo
}

	HavePosition struct{
		Buy bool
		Sell bool
	}
)

var sfdPosition SfdPosition
