package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// SubmitProposal create new proposal given a content
func (keeper Keeper) SubmitProposal(
	ctx sdk.Context,
	content types.Content,
) (types.Proposal, error) {
	if !keeper.handler.HasRoute(content.ProposalRoute()) {
		return types.Proposal{}, sdkerrors.Wrap(types.ErrNoProposalHandlerExists, content.ProposalRoute())
	}

	// Execute the proposal content in a new context branch (with branched store)
	// to validate the actual parameter changes before the proposal proceeds
	// through the governance process. State is not persisted.
	cacheCtx, _ := ctx.CacheContext()
	handler := keeper.handler.GetRoute(content.ProposalRoute())
	if err := handler(cacheCtx, content); err != nil {
		return types.Proposal{}, sdkerrors.Wrap(types.ErrInvalidProposalContent, err.Error())
	}

	proposalID, err := keeper.GetProposalID(ctx)
	if err != nil {
		return types.Proposal{}, err
	}

	submitTime := ctx.BlockHeader().Time
	depositPeriod := keeper.GetDepositParams(ctx).MaxDepositPeriod

	proposal, err := types.NewProposal(content, proposalID, submitTime, submitTime.Add(depositPeriod))
	if err != nil {
		return types.Proposal{}, err
	}

	keeper.SetProposal(ctx, proposal)
	keeper.InsertInactiveProposalQueue(ctx, proposalID, proposal.DepositEndTime)
	keeper.SetProposalID(ctx, proposalID+1)

	// called right after a proposal is submitted
	keeper.AfterProposalSubmission(ctx, proposalID)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitProposal,
			sdk.NewAttribute(types.AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
		),
	)

	return proposal, nil
}

// SubmitProposalV2 create new V2 proposal given an array of messages
func (keeper Keeper) SubmitProposalV2(
	ctx sdk.Context,
	messages []sdk.Msg,
) (types.ProposalV2, error) {

	// Loop through all messages and confirm that each has a handler and the gov module account
	// as the only signer
	for _, msg := range messages {
		signers := msg.GetSigners()
		if len(signers) != 1 {
			return types.ProposalV2{}, types.ErrInvalidSigner
		}

		if !signers[0].Equals(keeper.GetGovernanceAccount(ctx).GetAddress()) {
			return types.ProposalV2{}, sdkerrors.Wrap(types.ErrInvalidSigner, signers[0].String())
		}

		if keeper.router.Handler(msg) == nil {
			return types.ProposalV2{}, sdkerrors.Wrap(types.ErrUnroutableProposalMsg, sdk.MsgTypeURL(msg))
		}
	}

	proposalID, err := keeper.GetProposalID(ctx)
	if err != nil {
		return types.ProposalV2{}, err
	}

	submitTime := ctx.BlockHeader().Time
	depositPeriod := keeper.GetDepositParams(ctx).MaxDepositPeriod

	proposal, err := types.NewProposalV2(messages, proposalID, submitTime, submitTime.Add(depositPeriod))
	if err != nil {
		return types.ProposalV2{}, err
	}

	keeper.SetProposalV2(ctx, proposal)
	keeper.InsertInactiveProposalQueue(ctx, proposalID, proposal.DepositEndTime)
	keeper.SetProposalID(ctx, proposalID+1)

	// called right after a proposal is submitted
	keeper.AfterProposalSubmission(ctx, proposalID)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitProposal,
			sdk.NewAttribute(types.AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
		),
	)

	return proposal, nil
}

func (keeper Keeper) GetProposalVersion(ctx sdk.Context, proposalID uint64) int {
	store := ctx.KVStore(keeper.storeKey)

	switch {
	case store.Has(types.ProposalKey(proposalID)):
		return types.Version1
	case store.Has(types.ProposalKeyV2(proposalID)):
		return types.Version2
	default:
		return 0
	}
}

// GetProposal get proposal from store by ProposalID
func (keeper Keeper) GetProposal(ctx sdk.Context, proposalID uint64) (types.Proposal, bool) {
	store := ctx.KVStore(keeper.storeKey)

	bz := store.Get(types.ProposalKey(proposalID))
	if bz == nil {
		return types.Proposal{}, false
	}

	var proposal types.Proposal
	keeper.MustUnmarshalProposal(bz, &proposal)

	return proposal, true
}

