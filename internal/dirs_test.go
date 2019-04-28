package internal

import (
	"testing"
	"io/ioutil"
	"os"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListFilesFromDirectory(t *testing.T) {
	// create tmp directory
	dir, err := ioutil.TempDir(".", uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create tmp files
	_, err = ioutil.TempFile(dir, uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.TempFile(dir, uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}

	files, err := GetFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, files)
	assert.Equal(t, len(files), 2)
}

func TestListFilesFromEmptyDirectory(t *testing.T) {
	// create tmp directory
	dir, err := ioutil.TempDir(".", uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	files, err := GetFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	assert.Empty(t, files)
}

func TestListSingleFile(t *testing.T) {
	// create tmp directory
	dir, err := ioutil.TempDir(".", uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// create tmp file
	file, err := ioutil.TempFile(dir, uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}

	files, err := GetFiles(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, files)
	assert.Equal(t, len(files), 1)
}

