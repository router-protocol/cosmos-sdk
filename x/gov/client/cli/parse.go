package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
)

func parseSubmitProposalFlags(fs *pflag.FlagSet) (*proposal, error) {
	proposal := &proposal{}
	proposalFile, _ := fs.GetString(FlagProposal)

	if proposalFile == "" {
		proposalType, _ := fs.GetString(FlagProposalType)

		proposal.Title, _ = fs.GetString(FlagTitle)
		proposal.Description, _ = fs.GetString(FlagDescription)
		proposal.Type = govutils.NormalizeProposalType(proposalType)
		proposal.Deposit, _ = fs.GetString(FlagDeposit)
		return proposal, nil
	}

	for _, flag := range ProposalFlags {
		if v, _ := fs.GetString(flag); v != "" {
			return nil, fmt.Errorf("--%s flag provided alongside --proposal, which is a noop", flag)
		}
	}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, proposal)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func parseProposalMessages(fs *pflag.FlagSet) ([]sdk.Msg, sdk.Coins, error) {
	proposal := &proposal2{}
	proposalFile, _ := fs.GetString(FlagProposal)

	var (
		messagesString string
		depositString  string
		err            error
		msgs           []sdk.Msg
	)

	if proposalFile != "" {
		for _, flag := range ProposalFlags {
			if v, _ := fs.GetString(flag); v != "" {
				return []sdk.Msg{}, nil, fmt.Errorf("--%s flag provided alongside --proposal, which is a noop", flag)
			}
		}

		bz, err := ioutil.ReadFile(proposalFile)
		if err != nil {
			return []sdk.Msg{}, nil, err
		}

		err = json.Unmarshal(bz, proposal)
		if err != nil {
			return []sdk.Msg{}, nil, err
		}

		messagesString = proposal.Messages
		depositString = proposal.Deposit
	} else {
		depositString, err = fs.GetString(FlagDeposit)
		if err != nil {
			return []sdk.Msg{}, nil, err
		}

		messagesString, err = fs.GetString(FlagMessages)
		if err != nil {
			return []sdk.Msg{}, nil, err
		}

	}

	deposit, err := sdk.ParseCoinsNormalized(depositString)
	if err != nil {
		return []sdk.Msg{}, nil, err
	}

	json.Unmarshal([]byte(messagesString), msgs)

	return msgs, deposit, nil
}
