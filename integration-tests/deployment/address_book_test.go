package deployment

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestAddressBook_Save(t *testing.T) {
	ab := NewMemoryAddressBook()
	onRamp100 := NewTypeAndVersion("OnRamp", Version1_0_0)
	onRamp110 := NewTypeAndVersion("OnRamp", Version1_1_0)
	addr1 := common.HexToAddress("0x1").String()
	addr2 := common.HexToAddress("0x2").String()

	err := ab.Save(chainsel.TEST_90000001.Selector, addr1, onRamp100)
	require.NoError(t, err)

	// Invalid address
	err = ab.Save(chainsel.TEST_90000001.Selector, "asdlfkj", onRamp100)
	require.Error(t, err)
	assert.Equal(t, errors.Is(err, ErrInvalidAddress), true, "err %s", err)

	// Valid chain but not present.
	_, err = ab.AddressesForChain(chainsel.TEST_90000002.Selector)
	assert.Equal(t, errors.Is(err, ErrChainNotFound), true, "err %s", err)

	// Invalid selector
	err = ab.Save(0, addr1, onRamp100)
	require.Error(t, err)
	assert.Equal(t, errors.Is(err, ErrInvalidChainSelector), true)

	// Duplicate
	err = ab.Save(chainsel.TEST_90000001.Selector, addr1, onRamp100)
	require.Error(t, err)

	// Zero address
	err = ab.Save(chainsel.TEST_90000001.Selector, common.HexToAddress("0x0").Hex(), onRamp100)
	require.Error(t, err)

	// Distinct address same TV will not
	err = ab.Save(chainsel.TEST_90000001.Selector, addr2, onRamp100)
	require.NoError(t, err)
	// Same address different chain will not error
	err = ab.Save(chainsel.TEST_90000002.Selector, addr1, onRamp100)
	require.NoError(t, err)
	// We can save different versions of the same contract
	err = ab.Save(chainsel.TEST_90000002.Selector, addr2, onRamp110)
	require.NoError(t, err)

	addresses, err := ab.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, addresses, map[uint64]map[string]TypeAndVersion{
		chainsel.TEST_90000001.Selector: {
			addr1: onRamp100,
			addr2: onRamp100,
		},
		chainsel.TEST_90000002.Selector: {
			addr1: onRamp100,
			addr2: onRamp110,
		},
	})
}

func TestAddressBook_Merge(t *testing.T) {
	onRamp100 := NewTypeAndVersion("OnRamp", Version1_0_0)
	onRamp110 := NewTypeAndVersion("OnRamp", Version1_1_0)
	addr1 := common.HexToAddress("0x1").String()
	addr2 := common.HexToAddress("0x2").String()
	a1 := NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
		chainsel.TEST_90000001.Selector: {
			addr1: onRamp100,
		},
	})
	a2 := NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
		chainsel.TEST_90000001.Selector: {
			addr2: onRamp100,
		},
		chainsel.TEST_90000002.Selector: {
			addr1: onRamp110,
		},
	})
	require.NoError(t, a1.Merge(a2))

	addresses, err := a1.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, addresses, map[uint64]map[string]TypeAndVersion{
		chainsel.TEST_90000001.Selector: {
			addr1: onRamp100,
			addr2: onRamp100,
		},
		chainsel.TEST_90000002.Selector: {
			addr1: onRamp110,
		},
	})

	// Merge with conflicting addresses should error
	a3 := NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
		chainsel.TEST_90000001.Selector: {
			addr1: onRamp100,
		},
	})
	require.Error(t, a1.Merge(a3))
	// a1 should not have changed
	addresses, err = a1.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, addresses, map[uint64]map[string]TypeAndVersion{
		chainsel.TEST_90000001.Selector: {
			addr1: onRamp100,
			addr2: onRamp100,
		},
		chainsel.TEST_90000002.Selector: {
			addr1: onRamp110,
		},
	})
}

func TestAddressBookMap_Diff(t *testing.T) {
	addr1 := common.HexToAddress("0x1").String()
	addr2 := common.HexToAddress("0x2").String()
	type fields struct {
		AddressesByChain map[uint64]map[string]TypeAndVersion
	}
	type args struct {
		other AddressBook
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    AddressBook
		wantErr bool
	}{
		{
			name: "exact no diff",
			fields: fields{
				AddressesByChain: map[uint64]map[string]TypeAndVersion{
					chainsel.TEST_90000001.Selector: {
						addr1: NewTypeAndVersion("OnRamp", Version1_0_0),
					},
				},
			},
			args: args{
				other: NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
					chainsel.TEST_90000001.Selector: {
						addr1: NewTypeAndVersion("OnRamp", Version1_0_0),
					},
				}),
			},
			want: NewMemoryAddressBook(),
		},
		{
			name: "subset no diff",
			fields: fields{
				AddressesByChain: map[uint64]map[string]TypeAndVersion{
					chainsel.TEST_90000001.Selector: {
						addr1: NewTypeAndVersion("OnRamp", Version1_0_0),
						addr2: NewTypeAndVersion("ExistingContract", Version1_0_0),
					},
				},
			},
			args: args{
				other: NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
					chainsel.TEST_90000001.Selector: {
						addr1: NewTypeAndVersion("OnRamp", Version1_0_0),
					},
				}),
			},
			want: NewMemoryAddressBook(),
		},
		{
			name: "superset diff",
			fields: fields{
				AddressesByChain: map[uint64]map[string]TypeAndVersion{
					chainsel.TEST_90000001.Selector: {
						addr1: NewTypeAndVersion("OnRamp", Version1_0_0),
					},
				},
			},
			args: args{
				other: NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
					chainsel.TEST_90000001.Selector: {
						addr1: NewTypeAndVersion("OnRamp", Version1_0_0),
						addr2: NewTypeAndVersion("NewContract", Version1_0_0),
					},
				}),
			},
			want: NewMemoryAddressBookFromMap(map[uint64]map[string]TypeAndVersion{
				chainsel.TEST_90000001.Selector: {
					addr2: NewTypeAndVersion("NewContract", Version1_0_0),
				},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AddressBookMap{
				AddressesByChain: tt.fields.AddressesByChain,
			}
			got, err := m.Diff(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressBookMap.Diff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressBookMap.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
