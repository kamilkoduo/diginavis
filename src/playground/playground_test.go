package playground_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kamilkoduo/diginavis/src/playground"
)

var _ = Describe("Playground", func() {
	var (
		str1  string
		str12 string
	)
	BeforeEach(func() {
		str1 = "1"
		str12 = "12"
	})
	Describe("Calculating string entropy", func() {
		Context("with 1 symbol", func() {
			It("should be zero", func() {
				Expect(CalculateEntropy(str1)).Should(BeZero())
			})
		})

		Context("with 2 different symbols", func() {
			It("should be equal to 1", func() {
				Expect(CalculateEntropy(str12)).Should(BeEquivalentTo(1))
			})
		})
	})
})
