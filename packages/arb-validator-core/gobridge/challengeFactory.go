/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gobridge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type challengeFactory struct {
	contract common.Address
	client   *GoArbAuthClient
	auth     *TransOpts
}

func newChallengeFactory(address common.Address, client *GoArbAuthClient, auth *TransOpts) (*challengeFactory, error) {
	//vmCreatorContract, err := challengefactory.NewChallengeFactory(address, client)
	//if err != nil {
	//	return nil, errors2.Wrap(err, "Failed to connect to arbFactory")
	//}
	return &challengeFactory{address, client, auth}, nil
}

func (con *challengeFactory) CreateChallenge(
	ctx context.Context,
	asserter common.Address,
	challenger common.Address,
	challengePeriod common.TimeTicks,
	challengeHash common.Hash,
	challengeType *big.Int,
) (common.Address, error) {
	challenge := con.client.GoEthClient.getNextAddress()
	con.client.GoEthClient.challenges[challenge] = &challengeData{
		challengePeriodTicks: challengePeriod,
		asserter:             asserter,
		challenger:           challenger,
		challengeType:        challengeType,
	}
	return con.client.GoEthClient.getNextAddress(), nil
}