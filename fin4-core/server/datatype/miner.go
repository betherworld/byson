package datatype

import "fin4-core/server/decimaldt"

// Miner user type
type Miner struct {
	ID               ID
	UserName         string
	Mined            decimaldt.Decimal
	MiningPercentage string
}
