package constants

type UnitType string
type Layer string

var TroopTypes = []UnitType{
	Infantry,
	Tank,
	Fort,
	Plane,
}

const (
	// troop type names
	Infantry UnitType = "Infantry"
	Tank     UnitType = "Tank"
	Fort     UnitType = "Fort"
	Plane    UnitType = "Plane"
)

const (
	// capital names
	CapitalLv1 UnitType = "CapitalLv1"
	CapitalLv2 UnitType = "CapitalLv2"
	CapitalLv3 UnitType = "CapitalLv3"
	CapitalLv4 UnitType = "CapitalLv4"
)

const (
	// building types
	ResearchCenter UnitType = "ResearchCenter"
	TradingCenter  UnitType = "TradingCenter"
)

const (
	// layers
	Underground Layer = "Underground"
	Ground      Layer = "Ground"
	Air         Layer = "Air"
)
