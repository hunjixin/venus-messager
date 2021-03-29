package repo

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/ipfs-force-community/venus-messager/types"
)

type AddressRepo interface {
	HasAddress(ctx context.Context, addr address.Address) (bool, error)
	SaveAddress(ctx context.Context, address *types.Address) error
	UpdateNonce(ctx context.Context, addr address.Address, nonce uint64) error
	UpdateAddressState(ctx context.Context, addr address.Address, state types.AddressState) error
	GetAddress(ctx context.Context, addr address.Address) (*types.Address, error)
	ListAddress(ctx context.Context) ([]*types.Address, error)
	DelAddress(ctx context.Context, addr address.Address) error
	UpdateAddress(ctx context.Context, addr *types.Address) error
	SetSelectMsgNum(ctx context.Context, addr address.Address, num uint64) error
}