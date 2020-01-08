package salut

import "testing"

func TestSalut(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "Bonjour Monsieur!"},
		{0, "Bonjour Madame!"},
		{999, "Bonjour!"},
		{-100, "Bonjour!"},
	}

	for i, c := range cases {
		out := Salut(c.in)
		if out != c.want {
			t.Errorf("Case[%d]: Salut(%d) is %q, want %q", i, c.in, out, c.want)
		}
	}
}
