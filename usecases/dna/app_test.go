package dna

import "testing"

func TestDNAApplication(t *testing.T) {
	t.Run("Start Server", func(t *testing.T) {
		InitServer()
	})

	t.Run("Start User", func(t *testing.T) {
		InitUsers()
	})

	t.Run("Start Analyst", func(t *testing.T) {
		InitAnalyst()
	})
}

func TestOpenDataset(t *testing.T) {
	t.Run("open dataset", func(t *testing.T) {
		OpenDataset()
	})
}

func TestInitServer(t *testing.T) {
	t.Run("Start Server", func(t *testing.T) {
		InitServer()
	})
}

func TestInitUsers(t *testing.T) {
	t.Run("Start User", func(t *testing.T) {
		InitUsers()
	})
}

func TestInitAnalyst(t *testing.T) {
	t.Run("Start Analyst", func(t *testing.T) {
		InitAnalyst()
	})
}
