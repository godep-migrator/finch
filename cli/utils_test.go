package main

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"

	"os"
	"testing"
)

func TestGetTaskDBMem(t *testing.T) {
	os.Setenv("FINCH_STORAGE", "mem")

	_, err := getTaskDB()
	assert.Nil(t, err)
}

func TestGetTaskDBNone(t *testing.T) {
	os.Setenv("FINCH_STORAGE", "")
	folder, err := os.Getwd()
	assert.Nil(t, err)
	os.Setenv("HOME", folder)

	_, err = getTaskDB()
	assert.Nil(t, err)

	should := filepath.Join(folder, ".finchdb")
	_, err = os.Stat(should)
	if assert.False(t, os.IsNotExist(err)) {
		err := os.RemoveAll(should)
		assert.Nil(t, err)
	}
}

func TestGetTaskDBFile(t *testing.T) {
	folder, err := os.Getwd()
	assert.Nil(t, err)

	should := filepath.Join(folder, ".finchdb")
	os.Setenv("FINCH_STORAGE", should)

	_, err = getTaskDB()
	assert.Nil(t, err)

	_, err = os.Stat(should)
	if assert.False(t, os.IsNotExist(err)) {
		err := os.RemoveAll(should)
		assert.Nil(t, err)
	}
}
