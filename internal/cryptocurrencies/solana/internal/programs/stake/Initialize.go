// Copyright 2024 github.com/cordialsys
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

package stake

import (
	"errors"
	"fmt"

	"xrobot/internal/cryptocurrencies/solana/internal"
	bin "xrobot/internal/cryptocurrencies/solana/internal/binary"

	"xrobot/internal/cryptocurrencies/solana/internal/text/format"

	"github.com/gagliardetto/treeout"
)

type Initialize struct {
	// Authorization settings for stake account
	Authorized *Authorized

	// Lockup settings for stake account
	Lockup *Lockup

	// [0] = [WRITE SIGNER] StakeAccount
	// ··········· Stake account getting initialized
	//
	// [1] = [] RentSysvar
	// ··········· RentSysvar account
	//
	internal.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *Initialize) UnmarshalWithDecoder(dec *bin.Decoder) error {
	{
		err := dec.Decode(&inst.Authorized)
		if err != nil {
			return err
		}
	}
	{
		err := dec.Decode(&inst.Lockup)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *Initialize) MarshalWithEncoder(encoder *bin.Encoder) error {
	{
		err := encoder.Encode(*inst.Authorized)
		if err != nil {
			return err
		}
	}
	{
		err := encoder.Encode(*inst.Lockup)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Authorized == nil {
			return errors.New("authorized parameter is not set")
		}
		err := inst.Authorized.Validate()
		if err != nil {
			return err
		}
	}
	{
		if inst.Lockup == nil {
			return errors.New("lockup parameter is not set")
		}
		err := inst.Lockup.Validate()
		if err != nil {
			return err
		}
	}

	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}

// Stake account account
func (inst *Initialize) SetStakeAccount(stakeAccount internal.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = internal.Meta(stakeAccount).WRITE().SIGNER()
	return inst
}

// Rent sysvar account
func (inst *Initialize) SetRentSysvarAccount(rentSysvar internal.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = internal.Meta(rentSysvar)
	return inst
}
func (inst *Initialize) GetStakeAccount() *internal.AccountMeta      { return inst.AccountMetaSlice[0] }
func (inst *Initialize) GetRentSysvarAccount() *internal.AccountMeta { return inst.AccountMetaSlice[1] }

func (inst *Initialize) SetStaker(staker internal.PublicKey) *Initialize {
	inst.Authorized.Staker = &staker
	return inst
}

func (inst *Initialize) SetWithdrawer(withdrawer internal.PublicKey) *Initialize {
	inst.Authorized.Withdrawer = &withdrawer
	return inst
}

func (inst *Initialize) SetLockupTimestamp(unixTimestamp int64) *Initialize {
	inst.Lockup.UnixTimestamp = &unixTimestamp
	return inst
}

func (inst *Initialize) SetLockupEpoch(epoch uint64) *Initialize {
	inst.Lockup.Epoch = &epoch
	return inst
}

func (inst *Initialize) SetCustodian(custodian internal.PublicKey) *Initialize {
	inst.Lockup.Custodian = &custodian
	return inst
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.TypeIDFromUint32(Instruction_Initialize, bin.LE),
	}}
}

func (inst *Initialize) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child("Authorized").ParentFunc(func(authBranch treeout.Branches) {
							authBranch.Child(format.Account("    Staker", *inst.Authorized.Staker))
							authBranch.Child(format.Account("Withdrawer", *inst.Authorized.Withdrawer))
						})
						paramsBranch.Child("Lockup").ParentFunc(func(authBranch treeout.Branches) {
							authBranch.Child(format.Param("UnixTimestamp", inst.Lockup.UnixTimestamp))
							authBranch.Child(format.Param("        Epoch", inst.Lockup.Epoch))
							authBranch.Child(format.Account("    Custodian", *inst.Lockup.Custodian))
						})
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("           StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(format.Meta("           RentSysvar", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(internal.AccountMetaSlice, 2),
		Authorized:       &Authorized{},
		Lockup:           &Lockup{},
	}
	return nd
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// parameters:
	staker internal.PublicKey,
	withdrawer internal.PublicKey,
	// Accounts:
	stakeAccount internal.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetRentSysvarAccount(internal.SysVarRentPubkey).
		SetStaker(staker).
		SetWithdrawer(withdrawer).
		SetLockupEpoch(0).
		SetLockupTimestamp(0).
		SetCustodian(internal.SystemProgramID)
}
