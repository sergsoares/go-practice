package main

import "testing"

func Test_Sum(t *testing.T) {
	tables := []struct {
		x    int
		y    int
		want int
	}{
		{1, 2, 3},
		{10, 100, 110},
		{23, 10, 33},
		{8, 12, 20},
	}

	for _, item := range tables {
		got := sum(values{item.x, item.y})
		t.Logf("Teste %d, %d", item.x, item.y)

		if item.want != got {
			t.Errorf("Error %d, %d", item.want, got)
		}
	}
}
