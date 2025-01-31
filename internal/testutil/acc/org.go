package acc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/common/conversion"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/config"
)

func CheckDestroyOrgInvitation(s *terraform.State) error {
	conn := TestAccProviderSdkV2.Meta().(*config.MongoDBClient).Atlas

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mongodbatlas_invitations" {
			continue
		}

		ids := conversion.DecodeStateID(rs.Primary.ID)

		orgID := ids["org_id"]
		invitationID := ids["invitation_id"]

		// Try to find the invitation
		_, _, err := conn.Organizations.Invitation(context.Background(), orgID, invitationID)
		if err == nil {
			return fmt.Errorf("invitation (%s) still exists", invitationID)
		}
	}

	return nil
}
