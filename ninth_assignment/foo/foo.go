package foo

import (
	//"github.com/hkaya15/Go101/bar/internal"
	"github.com/hkaya15/Go101/bar/sibling"
)

func CantUseBigger(s string)string{
	// return internal.GetBigger(s) // Internal klasörü altındaki fonksiyonlar sadece siblingler tarafından kullanılabilir
	return sibling.WorkWithBigger(s)
}