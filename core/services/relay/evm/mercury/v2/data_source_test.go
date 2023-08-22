package mercury_v2

import (
	"context"
	"math/big"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	relaymercury "github.com/smartcontractkit/chainlink-relay/pkg/reportingplugins/mercury"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	mercurymocks "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/types"
)

var _ relaymercury.MercuryServerFetcher = &mockFetcher{}

type mockFetcher struct {
	ts             uint32
	tsErr          error
	linkPrice      *big.Int
	linkPriceErr   error
	nativePrice    *big.Int
	nativePriceErr error
}

var feedId types.FeedID = [32]byte{1}
var linkFeedId types.FeedID = [32]byte{2}
var nativeFeedId types.FeedID = [32]byte{3}

func (m *mockFetcher) FetchInitialMaxFinalizedBlockNumber(context.Context) (*int64, error) {
	return nil, nil
}

func (m *mockFetcher) LatestPrice(ctx context.Context, fId [32]byte) (*big.Int, error) {
	if fId == linkFeedId {
		return m.linkPrice, m.linkPriceErr
	} else if fId == nativeFeedId {
		return m.nativePrice, m.nativePriceErr
	}
	return nil, nil
}

func (m *mockFetcher) LatestTimestamp(context.Context) (uint32, error) {
	return m.ts, m.tsErr
}

