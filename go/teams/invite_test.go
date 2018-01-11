package teams

import (
	"context"
	"testing"

	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/stretchr/testify/require"
)

func TestObsoletingInvites1(t *testing.T) {
	// This chain has 3 keybase invites total:
	// 1) 579651b0d574971040b531b66efbc519%1
	// 2) 618d663af0f1ec88a5a19defa65a2f19%1
	// 3) 40903c59d19feef1d67c455499304c19%1
	//
	// 1 gets obsoleted by "change_membership" link that adds the same
	// person but does not complete the invite. 2 is canceled by
	// "invite" link. 3 should be still active when the chain is done
	// replaying.
	team := runUnitFromFilename(t, "invite_obsolete.json")

	hasInvite, err := team.HasActiveInvite(keybase1.TeamInviteName("579651b0d574971040b531b66efbc519%1"), "keybase")
	require.NoError(t, err)
	require.False(t, hasInvite)

	hasInvite, err = team.HasActiveInvite(keybase1.TeamInviteName("618d663af0f1ec88a5a19defa65a2f19%1"), "keybase")
	require.NoError(t, err)
	require.False(t, hasInvite)

	hasInvite, err = team.HasActiveInvite(keybase1.TeamInviteName("40903c59d19feef1d67c455499304c19%1"), "keybase")
	require.NoError(t, err)
	require.True(t, hasInvite)

	activeInvites := team.chain().inner.ActiveInvites
	require.Equal(t, 1, len(activeInvites))

	for _, invite := range activeInvites {
		require.Equal(t, keybase1.TeamRole_READER, invite.Role)
		require.EqualValues(t, "56eafff3400b5bcd8b40bff3d225ab27", invite.Id)
		require.EqualValues(t, "40903c59d19feef1d67c455499304c19%1", invite.Name)
		require.EqualValues(t, keybase1.UserVersion{Uid: "25852c87d6e47fb8d7d55400be9c7a19", EldestSeqno: 1}, invite.Inviter)
	}
}

func TestObsoletingInvites2(t *testing.T) {
	// This chain is a backwards-compatibility test to see if even if
	// someone got tricked into accepting obsolete invite, such chain
	// should still play and result in predictable end state.
	team := runUnitFromFilename(t, "invite_obsolete_trick.json")
	require.Equal(t, 0, len(team.chain().inner.ActiveInvites))
	require.True(t, team.IsMember(context.Background(), keybase1.UserVersion{Uid: "579651b0d574971040b531b66efbc519", EldestSeqno: 1}))
}
