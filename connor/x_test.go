package connor

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideOrders(t *testing.T) {
	c := &Connor{}

	f := NewCorderFactory("ETH")

	ex1 := f.FromParams(big.NewInt(1), 100, *newBenchmarkFromMap(map[string]uint64{}))
	ex2 := f.FromParams(big.NewInt(2), 200, *newBenchmarkFromMap(map[string]uint64{}))
	ex3 := f.FromParams(big.NewInt(3), 300, *newBenchmarkFromMap(map[string]uint64{}))

	req0 := f.FromParams(big.NewInt(1), 50, *newBenchmarkFromMap(map[string]uint64{}))
	req1 := f.FromParams(big.NewInt(1), 100, *newBenchmarkFromMap(map[string]uint64{}))
	req2 := f.FromParams(big.NewInt(2), 200, *newBenchmarkFromMap(map[string]uint64{}))
	req3 := f.FromParams(big.NewInt(3), 300, *newBenchmarkFromMap(map[string]uint64{}))
	req4 := f.FromParams(big.NewInt(4), 400, *newBenchmarkFromMap(map[string]uint64{}))

	existing := []*Corder{ex1, ex2, ex3}
	required := []*Corder{req0, req1, req2, req3, req4}

	set := c.divideOrdersSets(existing, required)
	assert.Len(t, set.toRestore, 3)
	assert.Len(t, set.toCreate, 2)
}