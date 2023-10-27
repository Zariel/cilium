// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tableTest struct {
	vecs         [][]string
	expectedSize int
	expectedVecs [][]string
}

func TestIteratorProduct(t *testing.T) {
	assert := assert.New(t)
	p := CartesianProduct(
		[]string{"CiliumNetworkPolicy", "CiliumClusterwideNetworkPolicy", "NetworkPolicyName"},
		[]string{"update", "delete"},
		[]string{"true", "false"},
		[]string{"0", "1"},
	)
	count := 0
	p.ForEach(func(v []string) {
		count++
	})
	assert.Equal(24, count)
}

func TestIteratorProductEmpty(t *testing.T) {
	assert := assert.New(t)
	p := CartesianProduct(
		[]string{"CiliumNetworkPolicy", "CiliumClusterwideNetworkPolicy", "NetworkPolicyName"},
		[]string{},
	)
	assert.Equal(0, p.Size())

	p = CartesianProduct[string]()
	assert.Equal(0, p.Size())
}
