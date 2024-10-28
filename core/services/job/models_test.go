package job_test

import (
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/pelletier/go-toml/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	pkgworkflows "github.com/smartcontractkit/chainlink-common/pkg/workflows"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestOCR2OracleSpec_RelayIdentifier(t *testing.T) {
	type fields struct {
		Relay       string
		ChainID     string
		RelayConfig job.JSONConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.RelayID
		wantErr bool
	}{
		{name: "err no chain id",
			fields:  fields{},
			want:    types.RelayID{},
			wantErr: true,
		},
		{
			name: "evm explicitly configured",
			fields: fields{
				Relay:   relay.NetworkEVM,
				ChainID: "1",
			},
			want: types.RelayID{Network: relay.NetworkEVM, ChainID: "1"},
		},
		{
			name: "evm implicitly configured",
			fields: fields{
				Relay:       relay.NetworkEVM,
				RelayConfig: map[string]any{"chainID": 1},
			},
			want: types.RelayID{Network: relay.NetworkEVM, ChainID: "1"},
		},
		{
			name: "evm implicitly configured with bad value",
			fields: fields{
				Relay:       relay.NetworkEVM,
				RelayConfig: map[string]any{"chainID": float32(1)},
			},
			want:    types.RelayID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &job.OCR2OracleSpec{
				Relay:       tt.fields.Relay,
				ChainID:     tt.fields.ChainID,
				RelayConfig: tt.fields.RelayConfig,
			}
			got, err := s.RelayID()
			if (err != nil) != tt.wantErr {
				t.Errorf("OCR2OracleSpec.RelayIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OCR2OracleSpec.RelayIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	//go:embed testdata/compact.toml
	compact string
	//go:embed testdata/pretty.toml
	pretty string
)

func TestOCR2OracleSpec(t *testing.T) {
	val := job.OCR2OracleSpec{
		Relay:                             relay.NetworkEVM,
		PluginType:                        types.Median,
		ContractID:                        "foo",
		OCRKeyBundleID:                    null.StringFrom("bar"),
		TransmitterID:                     null.StringFrom("baz"),
		ContractConfigConfirmations:       1,
		ContractConfigTrackerPollInterval: *models.NewInterval(time.Second),
		RelayConfig: map[string]interface{}{
			"chainID":   1337,
			"fromBlock": 42,
			"chainReader": evmtypes.ChainReaderConfig{
				Contracts: map[string]evmtypes.ChainContractReader{
					"median": {
						ContractABI: `[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "requester",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "configDigest",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "epoch",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint8",
        "name": "round",
        "type": "uint8"
      }
    ],
    "name": "RoundRequested",
    "type": "event"
  },
  {
    "inputs": [],
    "name": "latestTransmissionDetails",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "configDigest",
        "type": "bytes32"
      },
      {
        "internalType": "uint32",
        "name": "epoch",
        "type": "uint32"
      },
      {
        "internalType": "uint8",
        "name": "round",
        "type": "uint8"
      },
      {
        "internalType": "int192",
        "name": "latestAnswer_",
        "type": "int192"
      },
      {
        "internalType": "uint64",
        "name": "latestTimestamp_",
        "type": "uint64"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]
`,
						Configs: map[string]*evmtypes.ChainReaderDefinition{
							"LatestTransmissionDetails": {
								ChainSpecificName: "latestTransmissionDetails",
								OutputModifications: codec.ModifiersConfig{
									&codec.EpochToTimeModifierConfig{
										Fields: []string{"LatestTimestamp_"},
									},
									&codec.RenameModifierConfig{
										Fields: map[string]string{
											"LatestAnswer_":    "LatestAnswer",
											"LatestTimestamp_": "LatestTimestamp",
										},
									},
								},
							},
							"LatestRoundRequested": {
								ChainSpecificName: "RoundRequested",
								ReadType:          evmtypes.Event,
							},
						},
					},
				},
			},
			"codec": evmtypes.CodecConfig{
				Configs: map[string]evmtypes.ChainCodecConfig{
					"MedianReport": {
						TypeABI: `[
  {
    "Name": "Timestamp",
    "Type": "uint32"
  },
  {
    "Name": "Observers",
    "Type": "bytes32"
  },
  {
    "Name": "Observations",
    "Type": "int192[]"
  },
  {
    "Name": "JuelsPerFeeCoin",
    "Type": "int192"
  }
]
`,
					},
				},
			},
		},
		OnchainSigningStrategy: map[string]interface{}{
			"strategyName": "single-chain",
			"config": map[string]interface{}{
				"evm":       "",
				"publicKey": "0xdeadbeef",
			},
		},
		PluginConfig: map[string]interface{}{"juelsPerFeeCoinSource": `  // data source 1
  ds1          [type=bridge name="%s"];
  ds1_parse    [type=jsonparse path="data"];
  ds1_multiply [type=multiply times=2];

  // data source 2
  ds2          [type=http method=GET url="%s"];
  ds2_parse    [type=jsonparse path="data"];
  ds2_multiply [type=multiply times=2];

  ds1 -> ds1_parse -> ds1_multiply -> answer1;
  ds2 -> ds2_parse -> ds2_multiply -> answer1;

  answer1 [type=median index=0];
`,
		},
	}

	t.Run("marshal", func(t *testing.T) {
		gotB, err := toml.Marshal(val)
		require.NoError(t, err)
		t.Log("marshaled:", string(gotB))
		require.Equal(t, compact, string(gotB))
	})

	t.Run("round-trip", func(t *testing.T) {
		var gotVal job.OCR2OracleSpec
		require.NoError(t, toml.Unmarshal([]byte(compact), &gotVal))
		gotB, err := toml.Marshal(gotVal)
		require.NoError(t, err)
		require.Equal(t, compact, string(gotB))
		t.Run("pretty", func(t *testing.T) {
			var gotVal job.OCR2OracleSpec
			require.NoError(t, toml.Unmarshal([]byte(pretty), &gotVal))
			gotB, err := toml.Marshal(gotVal)
			require.NoError(t, err)
			t.Log("marshaled compact:", string(gotB))
			require.Equal(t, compact, string(gotB))
		})
	})
}

func TestWorkflowSpec_Validate(t *testing.T) {
	type fields struct {
		Workflow string
	}
	tests := []struct {
		name              string
		fields            fields
		wantWorkflowOwner string
		wantWorkflowName  string

		wantError bool
	}{
		{
			name: "valid",
			fields: fields{
				Workflow: pkgworkflows.WFYamlSpec(t, "workflow01", "0x0123456789012345678901234567890123456789"),
			},
			wantWorkflowOwner: "0123456789012345678901234567890123456789", // the workflow job spec strips the 0x prefix to limit to 40	characters
			wantWorkflowName:  "workflow01",
		},
		{
			name: "valid no name",
			fields: fields{
				Workflow: pkgworkflows.WFYamlSpec(t, "", "0x0123456789012345678901234567890123456789"),
			},
			wantWorkflowOwner: "0123456789012345678901234567890123456789", // the workflow job spec strips the 0x prefix to limit to 40	characters
			wantWorkflowName:  "",
		},
		{
			name: "valid no owner",
			fields: fields{
				Workflow: pkgworkflows.WFYamlSpec(t, "workflow01", ""),
			},
			wantWorkflowOwner: "",
			wantWorkflowName:  "workflow01",
		},
		{
			name: "invalid ",
			fields: fields{
				Workflow: "garbage",
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &job.WorkflowSpec{
				Workflow: tt.fields.Workflow,
			}
			err := w.Validate(testutils.Context(t))
			require.Equal(t, tt.wantError, err != nil)
			if !tt.wantError {
				assert.NotEmpty(t, w.WorkflowID)
				assert.Equal(t, tt.wantWorkflowOwner, w.WorkflowOwner)
				assert.Equal(t, tt.wantWorkflowName, w.WorkflowName)
			}
		})
	}

	t.Run("WASM can validate", func(t *testing.T) {
		configLocation := "testdata/config.json"

		w := &job.WorkflowSpec{
			Workflow: createTestBinary(t),
			SpecType: job.WASMFile,
			Config:   configLocation,
		}

		err := w.Validate(testutils.Context(t))
		require.NoError(t, err)
		assert.Equal(t, "owner", w.WorkflowOwner)
		assert.Equal(t, "name", w.WorkflowName)
		require.NotEmpty(t, w.WorkflowID)
	})
}

func TestOEVConfig(t *testing.T) {
	tests := []struct {
		name                 string
		shouldError          bool
		expectedErrorMessage string
		config               job.OEVConfig
	}{
		{
			name:                 "OEVTransmitterAddress not set",
			shouldError:          true,
			expectedErrorMessage: "no OEVTransmitterAddress found",
			config: job.OEVConfig{
				TransmitterAddress: nil,
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"calldata"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVContractAddress not set",
			shouldError:          true,
			expectedErrorMessage: "no OEVContractAddress found",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    nil,
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"calldata"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVBuilder not set",
			shouldError:          true,
			expectedErrorMessage: "no OEVBuilders found",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           nil,
				Hints:              []string{"calldata"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVBuilder empty",
			shouldError:          true,
			expectedErrorMessage: "OEVBuilders expects at least one builder",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{},
				Hints:              []string{"calldata"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVBuilder has empty element",
			shouldError:          true,
			expectedErrorMessage: "OEVBuilders should not contain empty element",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "", "builder2"},
				Hints:              []string{"calldata"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVHints not set",
			shouldError:          true,
			expectedErrorMessage: "no OEVHints found",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              nil,
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVHints empty",
			shouldError:          true,
			expectedErrorMessage: "OEVHints expects at least one hint",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVHints has empty element",
			shouldError:          true,
			expectedErrorMessage: "OEVHints should not contain empty element",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "", "hint2"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVPriceDelay not set",
			shouldError:          true,
			expectedErrorMessage: "OEVPriceDelay not set or smaller than 1s",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "hint2"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 1,
					},
				},
			},
		},
		{
			name:                 "OEVPriceDelay set to 50ms",
			shouldError:          true,
			expectedErrorMessage: "OEVPriceDelay not set or smaller than 1s",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "hint2"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 0,
					},
				},
				PriceDelay: time.Millisecond * 50,
			},
		},
		{
			name:                 "OEVRefund not set",
			shouldError:          true,
			expectedErrorMessage: "no OEVRefund found",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "hint2"},
				PriceDelay:         time.Second * 30,
			},
		},
		{
			name:                 "OEVRefund.Address not set",
			shouldError:          true,
			expectedErrorMessage: "OEVRefund.Address should not be empty",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "hint2"},
				Refund: []job.OEVRefund{
					{
						Percent: 1,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVRefund.Percent not set",
			shouldError:          true,
			expectedErrorMessage: "OEVRefund.Percent should be between 1 and 100",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "hint2"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
		{
			name:                 "OEVRefund.Percent over 100",
			shouldError:          true,
			expectedErrorMessage: "the sum of all OEVRefund.Percent should not be greater than 100",
			config: job.OEVConfig{
				TransmitterAddress: ptr(cltest.NewEIP55Address()),
				ContractAddress:    ptr(cltest.NewEIP55Address()),
				Builders:           []string{"builder1", "builder2"},
				Hints:              []string{"hint1", "hint2"},
				Refund: []job.OEVRefund{
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 50,
					},
					{
						Address: ptr(cltest.NewEIP55Address()),
						Percent: 51,
					},
				},
				PriceDelay: time.Second * 30,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.shouldError {
				require.ErrorContains(t, test.config.Validate(), test.expectedErrorMessage)
			} else {
				require.NoError(t, test.config.Validate())
			}
		})
	}
}
