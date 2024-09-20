package codec

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-viper/mapstructure/v2"

	commoncodec "github.com/smartcontractkit/chainlink-common/pkg/codec"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

// DecoderHooks
//
// decodeAccountAndAllowArraySliceHook allows:
//
//	strings to be converted to [32]byte allowing config to represent them as 0x...
//	slices or arrays to be converted to a pointer to that type
//
// BigIntHook allows *big.Int to be represented as any integer type or a string and to go back to them.
// Useful for config, or if when a model may use a go type that isn't a *big.Int when Pack expects one.
// Eg: int32 in a go struct from a plugin could require a *big.Int in Pack for int24, if it fits, we shouldn't care.
// SliceToArrayVerifySizeHook verifies that slices have the correct size when converting to an array
// sizeVerifyBigIntHook allows our custom types that verify the number fits in the on-chain type to be converted as-if
// it was a *big.Int
var DecoderHooks = []mapstructure.DecodeHookFunc{
	decodeAccountAndAllowArraySliceHook,
	commoncodec.BigIntHook,
	commoncodec.SliceToArrayVerifySizeHook,
	sizeVerifyBigIntHook,
	commoncodec.NumberHook,
	addressStringDecodeHook,
}

// NewCodec creates a new [commontypes.RemoteCodec] for EVM.
// Note that names in the ABI are converted to Go names using [abi.ToCamelCase],
// this is per convention in [abi.MakeTopics], [abi.Arguments.Pack] etc.
// This allows names on-chain to be in go convention when generated.
// It means that if you need to use a [codec.Modifier] to reference a field
// you need to use the Go name instead of the name on-chain.
// eg: rename FooBar -> Bar, not foo_bar_ to Bar if the name on-chain is foo_bar_
func NewCodec(conf types.CodecConfig) (commontypes.RemoteCodec, error) {
	parsed := &ParsedTypes{
		EncoderDefs: map[string]types.CodecEntry{},
		DecoderDefs: map[string]types.CodecEntry{},
	}

	for k, v := range conf.Configs {
		args := abi.Arguments{}
		if err := json.Unmarshal(([]byte)(v.TypeABI), &args); err != nil {
			return nil, err
		}

		mod, err := v.ModifierConfigs.ToModifier(DecoderHooks...)
		if err != nil {
			return nil, err
		}

		item := types.NewCodecEntry(args, nil, mod)
		if err = item.Init(); err != nil {
			return nil, err
		}

		parsed.EncoderDefs[k] = item
		parsed.DecoderDefs[k] = item
	}

	return parsed.ToCodec()
}

type evmCodec struct {
	*encoder
	*decoder
	*ParsedTypes
}

func (c *evmCodec) CreateType(itemType string, forEncoding bool) (any, error) {
	var itemTypes map[string]types.CodecEntry
	if forEncoding {
		itemTypes = c.EncoderDefs
	} else {
		itemTypes = c.DecoderDefs
	}

	def, ok := itemTypes[itemType]
	if !ok {
		return nil, fmt.Errorf("%w: cannot find type name %s", commontypes.ErrInvalidType, itemType)
	}

	return reflect.New(def.CheckedType()).Interface(), nil
}

func WrapItemType(contractName, itemType string, isParams bool) string {
	if isParams {
		return fmt.Sprintf("params.%s.%s", contractName, itemType)
	}

	return fmt.Sprintf("return.%s.%s", contractName, itemType)
}

var bigIntType = reflect.TypeOf((*big.Int)(nil))

func sizeVerifyBigIntHook(from, to reflect.Type, data any) (any, error) {
	if from.Implements(types.SizedBigIntType()) &&
		!to.Implements(types.SizedBigIntType()) &&
		!reflect.PointerTo(to).Implements(types.SizedBigIntType()) {
		return commoncodec.BigIntHook(from, bigIntType, reflect.ValueOf(data).Convert(bigIntType).Interface())
	}

	if !to.Implements(types.SizedBigIntType()) {
		return data, nil
	}

	var err error
	data, err = commoncodec.BigIntHook(from, bigIntType, data)
	if err != nil {
		return nil, err
	}

	bi, ok := data.(*big.Int)
	if !ok {
		return data, nil
	}

	converted := reflect.ValueOf(bi).Convert(to).Interface().(types.SizedBigInt)
	return converted, converted.Verify()
}

func decodeAccountAndAllowArraySliceHook(from, to reflect.Type, data any) (any, error) {
	if from.Kind() == reflect.String &&
		(to == reflect.TypeOf(common.Address{}) || to == reflect.TypeOf(&common.Address{})) {
		return decodeAddress(data)
	}

	if from.Kind() == reflect.Pointer && to.Kind() != reflect.Pointer && from != nil &&
		(from.Elem().Kind() == reflect.Slice || from.Elem().Kind() == reflect.Array) {
		return reflect.ValueOf(data).Elem().Interface(), nil
	}

	return data, nil
}

func decodeAddress(data any) (any, error) {
	decoded, err := hexutil.Decode(data.(string))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", commontypes.ErrInvalidType, err)
	} else if len(decoded) != common.AddressLength {
		return nil, fmt.Errorf(
			"%w: wrong number size for address expected %v got %v",
			commontypes.ErrSliceWrongLen,
			common.AddressLength, len(decoded))
	}

	return common.Address(decoded), nil
}

// addressStringDecodeHook converts between string and common.Address, and vice versa.
// It handles both converting a string (hex format) to a common.Address and converting a common.Address to a string.
func addressStringDecodeHook(from reflect.Type, to reflect.Type, value interface{}) (interface{}, error) {
	// Check if we're converting from string to common.Address
	if from == reflect.TypeOf("") && to == reflect.TypeOf(common.Address{}) {
		// Decode the string (hex format) to common.Address
		return decodeAddress(value)
	}

	// Check if we're converting from common.Address to string
	if from == reflect.TypeOf(common.Address{}) && to == reflect.TypeOf("") {
		// Convert the common.Address to its string (hex) representation
		return value.(common.Address).Hex(), nil
	}

	// If no valid conversion, return the original value unchanged
	return value, nil
}
