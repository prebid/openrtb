package openrtb2_test

import (
	"encoding/json"

	. "github.com/prebid/openrtb/v20/openrtb2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These fields (Regs.COPPA, Banner.TopFrame, Video.MinBitRate) were previously typed as
// plain int8/int64 with `omitempty`. Because encoding/json's omitempty treats a value type's
// zero value as empty, explicitly setting one of these fields to a meaningful 0 (COPPA: not
// subject to COPPA, TopFrame: not in the top frame, MinBitRate: a real 0 Kbps floor) was
// indistinguishable, once marshaled, from never having set the field at all - the sender's
// explicit signal was silently dropped. See https://github.com/prebid/openrtb/issues/13.
//
// Changing them to pointer types (matching the existing pattern used by Regs.GDPR,
// Banner.Vcm, and Device.DNT for the same "0/1, omission means Unknown" semantics) lets a
// caller distinguish "not set" (nil) from "explicitly set to 0" (Int8Ptr(0)/Int64Ptr(0)).
var _ = Describe("Zero-valued optional fields survive marshaling", func() {
	It("keeps Regs.COPPA when explicitly set to 0", func() {
		b, err := json.Marshal(Regs{COPPA: Int8Ptr(0)})
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(MatchJSON(`{"coppa":0}`))
	})

	It("omits Regs.COPPA when left unset", func() {
		b, err := json.Marshal(Regs{})
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(MatchJSON(`{}`))
	})

	It("keeps Banner.TopFrame when explicitly set to 0", func() {
		b, err := json.Marshal(Banner{TopFrame: Int8Ptr(0)})
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(MatchJSON(`{"topframe":0}`))
	})

	It("keeps Video.MinBitRate when explicitly set to 0", func() {
		b, err := json.Marshal(Video{MinBitRate: Int64Ptr(0)})
		Expect(err).NotTo(HaveOccurred())
		Expect(b).To(MatchJSON(`{"mimes":null,"minbitrate":0}`))
	})
})
