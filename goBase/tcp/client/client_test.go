package client

import "testing"

func TestCreateClient(t *testing.T) {
	t.Run("createClient", func(t *testing.T) {
		CreateClient(":8085", "第一条信息！")
	})
}
