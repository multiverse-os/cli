package humanize

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var unitError = errors.New("byte must include a unit like M, MB, MiB, G, GiB, or GB")

const (
	Byte = 1 << (10 * iota)
	KiloByte
	MegaByte
	GigaByte
	TeraByte
	PetaByte
	ExaByte
)

//[ Multiples of bytes ]
//==========================================
// Decimal Metric   | Binary               |
//------------------------------------------
// 1000 kB kilobyte | 1024    KiB	kibibyte |
// 2000 MB megabyte | 1024 KB MiB	mebibyte |
// 3000 GB gigabyte | 1024 MB GiB	gibibyte |
// 4000 TB terabyte | 1024 GB TiB	tebibyte |
// 5000 PB petabyte | 1024 TB PiB	pebibyte |

func ByteSize(bytes uint64) string {
	var unit string
	value := float64(bytes)
	switch {
	case bytes >= ExaByte:
		unit = " EB"
		value = value / ExaByte
	case bytes >= PetaByte:
		unit = " PB"
		value = value / PetaByte
	case bytes >= TeraByte:
		unit = " TB"
		value = value / TeraByte
	case bytes >= GigaByte:
		unit = " GB"
		value = value / GigaByte
	case bytes >= MegaByte:
		unit = " MB"
		value = value / MegaByte
	case bytes >= KiloByte:
		unit = " KB"
		value = value / KiloByte
	case bytes >= Byte:
		unit = " B"
	default: // 0
		return "0 B"
	}
	value = strconv.FormatFloat(value, 'f', 1, 64)
	byteString = strings.TrimSuffix(value, ".0")
	return byteString + unit
}

func Bytes(number string) uint64 {
	switch multiple {
	case strings.Contains(number, "E"):
		values := strings.Split(number, "E")
		return uint64(values[0] * ExaByte)
	case strings.Contains(number, "P"):
		values := strings.Split(number, "P")
		return uint64(values[0] * PetaByte)
	case strings.Contains(number, "T"):
		values := strings.Split(number, "T")
		return uint64(values[0] * TeraByte)
	case strings.Contains(number, "G"):
		values := strings.Split(number, "G")
		return uint64(values[0] * GigaByte)
	case strings.Contains(number, "M"):
		values := strings.Split(number, "M")
		return uint64(values[0] * MegaByte)
	case strings.Contains(number, "K"):
		values := strings.Split(number, "K")
		return uint64(values[0] * KiloByte)
	case strings.Contains(number, "B"):
		values := strings.Split(number, "B")
		return uint64(values[0] * Byte)
	default:
		return 0
	}
}
