package internal

import (
	"fmt"
	"testing"
	"user-service/internal/user"
	"user-service/pkg"
	"user-service/test/mock"
	"user-service/tools/dockertest"
	"user-service/tools/mongodb"
	"user-service/tools/redis"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var mongoConfig = pkg.MongoDBConfig{
	Username: "admin",
	Password: "admin",
	Host:     "localhost",
	Port:     "27017",
	Database: "test",
}

var redisConfig = pkg.RedisConfig{
	Host: "localhost",
	Port: "6379",
	//Password: "",
	DB: 0,
}

func TestUserRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Repository Suite")
}

var _ = Describe("User Repository", Ordered, func() {

	var mongo mongodb.MongoDBInterface
	var redisClient redis.RedisInterface
	var mongoContainer *dockertest.Dockertest
	var redisContainer *dockertest.Dockertest

	var repo IRepository

	BeforeAll(func() {
		mongoContainer = dockertest.NewDockertest("")
		err := mongoContainer.RunMongoDB(mongoConfig)
		Expect(err).Should(BeNil())

		redisContainer = dockertest.NewDockertest("")
		err = redisContainer.RunRedis(redisConfig)
		Expect(err).Should(BeNil())

		mongo = mongodb.NewMongoDB(mongoConfig)
		redisClient = redis.NewRedis(redisConfig)
	})

	AfterAll(func() {
		mongoContainer.Purge()
		redisContainer.Purge()
	})

	Context("NewUserRepository", func() {
		It("should return user repository", func() {
			repo = NewRepository(mongo, redisClient)
			Expect(mongo).ShouldNot(BeNil())
			Expect(redisClient).ShouldNot(BeNil())
		})
	})

	When("User collection is empty", func() {

		Context("IsEmailExists", func() {
			It("should return false", func() {
				isEmailExists, err := repo.IsEmailExists(mock.MockUser.Email)
				Expect(err).Should(BeNil())
				Expect(isEmailExists).Should(BeFalse())
			})
		})

		Context("IsNicknameExists", func() {
			It("should return false", func() {
				isNicknameExists, err := repo.IsNicknameExists(mock.MockUser.Nickname)
				Expect(err).Should(BeNil())
				Expect(isNicknameExists).Should(BeFalse())
			})
		})

		Context("GetUserByID", func() {
			It("shouldnt return user", func() {
				userObj, err := repo.GetUserByID(mock.MockUser.ID)
				fmt.Println(err)
				Expect(err).ShouldNot(BeNil())
				Expect(err).Should(Equal(pkg.ErrUserNotFound))
				Expect(userObj).Should(Equal(new(user.User)))
			})
		})

		Context("GetUsers", func() {
			It("should return empty slice", func() {
				users, err := repo.GetUsers()
				Expect(err).ShouldNot(BeNil())
				Expect(err).Should(Equal(pkg.ErrUserNotFound))
				Expect(users).Should(BeNil())
			})
		})

		Context("Upsert", func() {
			It("should insert user", func() {
				user := mock.MockUser
				err := repo.Upsert(user)
				Expect(err).Should(BeNil())
			})
		})
	})

	When("User collection is not empty", func() {

		Context("IsEmailExists", func() {
			It("should return true", func() {
				isEmailExists, err := repo.IsEmailExists(mock.MockUser.Email)
				Expect(err).Should(BeNil())
				Expect(isEmailExists).Should(BeTrue())
			})
		})

		Context("IsNicknameExists", func() {
			It("should return true", func() {
				isNicknameExists, err := repo.IsNicknameExists(mock.MockUser.Nickname)
				Expect(err).Should(BeNil())
				Expect(isNicknameExists).Should(BeTrue())
			})
		})

		Context("GetUserByID", func() {
			It("should return user", func() {
				userObj, err := repo.GetUserByID(mock.MockUser.ID)
				Expect(err).Should(BeNil())
				Expect(userObj).ShouldNot(BeNil())
				Expect(userObj.ID).Should(Equal(mock.MockUser.ID))
				Expect(userObj.FirstName).Should(Equal(mock.MockUser.FirstName))
				Expect(userObj.LastName).Should(Equal(mock.MockUser.LastName))
				Expect(userObj.Email).Should(Equal(mock.MockUser.Email))
				Expect(userObj.Nickname).Should(Equal(mock.MockUser.Nickname))
			})
		})

		Context("GetUsers", func() {
			It("should return slice of users", func() {
				users, err := repo.GetUsers()
				Expect(err).Should(BeNil())
				Expect(users).ShouldNot(BeNil())
				Expect(users).Should(Equal([]*user.User{mock.MockUser}))
			})
		})

		Context("DeleteUserByID", func() {
			It("should delete user", func() {
				err := repo.DeleteUserByID(mock.MockUser.ID)
				Expect(err).Should(BeNil())
			})
		})
	})
})
