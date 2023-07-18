package base

type Charset string

// Base64Chars is the base 64 alphabet used for encoding based on RFC 4648

// Base64CharsURL The "URL and Filename safe" Base 64 Alphabet used for encoding based on RFC 4648

/**
 * Base32 RFC 4648 Base32 alphabet
 */

// Base32 RFC 4648 Base32 alphabet

// Base32Hex Base 32 Encoding with Extended Hex Alphabet
// https://datatracker.ietf.org/doc/html/rfc4648#section-7
const (
	base32    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	base32Hex = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
	// Base32Crockford Crockford's Base32 character set which removes the
	// letters I, L, O, and U from the alphabet to avoid confusion with digits.
	// https://www.crockford.com/base32.html
	base32Crockford = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	// Base32ZBase32 ZBase32 character set as described in http://philzimmermann.com/docs/human-oriented-base-32-encoding.txt
	base32ZBase32 = "ybndrfg8ejkmcpqxot1uwisza345h769"
	// Base32GeoHash GeoHash Base 32 Geohash character set
	base32GeoHash = "0123456789bcdefghjkmnpqrstuvwxyz"
	base16        = "0123456789ABCDEF"
	base8         = "01234567"
	// Base32WordSafe Word-safe Base 32 character set
	base32WordSafe = "23456789CFGHJMPQRVWXcfghjmpqrvwx"
	base52         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	base36 = "0123456789abcdefghijklmnopqrstuvwxyz"

	base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Base45 is the character set used by the Base45 encoding scheme as defined in rfc9285
	// https://www.rfc-editor.org/rfc/rfc9285
	base45         = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"
	ascii85        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~"
	base91         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""
	basE91         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""
	base64Charset  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	base64CharsURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	base2          = "01"
)
