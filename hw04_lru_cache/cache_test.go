package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(5)

		list := []struct {
			Key   string
			Value int
		}{
			{Key: "10", Value: 10},
			{Key: "20", Value: 20},
			{Key: "30", Value: 30},
			{Key: "40", Value: 40},
			{Key: "50", Value: 50},
		}

		for _, v := range list {
			require.False(t, c.Set(Key(v.Key), v.Value))
		}

		c.Clear()

		for _, v := range list {
			value, ok := c.Get(Key(v.Key))
			require.False(t, ok)
			require.Nil(t, value)
		}
	})

	t.Run("capacity", func(t *testing.T) {
		type KV struct {
			Key   string
			Value int
		}

		c := NewCache(3)

		requestList := []KV{
			{Key: "10", Value: 10},
			{Key: "20", Value: 20},
			{Key: "30", Value: 30},
			{Key: "40", Value: 40},
			{Key: "50", Value: 50},
		}

		responseList := map[string]bool{
			"10": false,
			"20": false,
			"30": true,
			"40": true,
			"50": true,
		}

		for _, v := range requestList {
			c.Set(Key(v.Key), v.Value)
		}

		for i, v := range responseList {
			_, ok := c.Get(Key(i))
			require.Equal(t, v, ok)
		}
	})

	t.Run("capacity used element", func(t *testing.T) {
		t.Skip() // Remove me if you wrote how to check it
		type KV struct {
			Key   string
			Value int
		}

		c := NewCache(3)

		requestList := []KV{
			{Key: "10", Value: 10},
			{Key: "20", Value: 20},
			{Key: "30", Value: 30},
			{Key: "10", Value: 11},
			{Key: "30", Value: 31},
			{Key: "10", Value: 12},
			{Key: "20", Value: 21},
			{Key: "30", Value: 32},
			{Key: "20", Value: 22},
			{Key: "40", Value: 40},
		}

		responseList := []KV{
			{Key: "40", Value: 40},
			{Key: "20", Value: 22},
			{Key: "30", Value: 32},
		}

		for _, v := range requestList {
			c.Set(Key(v.Key), v.Value)
		}

		for _, v := range responseList {
			_ = v
			// how to check it
		}

	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
