package messages

import "testing"

func TestGreet(t *testing.T) {
	subTests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world\n",
		},
		{
			result: "Hello, Paul\n",
			items:  []string{"Paul"},
		},
		{
			result: "Hello, Paul, Ted\n",
			items:  []string{"Paul", "Ted"},
		},
		{
			result: "Hello, Paul, Ted, Zach\n",
			items:  []string{"Paul", "Ted", "Zach"},
		},
	}

	for _, st := range subTests {
		if s := Greet(st.items); s != st.result {
			t.Errorf("wanted %s, got %s", st.result, Greet(st.items))
		}
	}
}

func TestDepart(t *testing.T) {

}
