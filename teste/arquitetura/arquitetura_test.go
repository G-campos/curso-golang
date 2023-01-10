package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	// if runtime.GOARCH == "amd64" {
	// 	t.Skip("Não funciona em arquitetura amd64")
	// }
	if runtime.GOOS != "linux" {
		t.Skip("Feito para linux S2")
	}
	// ...
	t.Fail()
}
