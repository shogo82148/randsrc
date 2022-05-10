package sfmt132049

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = 132049

	// the pick up position of the array.
	pos1 = 110

	// the parameter of shift left as four 32-bit registers.
	sl1 = 19

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = 1

	// the parameter of shift right as four 32-bit registers.
	sr1 = 21

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = 1

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = 0xfb6ebf95ffffbb5f & maskR
	msk2 = 0xcff77ffffffefffa & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = 0x0000000000000001
	parity2 = 0xc7e91c7dcb520000
)
