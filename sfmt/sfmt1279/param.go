package sfmt1279

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = 1279

	// the pick up position of the array.
	pos1 = 7

	// the parameter of shift left as four 32-bit registers.
	sl1 = 14

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = 3

	// the parameter of shift right as four 32-bit registers.
	sr1 = 5

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = 1

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = 0x7fefcffff7fefffd & maskR
	msk2 = 0xb5ffff7faff3ef3f & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = 0x0000000000000001
	parity2 = 0x2000000000000000
)
