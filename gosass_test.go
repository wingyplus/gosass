package gosass

import "testing"

func TestRender(t *testing.T) {
	scss := `
                $color: #fff;

                #main {
                        color: $color;
                }`
	expected := `#main {
  color: #fff; }
`

	b, _ := Render(scss)
	if css := string(b); css != expected {
		t.Errorf("Expect \n%s \nbut got \n%s", expected, css)
	}
}
