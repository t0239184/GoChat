package crypto

import (
	"encoding/base64"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Test_Generate_Salt(t *testing.B) {
    length := 64
    salt := GenerateSalt(length)
    assert.Equal(t, length, len(salt))
}

func Benchmark_Test_Hash(t *testing.B) {
    hashedText := Hash("message")
    assert.Equal(t, "-Nr1ejNHzE1rnVdbMf5gd-LLSH9gqWIzwIy0edvzFTjMkV7G1IvbqpbdwaFttPT5bzcnbPyzUQuCRiQXcNWVLA==", hashedText)
}

func Benchmark_Test_Hash_With_Salt(t *testing.B) {
    salt, _ := base64.URLEncoding.DecodeString("UA1xIyzslG3yVmJint7snEK_je8Apu6EUL4WutOpFfjbVKzEmHwSHPKyixln2a3KnZi4ANm9v6qpsaJZYvm0ww==")
    hashedText := HashWithSalt("message", salt)
    // t.Log(hashedText, base64.URLEncoding.EncodeToString(salt))
    assert.Equal(t, "bjoQoYe9w8lA80HDXdCbzvocZKeWyOmFWLSQDnrE6DJNz-Xu0kDM6rxnTAj2JH1_bI0h4IqipJEh1D0oDMSEUQ==", hashedText)
}

func Benchmark_Test_Hash_With_Salt_And_Iteration(t *testing.B) {
    salt, _ := base64.URLEncoding.DecodeString("UA1xIyzslG3yVmJint7snEK_je8Apu6EUL4WutOpFfjbVKzEmHwSHPKyixln2a3KnZi4ANm9v6qpsaJZYvm0ww==")
    hashedText := HashWithSaltAndIteration("message", salt, 1000)
    // t.Log(hashedText, base64.URLEncoding.EncodeToString(salt))
    assert.Equal(t, "LcVxHQQHo4txBnN4ed3o_6inzvTBz5GnAQKV6FQNhj0gasjOtritpAfO3KkRu9IL9Xp9GTgSdPbqXz-SKlSYwA==", hashedText)
}

func Benchmark_Test_Hash_With_Salt_And_Iteration_using_two_goruntine(t *testing.B) {
    salt, _ := base64.URLEncoding.DecodeString("UA1xIyzslG3yVmJint7snEK_je8Apu6EUL4WutOpFfjbVKzEmHwSHPKyixln2a3KnZi4ANm9v6qpsaJZYvm0ww==")
    var hashedText string
    var wg = &sync.WaitGroup{}
    wg.Add(1)
    go func() {
        hashedText = HashWithSaltAndIteration("message", salt, 1000)
        wg.Done()
    }()
    wg.Wait()
    // t.Log(hashedText, base64.URLEncoding.EncodeToString(salt))
    assert.Equal(t, "LcVxHQQHo4txBnN4ed3o_6inzvTBz5GnAQKV6FQNhj0gasjOtritpAfO3KkRu9IL9Xp9GTgSdPbqXz-SKlSYwA==", hashedText)
}