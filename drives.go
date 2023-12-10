package microsoftgraph

import (
	"context"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

// GetDrive retrieves a user's OneDrive.
// ctx: The context to use for the request.
// requestConfiguration: The configuration for the request.
// Returns a Driveable object and an error, if any occurred.
func (c *Client) GetDrive(ctx context.Context, requestConfiguration *users.ItemDriveRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
	return c.client.Me().Drive().Get(ctx, requestConfiguration)
}

// ListDrives retrieves a list of the current user's drives.
// ctx: The context to use for the request.
// requestConfiguration: The configuration for the request.
// Returns a DriveCollectionResponseable object and an error, if any occurred.
func (c *Client) ListDrives(ctx context.Context, requestConfiguration *users.ItemDrivesRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveCollectionResponseable, error) {
	return c.client.Me().Drives().Get(ctx, requestConfiguration)
}

// ListChildrenByItemId retrieves a list of children of a DriveItem with a known ID.
// ctx: The context to use for the request.
// driveId: The ID of the drive.
// itemId: The ID of the item.
// requestConfiguration: The configuration for the request.
// Returns a DriveItemCollectionResponseable object and an error, if any occurred.
func (c *Client) ListChildrenByItemId(ctx context.Context, driveId, itemId string, requestConfiguration *drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable, error) {
	return c.client.Drives().ByDriveId(driveId).Items().ByDriveItemId(itemId).Children().Get(ctx, requestConfiguration)
}

// RecentFiles retrieves a list of recent files.
// ctx: The context to use for the request.
// driveId: The ID of the drive.
// requestConfiguration: The configuration for the request.
// Returns a DriveItemCollectionResponseable object and an error, if any occurred.
func (c *Client) RecentFiles(ctx context.Context, driveId string, requestConfiguration *drives.ItemRecentRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable, error) {
	return c.client.Drives().ByDriveId(driveId).Recent().Get(ctx, requestConfiguration)
}

// SharedWithMe retrieves driveItems shared with the current user.
// ctx: The context to use for the request.
// driveId: The ID of the drive.
// requestConfiguration: The configuration for the request.
// Returns an ItemSharedWithMeResponseable object and an error, if any occurred.
func (c *Client) SharedWithMe(ctx context.Context, driveId string, requestConfiguration *drives.ItemSharedWithMeRequestBuilderGetRequestConfiguration) (drives.ItemSharedWithMeResponseable, error) {
	return c.client.Drives().ByDriveId(driveId).SharedWithMe().Get(ctx, requestConfiguration)
}
