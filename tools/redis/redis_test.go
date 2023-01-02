package redis_test

import (
	"testing"
	"time"
	"user-service/internal/user"
	"user-service/test/mock"
	"user-service/tools/dockertest"
	"user-service/tools/redis"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var redisConfig = mock.RedisConfig

func TestRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redis Suite")
}

var _ = Describe("Redis", Ordered, func() {

	var redisClient redis.RedisInterface
	var dockerContainer *dockertest.Dockertest

	BeforeAll(func() {
		dockerContainer = dockertest.NewDockertest("")
		err := dockerContainer.RunRedis(redisConfig)
		Expect(err).Should(BeNil())
	})

	AfterAll(func() {
		dockerContainer.Purge()
	})

	Context("NewRedis", func() {
		It("should return redis client", func() {
			redisClient = redis.NewRedis(redisConfig)
			Expect(redisClient).ShouldNot(BeNil())
		})
	})

	Context("Connect", func() {
		It("Should return database", func() {
			db, err := redisClient.Connect()
			Expect(err).Should(BeNil())
			Expect(db).ShouldNot(BeNil())
		})
	})

	Context("GetRedisClient", func() {
		It("should return redis client", func() {
			redisClient := redisClient.GetRedisClient()
			Expect(redisClient).ShouldNot(BeNil())
		})
	})

	Context("Flush", func() {
		It("should return redis client", func() {
			err := redisClient.Flush()
			Expect(err).Should(BeNil())
		})
	})

	When("User is not exist", func() {
		Context("GetUserByID", func() {
			It("shouldnt return user", func() {
				userObj, err := redisClient.GetUserByID(user.MockUser.ID)
				Expect(err).Should(BeNil())
				Expect(userObj).Should(Equal(new(user.User)))
			})
		})

		Context("SetUser", func() {
			It("should set user", func() {
				err := redisClient.SetUser(user.MockUser, time.Minute)
				Expect(err).Should(BeNil())
			})
		})
	})

	When("User is exist", func() {
		Context("GetUserByID", func() {
			It("should return user", func() {
				userObj, err := redisClient.GetUserByID(user.MockUser.ID)
				Expect(err).Should(BeNil())
				Expect(userObj).ShouldNot(Equal(user.MockUser))
			})
		})

		Context("SetUser", func() {
			It("should set user", func() {
				err := redisClient.SetUser(user.MockUser, 0)
				Expect(err).Should(BeNil())
			})
		})
	})

	Context("GetUserByID", func() {
		It("should return user", func() {
			user, err := redisClient.GetUserByID(user.MockUser.ID)
			Expect(err).Should(BeNil())
			Expect(user).ShouldNot(BeNil())
		})
	})

})
