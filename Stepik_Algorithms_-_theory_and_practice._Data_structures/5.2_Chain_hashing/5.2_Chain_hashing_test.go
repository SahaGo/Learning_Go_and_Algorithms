package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

// тест хеш-функции со значеениями из комментариев
func TestHashFunc(t *testing.T) {
	assert := assert2.New(t)
	var m uint64 = 5
	assert.Equal("4", HashFunc("world", m), "should be equal")
	assert.Equal("4", HashFunc("HellO", m), "should be equal")
	assert.Equal("2", HashFunc("luck", m), "should be equal")
	assert.Equal("2", HashFunc("GooD", m), "should be equal")

	m = 17
	assert.Equal("12", HashFunc("a", m), "should be equal")
	assert.Equal("14", HashFunc("aaaaaaa", m), "should be equal")
	assert.Equal("10", HashFunc("aaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer wea", m), "should be eq")
	assert.Equal("11", HashFunc("aaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer wea", m), "should be eq")
	assert.Equal("16", HashFunc("aaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer wea", m), "should be eq")
	assert.Equal("8", HashFunc("aaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer wea", m), "should be eq")
	assert.Equal("0", HashFunc("aaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer weaaaaaaaaaaaaaaa sdd sre wew re wr swedrwe werwer wea", m), "should be eq")
}
