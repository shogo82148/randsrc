package sfmt86243

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = 86243

	// the pick up position of the array.
	pos1 = 366

	// the parameter of shift left as four 32-bit registers.
	sl1 = 6

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = 7

	// the parameter of shift right as four 32-bit registers.
	sr1 = 19

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = 1

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = 0xbff7ff3ffdbffbff & maskR
	msk2 = 0xbf9ff3fffd77efff & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = 0x0000000000000001
	parity2 = 0xe9528d8500000000
)
