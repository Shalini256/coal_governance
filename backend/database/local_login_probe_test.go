package database

import (
	"testing"

	"coal-governance-backend/config"
)

func TestLocalLoginProbe(t *testing.T) {
	if err := Connect(config.Load().DatabaseURL); err != nil {
		t.Fatal(err)
	}
	defer DB.Close()

	const seedHash = "$2b$10$V/goxVRUUnt78Fk7oRyUYeg/qRVgxhTR75.Flor9CKa6yMNB.VI5e"
	var adminExists, seededHashMatches bool
	var roleCount, userCount, mineCount int
	err := DB.QueryRow(`
		SELECT EXISTS (SELECT 1 FROM users WHERE email = ?),
		       COALESCE((SELECT password_hash = ? FROM users WHERE email = ?), false)`,
		"admin@coal.gov", seedHash, "admin@coal.gov").Scan(&adminExists, &seededHashMatches)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("admin_exists=%t seeded_hash_matches=%t", adminExists, seededHashMatches)
	if err := DB.QueryRow(`SELECT (SELECT COUNT(*) FROM roles), (SELECT COUNT(*) FROM users), (SELECT COUNT(*) FROM mines)`).Scan(&roleCount, &userCount, &mineCount); err != nil {
		t.Fatal(err)
	}
	t.Logf("roles=%d users=%d mines=%d", roleCount, userCount, mineCount)
}