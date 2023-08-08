package microsoftgraph

import (
	"context"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func (c *Client) CreateFolder(driveId, driveItemId, name string, requestConfiguration *drives.ItemItemsItemChildrenRequestBuilderPostRequestConfiguration) (graphmodels.DriveItemable, error) {
	requestBody := graphmodels.NewDriveItem()
	requestBody.SetName(&name)
	folder := graphmodels.NewFolder()
	requestBody.SetFolder(folder)
	additionalData := map[string]interface{}{
		"microsoftGraphConflictBehavior": "rename",
	}
	requestBody.SetAdditionalData(additionalData)

	return c.client.Drives().ByDriveId(driveId).Items().ByDriveItemId(driveItemId).Children().Post(context.Background(), requestBody, nil)
}
