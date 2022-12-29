package redis_test

import (
	"testing"
	"time"
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

	Context("SetUser", func() {
		It("should set user", func() {
			err := redisClient.SetUser(mock.MockUser, time.Second)
			Expect(err).Should(BeNil())
		})
	})

	Context("GetUserByID", func() {
		It("should return user", func() {
			user, err := redisClient.GetUserByID(mock.MockUser.ID)
			Expect(err).Should(BeNil())
			Expect(user).ShouldNot(BeNil())
		})
	})
})