// GetProposalV2 get proposal from store by ProposalID
func (keeper Keeper) GetProposalV2(ctx sdk.Context, proposalID uint64) (types.ProposalV2, bool) {
	store := ctx.KVStore(keeper.storeKey)

	bz := store.Get(types.ProposalKeyV2(proposalID))
	if bz == nil {
		return types.ProposalV2{}, false
	}

	var proposal types.ProposalV2
	keeper.cdc.MustUnmarshal(bz, &proposal)

	return proposal, true
}

// SetProposal set a proposal to store
func (keeper Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := ctx.KVStore(keeper.storeKey)

	bz := keeper.MustMarshalProposal(proposal)

	store.Set(types.ProposalKey(proposal.ProposalId), bz)
}

// SetProposalV2 persists a proposal to store
func (keeper Keeper) SetProposalV2(ctx sdk.Context, proposal types.ProposalV2) {
	store := ctx.KVStore(keeper.storeKey)

	bz := keeper.cdc.MustMarshal(&proposal)

	store.Set(types.ProposalKeyV2(proposal.ProposalId), bz)
}

// DeleteProposal deletes a proposal from store (both V1 and V2)
func (keeper Keeper) DeleteProposal(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		panic(fmt.Sprintf("couldn't find proposal with id#%d", proposalID))
	}
	keeper.RemoveFromInactiveProposalQueue(ctx, proposalID, proposal.DepositEndTime)
	keeper.RemoveFromActiveProposalQueue(ctx, proposalID, proposal.VotingEndTime)
	store.Delete(types.ProposalKey(proposalID))
	store.Delete(types.ProposalKeyV2(proposalID))
}

// IterateProposals iterates over the all the proposals and performs a callback function
func (keeper Keeper) IterateProposals(ctx sdk.Context, cb func(proposal types.Proposal) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ProposalsKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var proposal types.Proposal
		err := keeper.UnmarshalProposal(iterator.Value(), &proposal)
		if err != nil {
			panic(err)
		}

		if cb(proposal) {
			break
		}
	}
}

// IterateProposals iterates over the all the proposals and performs a callback function
func (keeper Keeper) IterateProposalsV2(ctx sdk.Context, cb func(proposal types.ProposalV2) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ProposalsKeyPrefixV2)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var proposal types.ProposalV2
		err := keeper.cdc.Unmarshal(iterator.Value(), &proposal)
		if err != nil {
			panic(err)
		}

		if cb(proposal) {
			break
		}
	}
}

// GetProposals returns all the proposals from store
func (keeper Keeper) GetProposals(ctx sdk.Context) (proposals types.Proposals) {
	keeper.IterateProposals(ctx, func(proposal types.Proposal) bool {
		proposals = append(proposals, proposal)
		return false
	})
	return
}

// GetProposals returns all the proposals from store
func (keeper Keeper) GetProposalsV2(ctx sdk.Context) (proposals types.ProposalsV2) {
	keeper.IterateProposalsV2(ctx, func(proposal types.ProposalV2) bool {
		proposals = append(proposals, proposal)
		return false
	})
	return
}

// GetProposalsFiltered retrieves proposals filtered by a given set of params which
// include pagination parameters along with voter and depositor addresses and a
// proposal status. The voter address will filter proposals by whether or not
// that address has voted on proposals. The depositor address will filter proposals
// by whether or not that address has deposited to them. Finally, status will filter
// proposals by status.
//
// NOTE: If no filters are provided, all proposals will be returned in paginated
// form.
func (keeper Keeper) GetProposalsFiltered(ctx sdk.Context, params types.QueryProposalsParams) types.Proposals {
	proposals := keeper.GetProposals(ctx)
	filteredProposals := make([]types.Proposal, 0, len(proposals))

	for _, p := range proposals {
		matchVoter, matchDepositor, matchStatus := true, true, true

		// match status (if supplied/valid)
		if types.ValidProposalStatus(params.ProposalStatus) {
			matchStatus = p.Status == params.ProposalStatus
		}

		// match voter address (if supplied)
		if len(params.Voter) > 0 {
			_, matchVoter = keeper.GetVote(ctx, p.ProposalId, params.Voter)
		}

		// match depositor (if supplied)
		if len(params.Depositor) > 0 {
			_, matchDepositor = keeper.GetDeposit(ctx, p.ProposalId, params.Depositor)
		}

		if matchVoter && matchDepositor && matchStatus {
			filteredProposals = append(filteredProposals, p)
		}
	}

	start, end := client.Paginate(len(filteredProposals), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		filteredProposals = []types.Proposal{}
	} else {
		filteredProposals = filteredProposals[start:end]
	}

	return filteredProposals
}

