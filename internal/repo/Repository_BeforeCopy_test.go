package repo

import (
	"os"
	"testing"

	"github.com/origadmin/toolkits/context"
)

// Successfully creates a cache directory if it does not exist
func TestCreatesCacheDirectory(t *testing.T) {
	// Arrr, let's set sail and test the creation of the cache directory!
	repo := Repository{
		cacheDir: "test_cache_dir",
		useCmd:   false,
	}
	err := repo.BeforeCopy(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if _, err := os.Stat("test_cache_dir"); os.IsNotExist(err) {
		t.Fatalf("Expected cache directory to be created, but it does not exist")
	}
	os.RemoveAll("test_cache_dir")
}

// Clones the repository using the go-git library when useCmd is false
func TestCloneRepoWithGoGit(t *testing.T) {
	// Avast ye! Testing the go-git library for cloning!
	repo := Repository{
		repoURL:  "https://github.com/goexts/ggb",
		cacheDir: "test_cache_dir",
		useCmd:   false,
	}
	err := repo.BeforeCopy(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
}

// Clones the repository using system command when useCmd is true
func TestCloneRepoWithCmd(t *testing.T) {
	// Shiver me timbers! Testing the system command for cloning!
	repo := Repository{
		repoURL:  "https://github.com/goexts/ggb",
		cacheDir: "test_cache_dir",
		useCmd:   true,
	}
	err := repo.BeforeCopy(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

}

// Handles error when cache directory creation fails
func TestErrorOnCacheDirCreation(t *testing.T) {
	// Arrr, testing the stormy seas of directory creation failure!
	repo := Repository{
		cacheDir: "/invalid/path/to/cache",
		useCmd:   false,
	}
	err := repo.BeforeCopy(context.Background())
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}

// Handles error when cloning fails using go-git
func TestErrorOnCloneWithGoGit(t *testing.T) {
	// Ahoy! Testing the treacherous waters of go-git cloning failure!
	repo := Repository{
		repoURL:  "invalid-url",
		cacheDir: "test_cache_dir",
		useCmd:   false,
	}
	err := repo.BeforeCopy(context.Background())
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}

// Handles error when cloning fails using system command
func TestErrorOnCloneWithCmd(t *testing.T) {
	// Yo ho ho! Testing the perilous voyage of system command cloning failure!
	repo := Repository{
		repoURL:  "invalid-url",
		cacheDir: "test_cache_dir",
		useCmd:   true,
	}
	err := repo.BeforeCopy(context.Background())
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}

// Handles error when checking out a tag fails
func TestErrorOnTagCheckout(t *testing.T) {
	// Arr matey! Testing the rough seas of tag checkout failure!
	repo := Repository{
		repoURL:  "https://github.com/goexts/ggb",
		cacheDir: "test_cache_dir",
		tag:      "non-existent-tag",
		useCmd:   false,
	}
	err := repo.BeforeCopy(context.Background())
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}
