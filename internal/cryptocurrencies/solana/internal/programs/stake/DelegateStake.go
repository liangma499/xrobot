// Copyright 2024 github.com/cordialsys
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stake

import (
	"fmt"

	"xrobot/internal/cryptocurrencies/solana/internal"
	bin "xrobot/internal/cryptocurrencies/solana/internal/binary"

	"xrobot/internal/cryptocurrencies/solana/internal/text/format"

	"github.com/gagliardetto/treeout"
)

type DelegateStake struct {
	// [0] = [WRITE SIGNER] StakeAccount
	// ··········· Stake account getting initialized
	//
	// [1] = [] Vote Account
	// ··········· The validator vote account being delegated to
	//
	// [2] = [] Clock Sysvar
	// ··········· The Clock Sysvar Account
	//
	// [3] = [] Stake History Sysvar
	// ··········· The Stake History Sysvar Account
	//
	// [4] = [] Stake Config Account
	// ··········· The Stake Config Account
	//
	// [5] = [WRITE SIGNER] Stake Authoriy
	// ··········· The Stake Authority
	//
	internal.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *DelegateStake) Validate() error {
	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}
func (inst *DelegateStake) SetStakeAccount(stakeAccount internal.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[0] = internal.Meta(stakeAccount).WRITE().SIGNER()
	return inst
}
func (inst *DelegateStake) SetVoteAccount(voteAcc internal.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[1] = internal.Meta(voteAcc)
	return inst
}
func (inst *DelegateStake) SetClockSysvar(clockSysVarAcc internal.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[2] = internal.Meta(clockSysVarAcc)
	return inst
}
func (inst *DelegateStake) SetStakeHistorySysvar(stakeHistorySysVarAcc internal.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[3] = internal.Meta(stakeHistorySysVarAcc)
	return inst
}
func (inst *DelegateStake) SetConfigAccount(stakeConfigAcc internal.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[4] = internal.Meta(stakeConfigAcc)
	return inst
}
func (inst *DelegateStake) SetStakeAuthority(stakeAuthority internal.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[5] = internal.Meta(stakeAuthority).WRITE().SIGNER()
	return inst
}
func (inst *DelegateStake) GetStakeAccount() *internal.AccountMeta { return inst.AccountMetaSlice[0] }
func (inst *DelegateStake) GetVoteAccount() *internal.AccountMeta  { return inst.AccountMetaSlice[1] }
func (inst *DelegateStake) GetClockSysvar() *internal.AccountMeta  { return inst.AccountMetaSlice[2] }
func (inst *DelegateStake) GetStakeHistorySysvar() *internal.AccountMeta {
	return inst.AccountMetaSlice[3]
}
func (inst *DelegateStake) GetConfigAccount() *internal.AccountMeta  { return inst.AccountMetaSlice[4] }
func (inst *DelegateStake) GetStakeAuthority() *internal.AccountMeta { return inst.AccountMetaSlice[5] }

func (inst DelegateStake) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.TypeIDFromUint32(Instruction_DelegateStake, bin.LE),
	}}
}

func (inst *DelegateStake) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("DelegateStake")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("           StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(format.Meta("           VoteAccount", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(format.Meta("           ClockSysvar", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(format.Meta("           StakeHistorySysvar", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(format.Meta("           StakeConfigAccount", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(format.Meta("           StakeAuthoriy", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

// NewDelegateStakeInstructionBuilder creates a new `DelegateStake` instruction builder.
func NewDelegateStakeInstructionBuilder() *DelegateStake {
	nd := &DelegateStake{
		AccountMetaSlice: make(internal.AccountMetaSlice, 6),
	}
	return nd
}

// NewDelegateStakeInstruction declares a new DelegateStake instruction with the provided parameters and accounts.
func NewDelegateStakeInstruction(
	// Accounts:
	validatorVoteAccount internal.PublicKey,
	stakeAuthority internal.PublicKey,
	stakeAccount internal.PublicKey,
) *DelegateStake {
	return NewDelegateStakeInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetVoteAccount(validatorVoteAccount).
		SetClockSysvar(internal.SysVarClockPubkey).
		SetStakeHistorySysvar(internal.SysVarStakeHistoryPubkey).
		SetConfigAccount(internal.SysVarStakeConfigPubkey).
		SetStakeAuthority(stakeAuthority)
}
