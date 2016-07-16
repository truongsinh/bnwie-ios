package facebook

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuth Facebook", func() {
	It("parses response error", func() {
		expectedBodyString := `{"error":{"message":"Error validating access token: Session has expired on Saturday, 02-May-15 01:07:02 PDT. The current time is Monday, 14-Sep-15 04:21:11 PDT.","type":"OAuthException","code":190,"error_subcode":463}}`
		body := bytes.NewBufferString(expectedBodyString)
		responseBody, bodyString, err := ParseError(body)
		Expect(responseBody.Message).To(Equal("Error validating access token: Session has expired on Saturday, 02-May-15 01:07:02 PDT. The current time is Monday, 14-Sep-15 04:21:11 PDT."))
		Expect(responseBody.Type).To(Equal("OAuthException"))
		Expect(responseBody.Code).To(Equal(190))
		Expect(responseBody.SubCode).To(Equal(463))
		Expect(responseBody.SubCode).To(Equal(463))
		Expect(responseBody.SubCode).To(Equal(463))
		Expect(bodyString).To(Equal(expectedBodyString))
		Expect(err).To(BeNil())
	})
})
