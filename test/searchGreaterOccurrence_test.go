package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fbonhomm/Knuth-Morris-Pratt-Algorithm/kmp"
)

var _ = Describe("SearchGreaterOccurrence", func() {
	Context("Success", func() {
		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abcdabdr\" as pattern", func() {
			// Prepare
			var index, length int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abcdabdr")

			// Process
			index, length = kmp.SearchGreaterOccurrence(buffer, pattern)

			// Expect
			Expect(length).To(Equal(7))
			Expect(index).To(Equal(15))
		})
	})

	Context("Success", func() {
		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abr\" as pattern", func() {
			// Prepare
			var index, length int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abr")

			// Process
			index, length = kmp.SearchGreaterOccurrence(buffer, pattern)

			// Expect
			Expect(length).To(Equal(2))
			Expect(index).To(Equal(0))
		})
	})

	Context("Success", func() {
		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abc\" as pattern", func() {
			// Prepare
			var index, length int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abc")

			// Process
			index, length = kmp.SearchGreaterOccurrence(buffer, pattern)

			// Expect
			Expect(length).To(Equal(3))
			Expect(index).To(Equal(0))
		})
	})
})