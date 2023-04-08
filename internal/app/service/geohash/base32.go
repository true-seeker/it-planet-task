package geohash

// invalid is a placeholder for invalid character decodings.
const invalid = 0xff

// encoding encapsulates an encoding defined by a given base32 alphabet.
type encoding struct {
	encode string
	decode [256]byte
}

// newEncoding constructs a new encoding defined by the given alphabet,
// which must be a 32-byte string.
func newEncoding(encoder string) *encoding {
	e := new(encoding)
	e.encode = encoder
	for i := 0; i < len(e.decode); i++ {
		e.decode[i] = invalid
	}
	for i := 0; i < len(encoder); i++ {
		e.decode[encoder[i]] = byte(i)
	}
	return e
}

// Encode bits of 64-bit word into a string.
func (e *encoding) Encode(x uint64) string {
	b := [12]byte{}
	for i := 0; i < 12; i++ {
		b[11-i] = e.encode[x&0x1f]
		x >>= 5
	}
	return string(b[:])
}

// Base32Encoding with the Geohash alphabet.
var base32encoding = newEncoding("0123456789bcdefghjkmnpqrstuvwxyz")