func (keeper Keeper) GetProposalsFilteredV2(ctx sdk.Context, params types.QueryProposalsParams) types.ProposalsV2 {
	proposals := keeper.GetProposalsV2(ctx)
	filteredProposals := make([]types.ProposalV2, 0, len(proposals))

	for _, p := range proposals {
		matchVoter, matchDepositor, matchStatus := true, true, true

		// match status (if supplied/valid)
		if types.ValidProposalStatus(params.ProposalStatus) {
			matchStatus = p.Status == params.ProposalStatus
		}

		// match voter address (if supplied)
		if len(params.Voter) > 0 {
			_, matchVoter = keeper.GetVote(ctx, p.ProposalId, params.Voter)
		}

		// match depositor (if supplied)
		if len(params.Depositor) > 0 {
			_, matchDepositor = keeper.GetDeposit(ctx, p.ProposalId, params.Depositor)
		}

		if matchVoter && matchDepositor && matchStatus {
			filteredProposals = append(filteredProposals, p)
		}
	}

	start, end := client.Paginate(len(filteredProposals), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		filteredProposals = []types.ProposalV2{}
	} else {
		filteredProposals = filteredProposals[start:end]
	}

	return filteredProposals
}

// GetProposalID gets the highest proposal ID
func (keeper Keeper) GetProposalID(ctx sdk.Context) (proposalID uint64, err error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.ProposalIDKey)
	if bz == nil {
		return 0, sdkerrors.Wrap(types.ErrInvalidGenesis, "initial proposal ID hasn't been set")
	}

	proposalID = types.GetProposalIDFromBytes(bz)
	return proposalID, nil
}

// SetProposalID sets the new proposal ID to the store
func (keeper Keeper) SetProposalID(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Set(types.ProposalIDKey, types.GetProposalIDBytes(proposalID))
}

func (keeper Keeper) ActivateVotingPeriod(ctx sdk.Context, proposal types.Proposal) {
	proposal.VotingStartTime = ctx.BlockHeader().Time
	votingPeriod := keeper.GetVotingParams(ctx).VotingPeriod
	proposal.VotingEndTime = proposal.VotingStartTime.Add(votingPeriod)
	proposal.Status = types.StatusVotingPeriod
	keeper.SetProposal(ctx, proposal)

	keeper.RemoveFromInactiveProposalQueue(ctx, proposal.ProposalId, proposal.DepositEndTime)
	keeper.InsertActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)
}

func (keeper Keeper) ActivateVotingPeriodV2(ctx sdk.Context, proposal types.ProposalV2) {
	proposal.VotingStartTime = ctx.BlockHeader().Time
	votingPeriod := keeper.GetVotingParams(ctx).VotingPeriod
	proposal.VotingEndTime = proposal.VotingStartTime.Add(votingPeriod)
	proposal.Status = types.StatusVotingPeriod
	keeper.SetProposalV2(ctx, proposal)

	keeper.RemoveFromInactiveProposalQueue(ctx, proposal.ProposalId, proposal.DepositEndTime)
	keeper.InsertActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)
}

func (keeper Keeper) MarshalProposal(proposal types.Proposal) ([]byte, error) {
	bz, err := keeper.cdc.Marshal(&proposal)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func (keeper Keeper) UnmarshalProposal(bz []byte, proposal *types.Proposal) error {
	err := keeper.cdc.Unmarshal(bz, proposal)
	if err != nil {
		return err
	}
	return nil
}

func (keeper Keeper) MustMarshalProposal(proposal types.Proposal) []byte {
	bz, err := keeper.MarshalProposal(proposal)
	if err != nil {
		panic(err)
	}
	return bz
}

func (keeper Keeper) MustUnmarshalProposal(bz []byte, proposal *types.Proposal) {
	err := keeper.UnmarshalProposal(bz, proposal)
	if err != nil {
		panic(err)
	}
}
