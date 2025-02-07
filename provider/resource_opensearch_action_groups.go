package provider 

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/olivere/elastic/uritemplates"
	elastic7 "github.com/olivere/elastic/v7"
)

var actionGroupsShema = map [string]*schema.Schema{
		"body": {

		}
}

func resourceOpensearchActionGroups() {}

func resourceOpensearchActionGroupsCreate() {}

func resourceOpensearchActionGroupsRead()  {}

func resourceOpensearchActionGroupsUpdate() {}

func resourceOpensearchActionGroupsGet() {}

func resourceOpensearchPostActionGroups() {}

func resourceOpensearchPutActionGroups() {}

func resourceOpensearchActionGroupsDelete() {}

type actionGroupsResponse struct {}