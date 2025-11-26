// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"xrobot/internal/cryptocurrencies/solana/internal"
	"xrobot/internal/cryptocurrencies/solana/internal/rpc"
	"xrobot/internal/cryptocurrencies/solana/internal/rpc/ws"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	ctx := context.Background()
	client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
	if err != nil {
		panic(err)
	}
	program := internal.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA") // token
	defer client.Close()

	sub, err := client.ProgramSubscribeWithOpts(
		program,
		rpc.CommitmentRecent,
		internal.EncodingBase64Zstd,
		nil,
	)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv(ctx)
		if err != nil {
			panic(err)
		}
		spew.Dump(got)

		decodedBinary := got.Value.Account.Data.GetBinary()
		if decodedBinary != nil {
			// spew.Dump(decodedBinary)
		}

		// or if you requested internal.EncodingJSONParsed and it is supported:
		rawJSON := got.Value.Account.Data.GetRawJSON()
		if rawJSON != nil {
			// spew.Dump(rawJSON)
		}
	}
}
