package gocache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/lwhile/gocache"
)

func genRandomSeq(size int) []int {
	seq := make([]int, size)
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < size; i++ {
		seq[i] = rd.Intn(9999999)
	}
	return seq
}

func TestNewCache(t *testing.T) {
	max := 999999
	cache := gocache.NewCache(max)
	seq := genRandomSeq(max)
	for i := 0; i < max; i++ {
		cache.Set(i, seq[i])
	}

	if cache.Size() != max {
		t.Fatalf("cache size %d != %d\n", cache.Size(), max)
	}

	for i := 0; i < max; i++ {
		v, err := cache.Get(i)
		if err != nil {
			t.Fatal(err)
		}
		if v.(int) != seq[i] {
			es := fmt.Sprintf("%v != %d\n", v.(int), seq[i])
			t.Fatal(es)
		}
	}
}

func TestCache_Del(t *testing.T) {
	// insert something
	max := 10
	cache := gocache.NewCache(max)

	for i := 0; i < max; i++ {
		cache.Set(i, i)
	}

	for i := 0; i < max; i++ {
		cache.Del(i)
		v, err := cache.Get(i)
		if err == nil {
			t.Fatalf("delete key %v fail", v)
		}

		if cache.Size() != max-i-1 {
			t.Fatalf("%d != %d\n", cache.Size(), max)
		}
	}

	if cache.Size() != 0 {
		t.Fatalf("%d !=0\n", cache.Size())
	}

}
