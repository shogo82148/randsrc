package sfmt2281

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = 2281

	// the pick up position of the array.
	pos1 = 12

	// the parameter of shift left as four 32-bit registers.
	sl1 = 19

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = 1

	// the parameter of shift right as four 32-bit registers.
	sr1 = 5

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = 1

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = 0xfdfffffebff7ffbf & maskR
	msk2 = 0xf2f7cbbff7ffef7f & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = 0x0000000000000001
	parity2 = 0x41dfa60000000000
)
