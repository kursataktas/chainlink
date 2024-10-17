package utils

import (
	"math"
	"math/big"
	"strings"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// ToDecimal converts an input to a decimal
func ToDecimal(input interface{}) (decimal.Decimal, error) {
	switch v := input.(type) {
	case string:
		answer, err := decimal.NewFromString(v)
		if err == nil {
			return answer, nil
		} else {
			hexAnswer, hexErr := hexStringToDecimal(v)
			if hexErr {
				return answer, err
			} else {
				return hexAnswer, nil
			}
		}
	case int:
		return decimal.New(int64(v), 0), nil
	case int8:
		return decimal.New(int64(v), 0), nil
	case int16:
		return decimal.New(int64(v), 0), nil
	case int32:
		return decimal.New(int64(v), 0), nil
	case int64:
		return decimal.New(v, 0), nil
	case uint:
		return decimal.New(int64(v), 0), nil
	case uint8:
		return decimal.New(int64(v), 0), nil
	case uint16:
		return decimal.New(int64(v), 0), nil
	case uint32:
		return decimal.New(int64(v), 0), nil
	case uint64:
		return decimal.New(int64(v), 0), nil
	case float64:
		if !validFloat(v) {
			return decimal.Decimal{}, errors.Errorf("invalid float %v, cannot convert to decimal", v)
		}
		return decimal.NewFromFloat(v), nil
	case float32:
		if !validFloat(float64(v)) {
			return decimal.Decimal{}, errors.Errorf("invalid float %v, cannot convert to decimal", v)
		}
		return decimal.NewFromFloat32(v), nil
	case big.Int:
		return decimal.NewFromBigInt(&v, 0), nil
	case *big.Int:
		return decimal.NewFromBigInt(v, 0), nil
	case decimal.Decimal:
		return v, nil
	case *decimal.Decimal:
		return *v, nil
	default:
		return decimal.Decimal{}, errors.Errorf("type %T cannot be converted to decimal.Decimal (%v)", input, input)
	}
}

func hexStringToDecimal(hexString string) (decimal.Decimal, bool) {
	hexString = strings.TrimPrefix(hexString, "0x")
	n := new(big.Int)
	_, err := n.SetString(hexString, 16)
	return decimal.NewFromBigInt(n, 0), err
}

func validFloat(f float64) bool {
	return !math.IsNaN(f) && !math.IsInf(f, 0)
}
