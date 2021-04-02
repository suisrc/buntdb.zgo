package buntdb_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suisrc/buntdb.zgo"
)

func TestStore(t *testing.T) {
	store, err := buntdb.NewStore(":memory:")
	assert.Nil(t, err)

	defer store.Close()

	key := "test"
	ctx := context.Background()
	err = store.Set1(ctx, key, 0)
	assert.Nil(t, err)

	b, err := store.Check(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, true, b)

	err = store.Delete(ctx, key)
	assert.Nil(t, err)
}
