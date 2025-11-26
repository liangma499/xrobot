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

package associatedtokenaccount

import (
	"fmt"

	"xrobot/internal/cryptocurrencies/solana/internal"
	bin "xrobot/internal/cryptocurrencies/solana/internal/binary"

	text "xrobot/internal/cryptocurrencies/solana/internal/text"

	spew "github.com/davecgh/go-spew/spew"
	treeout "github.com/gagliardetto/treeout"
)

var ProgramID internal.PublicKey = internal.SPLAssociatedTokenAccountProgramID

func SetProgramID(pubkey internal.PublicKey) {
	ProgramID = pubkey
	internal.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "AssociatedTokenAccount"

func init() {
	internal.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

type Instruction struct {
	bin.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent treeout.Branches) {
	if enToTree, ok := inst.Impl.(text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(spew.Sdump(inst))
	}
}

var InstructionImplDef = bin.NewVariantDefinition(
	bin.NoTypeIDEncoding, // NOTE: the associated-token-account program has no ID encoding.
	[]bin.VariantType{
		{
			"Create", (*Create)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() internal.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*internal.AccountMeta) {
	return inst.Impl.(internal.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	return []byte{}, nil
}

func (inst *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *bin.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst Instruction) MarshalWithEncoder(encoder *bin.Encoder) error {
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*internal.AccountMeta, data []byte) (any, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*internal.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := bin.NewBinDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(internal.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
