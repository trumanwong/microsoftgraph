package microsoftgraph

import (
	"context"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

// GetDrive Get a user's OneDrive
func (c *Client) GetDrive(requestConfiguration *users.ItemDriveRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
	return c.client.Me().Drive().Get(context.Background(), requestConfiguration)
}

// ListDrives List the current user's drives
func (c *Client) ListDrives(requestConfiguration *users.ItemDrivesRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveCollectionResponseable, error) {
	return c.client.Me().Drives().Get(context.Background(), requestConfiguration)
}

// ListChildrenByItemId List children of a DriveItem with a known ID
// If successful, this method returns the list of items in the children collection of the target item. The children collection will be composed of driveItem resources.
func (c *Client) ListChildrenByItemId(driveId, itemId string, requestConfiguration *drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable, error) {
	return c.client.Drives().ByDriveId(driveId).Items().ByDriveItemId(itemId).Children().Get(context.Background(), requestConfiguration)
}

// RecentFiles List recent files
func (c *Client) RecentFiles(driveId string, requestConfiguration *drives.ItemRecentRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemCollectionResponseable, error) {
	return c.client.Drives().ByDriveId(driveId).Recent().Get(context.Background(), requestConfiguration)
}

// SharedWithMe Get driveItems shared with me
func (c *Client) SharedWithMe(driveId string, requestConfiguration *drives.ItemSharedWithMeRequestBuilderGetRequestConfiguration) (drives.ItemSharedWithMeResponseable, error) {
	return c.client.Drives().ByDriveId(driveId).SharedWithMe().Get(context.Background(), requestConfiguration)
}
