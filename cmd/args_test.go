package cmd

import "testing"

func TestCheckAllNotEmpty(t *testing.T) {
	CheckAllNotEmpty("", "123")
}
