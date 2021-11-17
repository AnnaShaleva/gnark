package circuits

import (
	"github.com/consensys/gnark/frontend"
)

type xorCircuit struct {
	Op1, Op2, Res frontend.Variable
}

func (circuit *xorCircuit) Define(cs frontend.API) error {
	d := cs.Xor(circuit.Op1, circuit.Op2)

	cs.AssertIsEqual(d, circuit.Res)
	return nil
}

func init() {

	good := []frontend.Circuit{
		&xorCircuit{
			Op1: (1),
			Op2: (1),
			Res: (0),
		},
		&xorCircuit{
			Op1: (1),
			Op2: (0),
			Res: (1),
		},
		&xorCircuit{
			Op1: (0),
			Op2: (1),
			Res: (1),
		},
		&xorCircuit{
			Op1: (0),
			Op2: (0),
			Res: (0),
		},
	}

	bad := []frontend.Circuit{
		&xorCircuit{
			Op1: (1),
			Op2: (1),
			Res: (1),
		},
		&xorCircuit{
			Op1: (1),
			Op2: (0),
			Res: (0),
		},
		&xorCircuit{
			Op1: (0),
			Op2: (1),
			Res: (0),
		},
		&xorCircuit{
			Op1: (0),
			Op2: (0),
			Res: (1),
		},
		&xorCircuit{
			Op1: (42),
			Op2: (1),
			Res: (1),
		},
		&xorCircuit{
			Op1: (1),
			Op2: (1),
			Res: (42),
		},
	}

	addNewEntry("xor", &xorCircuit{}, good, bad)
}
