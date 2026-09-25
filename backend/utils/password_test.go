package utils

import "testing"

func TestSeedDemoPassword(t *testing.T) {
	const seedHash = "$2b$10$V/goxVRUUnt78Fk7oRyUYeg/qRVgxhTR75.Flor9CKa6yMNB.VI5e"
	if !CheckPassword(seedHash, "Coal@2026") {
		t.Fatal("seed hash does not match the documented demo password")
	}
}
