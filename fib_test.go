package main

import "testing"

func TestFib(t *testing.T) {
	t.Run("n=1", func(t *testing.T) {
		got := Fib(1)
		want := "1"
		if got != want {
			t.Errorf("Fib(1) == %s, want %s", got, want)
		}
	})
	t.Run("n=99", func(t *testing.T) {
		got := Fib(99)
		want := "218922995834555169026"
		if got != want {
			t.Errorf("Fib(99) == %s, want %s", got, want)
		}
	})
	t.Run("n=130", func(t *testing.T) {
		got := Fib(130)
		want := "659034621587630041982498215"
		if got != want {
			t.Errorf("Fib(130) == %s, want %s", got, want)
		}
	})
	t.Run("n=148", func(t *testing.T) {
		got := Fib(148)
		want := "3807901929474025356630904134051"
		if got != want {
			t.Errorf("Fib(148) == %s, want %s", got, want)
		}
	})
	t.Run("n=279", func(t *testing.T) {
		got := Fib(279)
		want := "9079598147510263717870894449029933369491131786514446266146"
		if got != want {
			t.Errorf("Fib(279) == %s, want %s", got, want)
		}
	})
	t.Run("n=539", func(t *testing.T) {
		got := Fib(539)
		want := "19717556437089196581097595986132868777484357362710277396538804282697604657932286856665006090315514324457397722361"
		if got != want {
			t.Errorf("Fib(539) == %s, want %s", got, want)
		}
	})
	t.Run("n=931", func(t *testing.T) {
		got := Fib(931)
		want := "165199249717838008158743058245780420215766768590131298708272645987824841020293881889311405773866850429566420076904922158529659340892491762955827899358214010685748100753200453335055149114435266869"
		if got != want {
			t.Errorf("Fib(931) == %s, want %s", got, want)
		}
	})
	t.Run("n=1000", func(t *testing.T) {
		got := Fib(1000)
		want := "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875"
		if got != want {
			t.Errorf("Fib(1000) == %s, want %s", got, want)
		}
	})
}
