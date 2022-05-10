package sfmt11213

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = 11213

	// the pick up position of the array.
	pos1 = 68

	// the parameter of shift left as four 32-bit registers.
	sl1 = 14

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = 3

	// the parameter of shift right as four 32-bit registers.
	sr1 = 7

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = 3

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = 0xffffffefeffff7fb & maskR
	msk2 = 0x7fffdbfddfdfbfff & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = 0x0000000000000001
	parity2 = 0xd0c7afa3e8148000
)
