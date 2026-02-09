package foundation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeIdGenerator(t *testing.T) {
	t.Run("generates sequential IDs with prefix", func(t *testing.T) {
		gen := MakeIdGenerator("test")
		id1 := gen()
		id2 := gen()
		id3 := gen()

		assert.Equal(t, "test-1", id1)
		assert.Equal(t, "test-2", id2)
		assert.Equal(t, "test-3", id3)
	})

	t.Run("multiple generators are independent", func(t *testing.T) {
		gen1 := MakeIdGenerator("gen1")
		gen2 := MakeIdGenerator("gen2")

		id1 := gen1()
		id2 := gen2()
		id3 := gen1()

		assert.Equal(t, "gen1-1", id1)
		assert.Equal(t, "gen2-1", id2)
		assert.Equal(t, "gen1-2", id3)
	})

	t.Run("empty prefix", func(t *testing.T) {
		gen := MakeIdGenerator("")
		id := gen()
		assert.Equal(t, "-1", id)
	})

	t.Run("generates many IDs", func(t *testing.T) {
		gen := MakeIdGenerator("many")
		ids := make(map[string]bool)

		for i := 0; i < 100; i++ {
			id := gen()
			// Verify uniqueness
			assert.False(t, ids[id], "ID %s was generated twice", id)
			ids[id] = true
		}

		assert.Len(t, ids, 100)
	})

	t.Run("prefix with special characters", func(t *testing.T) {
		gen := MakeIdGenerator("test-id_v1")
		id := gen()
		assert.Equal(t, "test-id_v1-1", id)
	})
}

func TestMakeIntIdGenerator(t *testing.T) {
	t.Run("generates sequential uint64 IDs", func(t *testing.T) {
		gen := MakeIntIdGenerator()
		id1 := gen()
		id2 := gen()
		id3 := gen()

		assert.Equal(t, uint64(1), id1)
		assert.Equal(t, uint64(2), id2)
		assert.Equal(t, uint64(3), id3)
	})

	t.Run("multiple generators are independent", func(t *testing.T) {
		gen1 := MakeIntIdGenerator()
		gen2 := MakeIntIdGenerator()

		id1 := gen1()
		id2 := gen2()
		id3 := gen1()

		assert.Equal(t, uint64(1), id1)
		assert.Equal(t, uint64(1), id2)
		assert.Equal(t, uint64(2), id3)
	})

	t.Run("generates many IDs", func(t *testing.T) {
		gen := MakeIntIdGenerator()
		ids := make(map[uint64]bool)

		for i := 0; i < 1000; i++ {
			id := gen()
			// Verify uniqueness
			assert.False(t, ids[id], "ID %d was generated twice", id)
			ids[id] = true
		}

		assert.Len(t, ids, 1000)
	})

	t.Run("starts from 1 not 0", func(t *testing.T) {
		gen := MakeIntIdGenerator()
		id := gen()
		assert.Equal(t, uint64(1), id)
	})

	t.Run("sequential ordering", func(t *testing.T) {
		gen := MakeIntIdGenerator()
		var prev uint64 = 0

		for i := 0; i < 100; i++ {
			id := gen()
			assert.Greater(t, id, prev)
			prev = id
		}
	})
}

func TestNewUUID(t *testing.T) {
	t.Run("generates valid UUID", func(t *testing.T) {
		uuid := NewUUID()
		assert.NotEmpty(t, uuid)
		// UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (36 characters)
		assert.Len(t, uuid, 36)
		assert.Contains(t, uuid, "-")
	})

	t.Run("generates unique UUIDs", func(t *testing.T) {
		uuids := make(map[string]bool)

		for i := 0; i < 100; i++ {
			uuid := NewUUID()
			assert.False(t, uuids[uuid], "UUID %s was generated twice", uuid)
			uuids[uuid] = true
		}

		assert.Len(t, uuids, 100)
	})

	t.Run("multiple calls return different UUIDs", func(t *testing.T) {
		uuid1 := NewUUID()
		uuid2 := NewUUID()
		uuid3 := NewUUID()

		assert.NotEqual(t, uuid1, uuid2)
		assert.NotEqual(t, uuid2, uuid3)
		assert.NotEqual(t, uuid1, uuid3)
	})
}
