package challenge

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
)

type ExecutionImpl struct {
	initialGasUsed *big.Int
}

func (e ExecutionImpl) GetCuts(lookup core.ValidatorLookup, offsets []*big.Int) ([]core.Cut, error) {
	panic("implement me")
}

func (e ExecutionImpl) FindFirstDivergence(lookup core.ValidatorLookup, offsets []*big.Int, cuts []core.Cut) (int, error) {
	panic("implement me")
}

func (e ExecutionImpl) Bisect(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	inconsistentSegment *core.ChallengeSegment,
	subCuts []core.Cut,
) (*types.Transaction, error) {
	return challenge.BisectExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		inconsistentSegment,
		subCuts,
	)
}

func (e ExecutionImpl) OneStepProof(
	ctx context.Context,
	challenge *ethbridge.Challenge,
	lookup core.ValidatorLookup,
	prevBisection *core.Bisection,
	segmentToChallenge int,
	challengedSegment *core.ChallengeSegment,
) (*types.Transaction, error) {
	startMachine, err := lookup.GetMachine(e.initialGasUsed)
	if err != nil {
		return nil, err
	}
	beforeAssertion, err := lookup.GetExecutionInfo(startMachine, challengedSegment.Start)
	if err != nil {
		return nil, err
	}

	beforeMachine, err := lookup.GetMachine(new(big.Int).Add(e.initialGasUsed, challengedSegment.Start))
	if err != nil {
		return nil, err
	}

	proofData, err := beforeMachine.MarshalForProof()
	if err != nil {
		return nil, err
	}

	bufferProofData, err := beforeMachine.MarshalBufferProof()
	if err != nil {
		return nil, err
	}

	return challenge.OneStepProveExecution(
		ctx,
		prevBisection,
		segmentToChallenge,
		beforeAssertion,
		proofData,
		bufferProofData,
	)
}