// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package test_data

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/fakes"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/flop_kick"
	"math/big"
	"strconv"
	"time"
)

var (
	FlopKickLog = types.Log{
		Address: common.HexToAddress(KovanFlopperContractAddress),
		Topics: []common.Hash{
			common.HexToHash("0xefa52d9342a199cb30efd2692463f2c2bef63cd7186b50382d4fb94ad207880e"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000005"),
		},
		Data:        hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000f00000000000000000000000000000000000000000000000000000000000000020000000000000000000000007d7bee5fcfd8028cf7b00876c5b1421c800561a6000000000000000000000000000000000000000000000000000000000002a300"),
		BlockNumber: 19,
		TxHash:      common.HexToHash("0xd8fd67b37a6aa64a3cef4937204765183b180d8dc92eecd0d233f445526d31b5"),
		TxIndex:     flopTxIndex,
		BlockHash:   fakes.FakeHash,
		Index:       32,
		Removed:     false,
	}

	flopTxIndex       = uint(33)
	flopBidId         = int64(5)
	flopLot           = int64(15)
	flopBid           = int64(2)
	flopGal           = "0x7d7bEe5fCfD8028cf7b00876C5b1421c800561A6"
	rawFlopLogJson, _ = json.Marshal(FlopKickLog)
	flopEnd           = int64(172800)

	FlopKickEntity = flop_kick.Entity{
		Id:               big.NewInt(flopBidId),
		Lot:              big.NewInt(flopLot),
		Bid:              big.NewInt(flopBid),
		Gal:              common.HexToAddress(flopGal),
		End:              big.NewInt(flopEnd),
		TransactionIndex: flopTxIndex,
		LogIndex:         FlopKickLog.Index,
		Raw:              FlopKickLog,
	}

	FlopKickModel = flop_kick.Model{
		BidId:            strconv.FormatInt(flopBidId, 10),
		Lot:              strconv.FormatInt(flopLot, 10),
		Bid:              strconv.FormatInt(flopBid, 10),
		Gal:              flopGal,
		End:              time.Unix(flopEnd, 0),
		TransactionIndex: flopTxIndex,
		LogIndex:         FlopKickLog.Index,
		Raw:              rawFlopLogJson,
	}
)

type FlopKickDBResult struct {
	Id       int64
	HeaderId int64 `db:"header_id"`
	flop_kick.Model
}