/**
* @file
* @copyright defined in go-seele/LICENSE
 */

package state

import (
	"math/big"

	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/core/types"
)

// CreateAccount creates a new account in statedb.
func (s *Statedb) CreateAccount(address common.Address) {
	s.GetOrNewStateObject(address)
}

// GetCodeHash returns contract code hash associated with the specified address if any.
// Otherwise, return an empty hash.
func (s *Statedb) GetCodeHash(address common.Address) common.Hash {
	// @todo
	return common.EmptyHash
}

// GetCode returns the contract code associated with the specified address if any.
// Otherwise, return nil.
func (s *Statedb) GetCode(address common.Address) []byte {
	// @todo
	return nil
}

// SetCode sets the contract code of the specified address if exists.
func (s *Statedb) SetCode(address common.Address, code []byte) {
	// @todo
}

// GetCodeSize returns the contract code size associated with the specified address if any.
// Otherwise, return 0.
func (s *Statedb) GetCodeSize(address common.Address) int {
	// @todo
	return 0
}

// AddRefund refunds gas
func (s *Statedb) AddRefund(gas uint64) {
	// @todo
}

// GetRefund returns the current value of the refund counter.
func (s *Statedb) GetRefund() uint64 {
	// @todo
	return 0
}

// GetState returns the value of specified key in account storage if exists.
// Otherwise, return empty hash.
func (s *Statedb) GetState(address common.Address, key common.Hash) common.Hash {
	// @todo
	return common.EmptyHash
}

// SetState adds or updates key-value pair in account storage.
func (s *Statedb) SetState(address common.Address, key common.Hash, value common.Hash) {
	// @todo
}

// Suicide marks the given account as suicided and clears the account balance.
// Note, The account's state object is still available until the state is committed.
// Return true if the specified account exists, otherwise false.
func (s *Statedb) Suicide(address common.Address) bool {
	stateObj := s.getStateObject(address)
	if stateObj == nil {
		return false
	}

	stateObj.SetAmount(new(big.Int))
	// @todo mark the state object as suicided

	return true
}

// HasSuicided returns true if the specified account exists and suicided, otherwise false.
func (s *Statedb) HasSuicided(address common.Address) bool {
	stateObj := s.getStateObject(address)
	if stateObj == nil {
		return false
	}

	// @todo return stateObj.suicided

	return false
}

// Exist reports whether the given account exists in state.
// Notably this should also return true for suicided accounts.
func (s *Statedb) Exist(address common.Address) bool {
	return s.getStateObject(address) != nil
}

// Empty returns whether the given account satisfy (balance = nonce = code = 0).
func (s *Statedb) Empty(address common.Address) bool {
	// @todo
	return false
}

// RevertToSnapshot reverts all state changes made since the given revision.
func (s *Statedb) RevertToSnapshot(revid int) {
	// @todo
}

// Snapshot returns an identifier for the current revision of the state.
func (s *Statedb) Snapshot() int {
	// @todo
	return 0
}

// AddLog add a log.
func (s *Statedb) AddLog(log *types.Log) {
	// @todo
}

// AddPreimage records a SHA3 preimage seen by the VM.
func (s *Statedb) AddPreimage(common.Hash, []byte) {
	// @todo
}

// ForEachStorage visits all the key-value pairs for the specified account storage.
func (s *Statedb) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) {
	// @todo
}