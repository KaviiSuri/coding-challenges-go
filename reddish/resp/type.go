package resp

const TERMINATOR = "\r\n"

type DataType uint

const (
	SimpleString DataType = iota
	SimpleError
	Integer
	BulkString
	Array
	// Null
	// Boolean
	// Double
	// BigNumber
	// BulkError
	// VerbatimString
	// Map
	// Set
	// Push
)

var firstByteForType = [...]byte{
	SimpleString: '+',
	SimpleError:  '-',
	Integer:      ':',
	BulkString:   '$',
	Array:        '*',
	// Null:           '_',
	// Boolean:        '#',
	// Double:         ',',
	// BigNumber:      '(',
	// BulkError:      '!',
	// VerbatimString: '=',
	// Map:            '%',
	// Set:            '~',
	// Push:           '>',
}

var typeForFirstByte = map[byte]DataType{
	'+': SimpleString,
	'-': SimpleError,
	':': Integer,
	'$': BulkString,
	'*': Array,
	// '_': Null,
	// '#': Boolean,
	// ',': Double,
	// '(': BigNumber,
	// '!': BulkError,
	// '=': VerbatimString,
	// '%': Map,
	// '~': Set,
	// '>': Push,
}

func GetDataType(symbol byte) (DataType, bool) {
	rType, found := typeForFirstByte[symbol]
	return rType, found
}

func GetFirstByteFor(t DataType) byte {
	return firstByteForType[t]
}
