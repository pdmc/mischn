// Recommendations
//
// The SHAKE functions are recommended for most new uses. They can produce
// output of arbitrary length. SHAKE256, with an output length of at least
// 64 bytes, provides 256-bit security against all attacks.  The Keccak team
// recommends it for most applications upgrading from SHA2-512. (NIST chose a
// much stronger, but much slower, sponge instance for SHA3-512.)
//
// The SHA-3 functions are "drop-in" replacements for the SHA-2 functions.
// They produce output of the same length, with the same security strengths
// against all attacks. This means, in particular, that SHA3-256 only has
// 128-bit collision resistance, because its output length is 32 bytes.
package sha3
