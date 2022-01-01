package kubeflake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachineID(t *testing.T) {
	id, err := machineID()
	assert.NotZero(t, id)
	assert.Nil(t, err)
}

func TestNew(t *testing.T) {
	id, err := New()
	assert.NotZero(t, id)
	assert.Nil(t, err)
}

func TestMust(t *testing.T) {
	id := Must(New())
	assert.NotZero(t, id)
}

func TestCompareOrder(t *testing.T) {
	id1 := Must(New())
	id2 := Must(New())
	assert.Less(t, id1, id2)
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, _ = New()
	}
}

func BenchmarkMust(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = Must(New())
	}
}
