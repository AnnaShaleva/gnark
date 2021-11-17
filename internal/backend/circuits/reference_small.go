package circuits

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
)

const nbConstraintsRefSmall = 5

type referenceSmallCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *referenceSmallCircuit) Define(cs frontend.API) error {
	for i := 0; i < nbConstraintsRefSmall; i++ {
		circuit.X = cs.Mul(circuit.X, circuit.X)
	}
	cs.AssertIsEqual(circuit.X, circuit.Y)
	return nil
}

func init() {
	var circuit, good, bad referenceSmallCircuit

	good.X = (2)

	// compute expected Y
	var expectedY big.Int
	expectedY.SetUint64(2)

	for i := 0; i < nbConstraintsRefSmall; i++ {
		expectedY.Mul(&expectedY, &expectedY)
	}

	good.Y = (expectedY)

	bad.X = (3)
	bad.Y = (expectedY)

	addEntry("reference_small", &circuit, &good, &bad)
}
