package playground

import (
	. "github.com/onsi/gomega"
	"math"
	"testing"
	"time"
)

func TestCalculateEntropy(t *testing.T) {
	ent, err := CalculateEntropy("12345")
	delta := math.Abs(float64(math.Round(1e3*ent)-2322) / 1e3)

	if err != nil || delta >= 1e-3 {
		t.Error("Expected H(\"12345\") = 2.322 ")
	}
}

func TestCalculateEntropyGomega(t *testing.T) {
	ent, err := CalculateEntropy("12345")
	delta := math.Abs(float64(math.Round(1e3*ent)-2322) / 1e3)

	g := NewGomegaWithT(t)

	annotation := "Expected H(\"12345\") = 2.322 with error < 0.001"

	//g.Expect(err).ShouldNot(HaveOccurred())
	//g.Expect(delta < 1e-3).To(BeTrue(), annotation)

	g.Expect(delta, err).Should(BeNumerically("<", 1e-3), annotation)

}

func TestPrintNonEmptyStr(t *testing.T) {
	g := NewGomegaWithT(t)

	annotation := "Expected error on empty string"
	g.Expect(PrintNonEmptyStr("something\n")).Should(Succeed(), annotation)
}

func TestPrintEmptyStr(t *testing.T) {
	g := NewGomegaWithT(t)

	annotation := "Expected error on empty string"
	g.Expect(PrintNonEmptyStr("")).ShouldNot(Succeed(), annotation)
}

func TestPrintNonEmptyStrEventually(t *testing.T) {
	g := NewGomegaWithT(t)
	annotation := "Eventually has not secceeded"
	var r string
	go func() {
		r, _ = ReturnStringAfter2s("test")
	}()
	g.Eventually(func() string {
		return r
	}, "2.1s", "0.1s").Should(Equal("test"), annotation)
}

func TestSummationConsistently(t *testing.T) {
	g := NewGomegaWithT(t)
	annotation := "Consistently has not secceeded"
	sum := 0
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Duration(1 * 1e9))
			sum++
		}
	}()
	g.Consistently(func() int {
		return sum
	}, "2s", "0.1s").Should(BeNumerically("<=", 3), annotation)
}

