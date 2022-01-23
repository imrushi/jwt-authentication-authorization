package handler

import "testing"

func TestInitalMigration(t *testing.T) {
	connection := GetDatabase()
	t.Log(connection)
}
