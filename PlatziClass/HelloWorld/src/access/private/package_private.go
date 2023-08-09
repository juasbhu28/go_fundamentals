package private

import (
	"fmt"

	"github.com/go_fundamentals/PlatziClass/HelloWorld/src/access/public"
)

type carPrivate struct {
	// Propiedades privadas
	privatebrand string
	privateYear  int
}

func PrintCarPrivate(c public.CarPublic) {
	var cPrivate carPrivate
	cPrivate.privatebrand = c.PublicBrand
	cPrivate.privateYear = c.PublicYear
	fmt.Println("Car private", cPrivate)
}
