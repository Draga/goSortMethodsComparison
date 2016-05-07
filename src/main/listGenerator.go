package main

import (
	"math/rand"
	"errors"
)

type ListOrder int

const (
	ordered ListOrder = iota
	inverse
	random
)

func generateList(length uint32, listOrder ListOrder) (*[]uint32, error) {
	list := make([]uint32, length)
	switch listOrder {
		case ordered:
			for i := uint32(0); i < length; i++ {
				list[i] = i;
			}
		case inverse:
			for i := length - 1; i >= 0; i-- {
				list[i] = i;
			}
		case random:
			for i := uint32(0); i < length; i++ {
				list[i] = rand.Uint32();
			}
		default:
			return nil, errors.New("Unimplemented list order #" + string(listOrder))
	}

	return &list, nil
}