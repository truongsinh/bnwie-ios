package model

import (
	"testing"

	"github.com/Smarp/socnet"
	_ "github.com/lib/pq"
	. "github.com/onsi/gomega"
)

func TestUpsertAndGetDesc(t *testing.T) {
	RegisterTestingT(t)
	prepareUserUpsert(prepareDB())
	u := &User{
		Profile: &socnet.Profile{
			Id:        "123",
			FullName:  "MyFullName",
			AvatarUrl: "MyAvatarUrl",
		},
		ExtendedData: &ExtendedData{
			Description: "MyNewDesc",
		},
	}
	err := UpsertUserModel(u)
	Expect(err).To(BeNil())
	Expect(u.Provider).To(Equal([]string{}))
	Expect(u.Consumer).To(Equal([]string{}))
	Expect(u.Description).To(Equal(""))
}
