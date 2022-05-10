package sfmt216091

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = 216091

	// the pick up position of the array.
	pos1 = 627

	// the parameter of shift left as four 32-bit registers.
	sl1 = 11

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = 3

	// the parameter of shift right as four 32-bit registers.
	sr1 = 10

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = 1

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = 0xbfffffffbff7bff7 & maskR
	msk2 = 0xffddfbfbbffffa7f & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = 0x89e80709f8000001
	parity2 = 0x0c64b1e43bd2b64b
)
