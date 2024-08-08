package main_test

import "testing"

func BenchmarkValidateLoop3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateLoop("foo", noop, noop, noop)
	}
}

func BenchmarkValidateLoop5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateLoop("foo", noop, noop, noop, noop, noop)
	}
}

func BenchmarkValidateLoop10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateLoop("foo", noop, noop, noop, noop, noop, noop, noop, noop, noop, noop)
	}
}

func BenchmarkValidateDepth3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateDepth("foo", noop, noop, noop)
	}
}

func BenchmarkValidateDepth5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateDepth("foo", noop, noop, noop, noop, noop)
	}
}

func BenchmarkValidateDepth10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateDepth("foo", noop, noop, noop, noop, noop, noop, noop, noop, noop, noop)
	}
}

func BenchmarkValidateCombine3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateCombine3("foo")
	}
}

func BenchmarkValidateCombine5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateCombine5("foo")
	}
}

func BenchmarkValidateCombine10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateCombine10("foo")
	}
}

func BenchmarkValidatePlain3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validatePlain3("foo")
	}
}

func BenchmarkValidatePlain5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validatePlain5("foo")
	}
}

func BenchmarkValidatePlain10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validatePlain10("foo")
	}
}

func validatePlain3(s string) error {
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	return nil
}

func validatePlain5(s string) error {
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	return nil
}

func validatePlain10(s string) error {
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	if err := noop(s); err != nil {
		return err
	}
	return nil
}

func validateLoop(s string, fns ...func(string) error) error {
	for _, fn := range fns {
		if err := fn(s); err != nil {
			return err
		}
	}

	return nil
}

var validateCombine3 = combine(noop, noop, noop)
var validateCombine5 = combine(noop, noop, noop, noop, noop)
var validateCombine10 = combine(noop, noop, noop, noop, noop, noop, noop, noop, noop, noop)

func combine[T any](fns ...func(T) error) func(T) error {
	switch len(fns) {
	case 0:
		return nil
	case 1:
		return fns[0]
	default:
		a, b, c := fns[:len(fns)-2], fns[len(fns)-2], fns[len(fns)-1]
		return combine(append(a, func(v T) error {
			if err := b(v); err != nil {
				return err
			}

			return c(v)
		})...)
	}
}

func validateDepth(s string, fns ...func(string) error) error {
	if len(fns) == 0 {
		return nil
	}

	if err := fns[0](s); err != nil {
		return err
	}
	return validateDepth(s, fns[1:]...)
}

func noop(string) error {
	return nil
}