func Test_Datasource(t *testing.T) {
	ds := &datasource{lggr: logger.TestLogger(t)}
	ctx := testutils.Context(t)
	repts := ocrtypes.ReportTimestamp{}

	fetcher := &mockFetcher{}
	ds.fetcher = fetcher

	goodTrrs := []pipeline.TaskRunResult{
		{
			// bp
			Result: pipeline.Result{Value: "122.345"},
			Task:   &mercurymocks.MockTask{},
		},
	}

	ds.pipelineRunner = &mercurymocks.MockRunner{
		Trrs: goodTrrs,
	}

	spec := pipeline.Spec{}
	ds.spec = spec

	t.Run("when fetchMaxFinalizedTimestamp=true", func(t *testing.T) {
		t.Run("if LatestTimestamp returns error", func(t *testing.T) {
			fetcher.tsErr = errors.New("some error")

			obs, err := ds.Observe(ctx, repts, true)
			assert.NoError(t, err)

			assert.EqualError(t, obs.MaxFinalizedTimestamp.Err, "some error")
			assert.Zero(t, obs.MaxFinalizedTimestamp.Val)
		})

		t.Run("if LatestTimestamp succeeds", func(t *testing.T) {
			fetcher.tsErr = nil
			fetcher.ts = 123

			obs, err := ds.Observe(ctx, repts, true)
			assert.NoError(t, err)

			assert.Equal(t, uint32(123), obs.MaxFinalizedTimestamp.Val)
			assert.NoError(t, obs.MaxFinalizedTimestamp.Err)
		})

		t.Run("if LatestTimestamp succeeds but ts=0 (new feed)", func(t *testing.T) {
			fetcher.tsErr = nil
			fetcher.ts = 0

			obs, err := ds.Observe(ctx, repts, true)
			assert.NoError(t, err)

			assert.NoError(t, obs.MaxFinalizedTimestamp.Err)
			assert.Zero(t, obs.MaxFinalizedTimestamp.Val)
		})

		t.Run("when run execution succeeded", func(t *testing.T) {
			t.Run("when feedId=linkFeedID=nativeFeedId", func(t *testing.T) {
				t.Cleanup(func() {
					ds.feedID, ds.linkFeedID, ds.nativeFeedID = feedId, linkFeedId, nativeFeedId
				})

				ds.feedID, ds.linkFeedID, ds.nativeFeedID = feedId, feedId, feedId

				fetcher.ts = 123123
				fetcher.tsErr = nil

				obs, err := ds.Observe(ctx, repts, true)
				assert.NoError(t, err)

				assert.Equal(t, big.NewInt(122), obs.BenchmarkPrice.Val)
				assert.NoError(t, obs.BenchmarkPrice.Err)
				assert.Equal(t, uint32(123123), obs.MaxFinalizedTimestamp.Val)
				assert.NoError(t, obs.MaxFinalizedTimestamp.Err)
				assert.Equal(t, big.NewInt(122), obs.LinkPrice.Val)
				assert.NoError(t, obs.LinkPrice.Err)
				assert.Equal(t, big.NewInt(122), obs.NativePrice.Val)
				assert.NoError(t, obs.NativePrice.Err)
			})
		})
	})

	t.Run("when fetchMaxFinalizedTimestamp=false", func(t *testing.T) {
		t.Run("when run execution fails, returns error", func(t *testing.T) {
			t.Cleanup(func() {
				ds.pipelineRunner = &mercurymocks.MockRunner{
					Trrs: goodTrrs,
					Err:  nil,
				}
			})

			ds.pipelineRunner = &mercurymocks.MockRunner{
				Trrs: goodTrrs,
				Err:  errors.New("run execution failed"),
			}

			_, err := ds.Observe(ctx, repts, false)
			assert.EqualError(t, err, "Observe failed while executing run: error executing run for spec ID 0: run execution failed")
		})

		t.Run("when parsing run results fails, return error", func(t *testing.T) {
			t.Cleanup(func() {
				runner := &mercurymocks.MockRunner{
					Trrs: goodTrrs,
					Err:  nil,
				}
				ds.pipelineRunner = runner
			})

			badTrrs := []pipeline.TaskRunResult{
				{
					// benchmark price
					Result: pipeline.Result{Error: errors.New("some error with bp")},
					Task:   &mercurymocks.MockTask{},
				},
			}

			ds.pipelineRunner = &mercurymocks.MockRunner{
				Trrs: badTrrs,
				Err:  nil,
			}

			_, err := ds.Observe(ctx, repts, false)
			assert.EqualError(t, err, "Observe failed while parsing run results: some error with bp")
		})

		t.Run("when run execution succeeded", func(t *testing.T) {
			t.Run("when feedId=linkFeedID=nativeFeedId", func(t *testing.T) {
				t.Cleanup(func() {
					ds.feedID, ds.linkFeedID, ds.nativeFeedID = feedId, linkFeedId, nativeFeedId
				})

				var feedId types.FeedID = [32]byte{1}
				ds.feedID, ds.linkFeedID, ds.nativeFeedID = feedId, feedId, feedId

				obs, err := ds.Observe(ctx, repts, false)
				assert.NoError(t, err)

				assert.Equal(t, big.NewInt(122), obs.BenchmarkPrice.Val)
				assert.NoError(t, obs.BenchmarkPrice.Err)
				assert.Equal(t, uint32(0), obs.MaxFinalizedTimestamp.Val)
				assert.NoError(t, obs.MaxFinalizedTimestamp.Err)
				assert.Equal(t, big.NewInt(122), obs.LinkPrice.Val)
				assert.NoError(t, obs.LinkPrice.Err)
				assert.Equal(t, big.NewInt(122), obs.NativePrice.Val)
				assert.NoError(t, obs.NativePrice.Err)
			})

			t.Run("when fails to fetch linkPrice or nativePrice", func(t *testing.T) {
				t.Cleanup(func() {
					fetcher.linkPriceErr = nil
					fetcher.nativePriceErr = nil
				})

				fetcher.linkPriceErr = errors.New("some error fetching link price")
				fetcher.nativePriceErr = errors.New("some error fetching native price")

				obs, err := ds.Observe(ctx, repts, false)
				assert.NoError(t, err)

				assert.Nil(t, obs.LinkPrice.Val)
				assert.EqualError(t, obs.LinkPrice.Err, "some error fetching link price")
				assert.Nil(t, obs.NativePrice.Val)
				assert.EqualError(t, obs.NativePrice.Err, "some error fetching native price")
			})

			t.Run("when succeeds to fetch linkPrice or nativePrice but got nil (new feed)", func(t *testing.T) {
				obs, err := ds.Observe(ctx, repts, false)
				assert.NoError(t, err)

				assert.Equal(t, obs.LinkPrice.Val, maxInt192)
				assert.Nil(t, obs.LinkPrice.Err)
				assert.Equal(t, obs.NativePrice.Val, maxInt192)
				assert.Nil(t, obs.NativePrice.Err)
			})
		})
	})
}
