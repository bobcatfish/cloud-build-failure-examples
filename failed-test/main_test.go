package main

import (
	"fmt"
    	"testing"
)


func TestCountTheCats(t *testing.T) {
	for _, c := range []struct {
		cats     []cat
		want	int
	}{{
		cats: []cat{{
			name: "sparkles",
		}},
		want: 1,
	}, {
		cats: []cat{{
			name: "sparkles",
		},{
			name: "colonel mustard",
		}},
		want: 2,
	}, {
		cats: []cat{{
			name: "sparkles",
		},{
			name: "colonel mustard",
		},{
			name: "mcfluffles",
		}},
		want: 3,
	}} {
		t.Run(fmt.Sprintf("%d cats", c.want), func(t *testing.T) {
			got := countTheCats(c.cats)
			if got != c.want {
				t.Errorf("Wanted %d cats got %d", c.want, got)
			}
		})
	}
}
