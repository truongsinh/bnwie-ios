package model

import (
	"testing"

	_ "github.com/lib/pq"
	. "github.com/onsi/gomega"
)

func TestUpdateService(t *testing.T) {
	RegisterTestingT(t)
	prepareUpdateSerivceUpsert(prepareDB())
	u := &ExtendedData{
		UserId:   11,
		Provider: []string{},
		Consumer: []string{},
	}
	err := UpdateServiceByIdModel(u)
	Expect(err).To(BeNil())
}
