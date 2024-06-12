package logprovider

import (
	"math/big"
	"sync"
)

type DequeueCoordinator interface {
	// DequeueBlockWindow identifies a block window ready for processing between the given start and latest block numbers.
	// It prioritizes windows that need to have the minimum guaranteed logs dequeued before considering windows with
	// remaining logs to be dequeued, as a best effort.
	DequeueBlockWindow(start int64, latestBlock int64, blockRate int) (int64, int64, bool)
	// GetUpkeepSelector returns a function that accepts an upkeep ID, and performs a modulus against the number of
	// iterations, and compares the result against the current iteration. When this comparison returns true, the
	// upkeep is selected for the dequeuing. This means that, for a given set of upkeeps, a different subset of
	// upkeeps will be dequeued for each iteration once only, and, across all iterations, all upkeeps will be
	// dequeued once.
	GetUpkeepSelector(startWindow int64, logLimitLow, iterations, currentIteration int) func(id *big.Int) bool
	// TrackUpkeeps tracks how many times an upkeep has been dequeued for a given block window.
	TrackUpkeeps(startWindow int64, upkeepID *big.Int)
	// UpdateBlockWindow updates the status of a block window based on the number of logs dequeued,
	// remaining logs, and the number of upkeeps. This function tracks remaining and dequeued logs for the specified
	// block window, determines if a block window has had the minimum number of guaranteed logs dequeued, and marks a
	// window as not ready if there are not yet any logs available to dequeue from the window.
	UpdateBlockWindow(startWindow int64, logs, remaining, numberOfUpkeeps, logLimitLow int)
	// MarkReorg handles the detection of a reorg  by resetting the state of the affected block window. It ensures that
	// upkeeps within the specified block window are marked as not having the minimum number of guaranteed logs dequeued.
	MarkReorg(block int64, blockRate uint32)
}

type dequeueCoordinator struct {
	dequeuedMinimum map[int64]bool
	notReady        map[int64]bool
	remainingLogs   map[int64]int
	dequeuedLogs    map[int64]int
	completeWindows map[int64]bool
	dequeuedUpkeeps map[int64]map[string]int
	mu              sync.Mutex
}

func NewDequeueCoordinator() *dequeueCoordinator {
	return &dequeueCoordinator{
		dequeuedMinimum: map[int64]bool{},
		notReady:        map[int64]bool{},
		remainingLogs:   map[int64]int{},
		dequeuedLogs:    map[int64]int{},
		completeWindows: map[int64]bool{},
		dequeuedUpkeeps: map[int64]map[string]int{},
	}
}

func (c *dequeueCoordinator) DequeueBlockWindow(start int64, latestBlock int64, blockRate int) (int64, int64, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// check if minimum logs have been dequeued
	for i := start; i <= latestBlock; i += int64(blockRate) {
		startWindow, end := getBlockWindow(i, blockRate)
		if latestBlock >= end {
			c.completeWindows[startWindow] = true
		} else if c.notReady[startWindow] { // the window is incomplete and has no logs to provide as of yet
			break
		}

		if hasDequeued, ok := c.dequeuedMinimum[startWindow]; !ok || !hasDequeued {
			return startWindow, end, true
		}
	}

	// check best effort dequeue
	for i := start; i < latestBlock; i += int64(blockRate) {
		startWindow, end := getBlockWindow(i, blockRate)

		if remainingLogs, ok := c.remainingLogs[startWindow]; ok {
			if remainingLogs > 0 {
				return startWindow, end, true
			}
		}
	}

	return 0, 0, false
}

func (c *dequeueCoordinator) GetUpkeepSelector(startWindow int64, logLimitLow, iterations, currentIteration int) func(id *big.Int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	bestEffort := false

	if hasDequeued, ok := c.dequeuedMinimum[startWindow]; ok {
		if hasDequeued {
			bestEffort = true
		}
	}

	return func(id *big.Int) bool {
		// query the map of block number to upkeep ID for dequeued count here when the block window is incomplete
		dequeueUpkeep := true
		if !bestEffort {
			if windowUpkeeps, ok := c.dequeuedUpkeeps[startWindow]; ok {
				if windowUpkeeps[id.String()] >= logLimitLow {
					dequeueUpkeep = false
				}
			}
		}
		return dequeueUpkeep && id.Int64()%int64(iterations) == int64(currentIteration)
	}
}

func (c *dequeueCoordinator) TrackUpkeeps(startWindow int64, upkeepID *big.Int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if windowUpkeeps, ok := c.dequeuedUpkeeps[startWindow]; ok {
		windowUpkeeps[upkeepID.String()] = windowUpkeeps[upkeepID.String()] + 1
		c.dequeuedUpkeeps[startWindow] = windowUpkeeps
	} else {
		c.dequeuedUpkeeps[startWindow] = map[string]int{
			upkeepID.String(): 1,
		}
	}
}

func (c *dequeueCoordinator) UpdateBlockWindow(startWindow int64, logs, remaining, numberOfUpkeeps, logLimitLow int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.remainingLogs[startWindow] = remaining
	c.dequeuedLogs[startWindow] += logs

	if isComplete, ok := c.completeWindows[startWindow]; ok {
		if isComplete {
			// if the window is complete, and there are no more logs, then we have to consider this as min dequeued, even if no logs were dequeued
			if c.remainingLogs[startWindow] == 0 || c.dequeuedLogs[startWindow] >= numberOfUpkeeps*logLimitLow {
				c.dequeuedMinimum[startWindow] = true
			}
		} else if c.dequeuedLogs[startWindow] >= numberOfUpkeeps*logLimitLow { // this assumes we don't dequeue the same upkeeps more than logLimitLow in min commitment
			c.dequeuedMinimum[startWindow] = true
		}
	} else if c.dequeuedLogs[startWindow] >= numberOfUpkeeps*logLimitLow { // this assumes we don't dequeue the same upkeeps more than logLimitLow in min commitment
		c.dequeuedMinimum[startWindow] = true
	} else if logs == 0 && remaining == 0 {
		c.notReady[startWindow] = true
	}
}

func (c *dequeueCoordinator) MarkReorg(block int64, blockRate uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	startWindow, _ := getBlockWindow(block, int(blockRate))
	c.dequeuedMinimum[startWindow] = false
	// TODO instead of wiping the count for all upkeeps, should we wipe for upkeeps only impacted by the reorg?
	for upkeepID := range c.dequeuedUpkeeps[startWindow] {
		c.dequeuedUpkeeps[startWindow][upkeepID] = 0
	}
}
