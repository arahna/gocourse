package main

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		first, second string
		want          bool
	}{
		{"", "", false},
		{"Апельсин", "Спаниель", true},
		{"Апельсина", "Спаниель", false},
	}
	for _, c := range cases {
		got := isAnagram1(c.first, c.second)
		if got != c.want {
			t.Errorf("isAnagram1(\"%v\", \"%v\") == %v, want %v\n", c.first, c.second, got, c.want)
		}
	}
	for _, c := range cases {
		got := isAnagram2(c.first, c.second)
		if got != c.want {
			t.Errorf("isAnagram2(\"%v\", \"%v\") == %v, want %v\n", c.first, c.second, got, c.want)
		}
	}
}

func BenchmarkAnagram1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isAnagram1("Апельсин", "Спаниель")
	}
}

func BenchmarkAnagram2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isAnagram2("Апельсин", "Спаниель")
	}
}
