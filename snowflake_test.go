package go-snowflake

import (
	"testing"
)

func TestSnowflake_Generate1(t *testing.T) {
	node := NewSnowflake(1, 1)
	ids := make(map[int64]bool)

	for i := 0; i < 1; i++ {
		id := node.Generate()
		if _, exists := ids[id]; exists {
			t.Fatalf("ID is not unique: %d", id)
		}
		ids[id] = true
	}
}

func TestSnowflake_Generate1000(t *testing.T) {
	node := NewSnowflake(1, 1)
	ids := make(map[int64]bool)

	for i := 0; i < 1000; i++ {
		id := node.Generate()
		if _, exists := ids[id]; exists {
			t.Fatalf("ID is not unique: %d", id)
		}
		ids[id] = true
	}
}

func TestSnowflake_GenerateOneMillion(t *testing.T) {
	node := NewSnowflake(1, 1)
	ids := make(map[int64]bool)

	for i := 0; i < 1000000; i++ {
		id := node.Generate()
		if _, exists := ids[id]; exists {
			t.Fatalf("ID is not unique: %d", id)
		}
		ids[id] = true
	}
}
