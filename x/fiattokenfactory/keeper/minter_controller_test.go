package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/wfblockchain/noble-fiattokenfactory/testutil/keeper"
	"github.com/wfblockchain/noble-fiattokenfactory/testutil/nullify"
	"github.com/wfblockchain/noble-fiattokenfactory/x/fiattokenfactory/keeper"
	"github.com/wfblockchain/noble-fiattokenfactory/x/fiattokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMinterController(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MinterController {
	items := make([]types.MinterController, n)
	for i := range items {
		items[i].Controller = strconv.Itoa(i)

		keeper.SetMinterController(ctx, items[i])
	}
	return items
}

func TestMinterControllerGet(t *testing.T) {
	keeper, ctx := keepertest.FiatTokenfactoryKeeper(t)
	items := createNMinterController(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMinterController(ctx,
			item.Controller,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestMinterControllerRemove(t *testing.T) {
	keeper, ctx := keepertest.FiatTokenfactoryKeeper(t)
	items := createNMinterController(keeper, ctx, 10)
	for _, item := range items {
		keeper.DeleteMinterController(ctx,
			item.Minter,
		)
		_, found := keeper.GetMinterController(ctx,
			item.Minter,
		)
		require.False(t, found)
	}
}

func TestMinterControllerGetAll(t *testing.T) {
	keeper, ctx := keepertest.FiatTokenfactoryKeeper(t)
	items := createNMinterController(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMinterControllers(ctx)),
	)
}
