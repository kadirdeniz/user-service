package user

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = Describe("User", Ordered, func() {
	Context("NewUser", func() {
		It("should be success", func() {
			user := New(MockUser.FirstName, MockUser.LastName, MockUser.Nickname, MockUser.Email, MockUser.Password)
			Expect(user.FirstName).Should(Equal(MockUser.FirstName))
			Expect(user.LastName).Should(Equal(MockUser.LastName))
			Expect(user.Nickname).Should(Equal(MockUser.Nickname))
			Expect(user.Email).Should(Equal(MockUser.Email))
			Expect(user.Password).Should(Equal(MockUser.Password))
			Expect(user.ID).ShouldNot(BeNil())
		})
	})

	Context("IsEmpty", func() {
		It("should be true", func() {
			userObj := new(User)
			Expect(userObj.IsEmpty()).Should(BeTrue())
		})

		It("should be false", func() {
			userObj := new(User)
			userObj.ID = MockUser.ID
			Expect(userObj.IsEmpty()).Should(BeFalse())
		})
	})
})
