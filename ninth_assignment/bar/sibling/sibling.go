package sibling

import "github.com/hkaya15/Go101/bar/internal"

func WorkWithBigger(s string)string{
	return internal.GetBigger(s)
}