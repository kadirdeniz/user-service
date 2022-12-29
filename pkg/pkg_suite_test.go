package pkg

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"user-service/internal/user"
)

func TestPkg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkg Suite")
}

var userId, _ = primitive.ObjectIDFromHex("63add115ed3628749870398e")

var _ = Describe("Pkg", Ordered, func() {

	Context("JSONEncoder", func() {
		var userObj user.User

		decodedUserObj := []byte("{\"id\":\"63add115ed3628749870398e\",\"first_name\":\"John\",\"last_name\":\"Doe\",\"nickname\":\"johndoe\",\"email\":\"johndoe@mail.com\"}\n")

		It("should be success", func() {
			err := JSONEncoder(decodedUserObj, &userObj)
			Expect(err).Should(BeNil())
			Expect(userObj).ShouldNot(BeNil())
			Expect(userObj.ID).Should(Equal(userId))
		})
	})

	Context("JSONDecoder", func() {
		It("should be success", func() {
			decodedUserObj, err := JSONDecoder(&user.User{
				ID:        userId,
				FirstName: "John",
				LastName:  "Doe",
				Nickname:  "johndoe",
				Email:     "johndoe@mail.com",
				Password:  "123456",
			})
			Expect(err).Should(BeNil())
			Expect(decodedUserObj).ShouldNot(BeNil())
		})
	})

})
