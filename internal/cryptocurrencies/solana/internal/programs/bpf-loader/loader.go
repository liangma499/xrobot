package bpfloader

import (
	"encoding/binary"
	"fmt"

	"tron_robot/internal/cryptocurrencies/solana/internal"
	"tron_robot/internal/cryptocurrencies/solana/internal/programs/system"
	"tron_robot/internal/cryptocurrencies/solana/internal/rpc"
)

const (
	PACKET_DATA_SIZE int = 1280 - 40 - 8
)

// https://github.com/solana-labs/solana/blob/v1.7.15/cli/src/program.rs#L1683
func calculateMaxChunkSize(
	createBuilder func(offset int, data []byte) *internal.TransactionBuilder,
) (size int, err error) {
	transaction, err := createBuilder(0, []byte{}).Build()
	if err != nil {
		return
	}
	signatures := make(
		[]internal.Signature,
		transaction.Message.Header.NumRequiredSignatures,
	)
	transaction.Signatures = append(transaction.Signatures, signatures...)
	serialized, err := transaction.MarshalBinary()
	if err != nil {
		return
	}
	size = PACKET_DATA_SIZE - len(serialized) - 1
	return
}

// https://github.com/solana-labs/solana/blob/v1.7.15/cli/src/program.rs#L2006
func completePartialProgramInit(
	loaderId internal.PublicKey,
	payerPubkey internal.PublicKey,
	elfPubkey internal.PublicKey,
	account *rpc.Account,
	accountDataLen int,
	minimumBalance uint64,
	allowExcessiveBalance bool,
) (instructions []internal.Instruction, balanceNeeded uint64, err error) {
	if account.Executable {
		err = fmt.Errorf("buffer account is already executable")
		return
	}
	if !account.Owner.Equals(loaderId) &&
		!account.Owner.Equals(internal.SystemProgramID) {
		err = fmt.Errorf(
			"buffer account passed is already in use by another program",
		)
		return
	}
	if len(account.Data.GetBinary()) > 0 &&
		len(account.Data.GetBinary()) < accountDataLen {
		err = fmt.Errorf(
			"buffer account passed is not large enough, may have been for a " +
				" different deploy?",
		)
		return
	}

	if len(account.Data.GetBinary()) == 0 &&
		account.Owner.Equals(internal.SystemProgramID) {
		instructions = append(
			instructions,
			system.NewAllocateInstruction(uint64(accountDataLen), elfPubkey).
				Build(),
		)
		instructions = append(
			instructions,
			system.NewAssignInstruction(loaderId, elfPubkey).Build(),
		)
		if account.Lamports < minimumBalance {
			balance := minimumBalance - account.Lamports
			instructions = append(
				instructions,
				system.NewTransferInstruction(balance, payerPubkey, elfPubkey).
					Build(),
			)
			balanceNeeded = balance
		} else if account.Lamports > minimumBalance &&
			account.Owner.Equals(internal.SystemProgramID) &&
			!allowExcessiveBalance {
			err = fmt.Errorf(
				"buffer account has a balance: %v.%v; it may already be in use",
				account.Lamports/internal.LAMPORTS_PER_SOL,
				account.Lamports%internal.LAMPORTS_PER_SOL,
			)
			return
		}
	}
	return
}

func load(
	payerPubkey internal.PublicKey,
	account *rpc.Account,
	programData []byte,
	bufferDataLen int,
	minimumBalance uint64,
	loaderId internal.PublicKey,
	bufferPubkey internal.PublicKey,
	allowExcessiveBalance bool,
) (
	initialBuilder *internal.TransactionBuilder,
	writeBuilders []*internal.TransactionBuilder,
	finalBuilder *internal.TransactionBuilder,
	balanceNeeded uint64,
	err error,
) {
	var instructions []internal.Instruction
	if account != nil {
		instructions, balanceNeeded, err = completePartialProgramInit(
			loaderId,
			payerPubkey,
			bufferPubkey,
			account,
			bufferDataLen,
			minimumBalance,
			allowExcessiveBalance,
		)
		if err != nil {
			return
		}
	} else {
		instructions = append(
			instructions,
			system.NewCreateAccountInstruction(
				minimumBalance,
				uint64(bufferDataLen),
				loaderId,
				payerPubkey,
				bufferPubkey,
			).Build(),
		)
		balanceNeeded = minimumBalance
	}
	if len(instructions) > 0 {
		initialBuilder = internal.NewTransactionBuilder().SetFeePayer(payerPubkey)
		for _, instruction := range instructions {
			initialBuilder = initialBuilder.AddInstruction(instruction)
		}
	}

	createBuilder := func(offset int, chunk []byte) *internal.TransactionBuilder {
		data := make([]byte, len(chunk)+16)
		binary.LittleEndian.PutUint32(data[0:], 0)
		binary.LittleEndian.PutUint32(data[4:], uint32(offset))
		binary.LittleEndian.PutUint32(data[8:], uint32(len(chunk)))
		binary.LittleEndian.PutUint32(data[12:], 0)
		copy(data[16:], chunk)
		instruction := internal.NewInstruction(
			loaderId,
			internal.AccountMetaSlice{
				internal.NewAccountMeta(bufferPubkey, true, true),
			},
			data,
		)
		return internal.NewTransactionBuilder().
			AddInstruction(instruction).
			SetFeePayer(payerPubkey)
	}

	chunkSize, err := calculateMaxChunkSize(createBuilder)
	if err != nil {
		return
	}
	writeBuilders = []*internal.TransactionBuilder{}
	for i := 0; i < len(programData); i += chunkSize {
		end := i + chunkSize
		if end > len(programData) {
			end = len(programData)
		}
		writeBuilders = append(
			writeBuilders,
			createBuilder(i, programData[i:end]),
		)
	}

	finalBuilder = internal.NewTransactionBuilder().SetFeePayer(payerPubkey)
	{
		data := make([]byte, 4)
		binary.LittleEndian.PutUint32(data[0:], 1)
		instruction := internal.NewInstruction(
			loaderId,
			internal.AccountMetaSlice{
				internal.NewAccountMeta(bufferPubkey, true, true),
			},
			data,
		)
		finalBuilder.AddInstruction(instruction)
	}
	return
}

func Deploy(
	payerPubkey internal.PublicKey,
	account *rpc.Account,
	programData []byte,
	minimumBalance uint64,
	loaderId internal.PublicKey,
	bufferPubkey internal.PublicKey,
	allowExcessiveBalance bool,
) (
	initialBuilder *internal.TransactionBuilder,
	writeBuilders []*internal.TransactionBuilder,
	finalBuilder *internal.TransactionBuilder,
	balanceNeeded uint64,
	err error,
) {
	return load(
		payerPubkey,
		account,
		programData,
		len(programData),
		minimumBalance,
		loaderId,
		bufferPubkey,
		allowExcessiveBalance,
	)
}
