package simple

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	provider := newClient("qq")
	rs := provider.send()
	fmt.Println(rs)
}
