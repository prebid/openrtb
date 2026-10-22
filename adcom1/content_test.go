package adcom1_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/prebid/openrtb/v20/adcom1"
)

var _ = DescribeTable(
	"Content broadcast flags preserve omitted and explicit values",
	func(payload string) {
		var content Content
		Expect(json.Unmarshal([]byte(payload), &content)).To(Succeed())

		actual, err := json.Marshal(content)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(MatchJSON(payload))
	},
	Entry("omitted", `{}`),
	Entry("realtime zero", `{"realtime":0}`),
	Entry("realtime one", `{"realtime":1}`),
	Entry("firstbroadcast zero", `{"firstbroadcast":0}`),
	Entry("firstbroadcast one", `{"firstbroadcast":1}`),
)
