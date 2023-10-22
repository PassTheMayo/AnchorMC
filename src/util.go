package main

func pointerOf[T any](v T) *T {
	return &v
}
