package hello

import "testing"



func TestGreetsGitHub(t *testing.T) {
	result := hello.Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}
