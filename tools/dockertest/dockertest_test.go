package dockertest

import (
	"testing"
	"user-service/test/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var mongoConfig = mock.MongoConfig
var redisConfig = mock.RedisConfig

func TestDockertest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dockertest Suite")
}

var _ = Describe("Dockertest", Ordered, func() {

	Context("NewDockertest", func() {
		It("should be success", func() {
			dockertest := NewDockertest("")
			Expect(dockertest).ShouldNot(BeNil())
		})
	})

	Context("RunMongoDB", func() {
		It("should be success", func() {
			dockertest := NewDockertest("")
			err := dockertest.RunMongoDB(mongoConfig)
			Expect(err).Should(BeNil())
		})
	})

	Context("RunRedis", func() {
		It("should be success", func() {
			dockertest := NewDockertest("")
			err := dockertest.RunRedis(redisConfig)
			Expect(err).Should(BeNil())
		})
	})

	Context("Purge", func() {
		It("should be purge redis", func() {
			dockertest := NewDockertest("")
			dockertest.RunRedis(redisConfig)
			err := dockertest.Purge()
			Expect(err).Should(BeNil())
		})

		It("should be purge mongo", func() {
			dockertest := NewDockertest("")
			dockertest.RunMongoDB(mongoConfig)
			err := dockertest.Purge()
			Expect(err).Should(BeNil())
		})
	})
})
