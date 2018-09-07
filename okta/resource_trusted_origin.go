package okta

import (
  "fmt"
  "log"
  "github.com/articulate/oktasdk-go/okta"
  "github.com/hashicorp/terraform/helper/schema"
)

func resourceTrustedOrigin() *schema.Resource {
  return &schema.Resource{
    Create: resourceTrustedOriginCreate,
    Read:   resourceTrustedOriginRead,
    Update: resourceTrustedOriginUpdate,
    Delete: resourceTrustedOriginDelete,
    Exists: trustedOriginExists,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },

    Schema: map[string]*schema.Schema{
      "active": &schema.Schema{
        Type:        schema.TypeBool,
        Optional:    true,
        Default:     true,
        Description: "Whether the Trusted Origin is active or not - can only be issued post-creation",
      },
      "name": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        Description: "Name of the Trusted Origin Resource",
      },
      "origin": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        Description: "The origin to trust",
      },
      "scopes": &schema.Schema{
        Type:        schema.TypeList,
        Optional:    true,
        Elem:        &schema.Schema{Type: schema.TypeString},
        Description: "Scopes of the Trusted Origin - can either be CORS or REDIRECT only",
      },
    },
  }
}

func assembleTrustedOrigin() *okta.TrustedOrigin {
  deactivate := &okta.TrustedOriginDeactive{
    Hints: &okta.TrustedOriginHints{},
  }

  self := &okta.TrustedOriginSelf{
    Hints: &okta.TrustedOriginHints{},
  }

  links := &okta.TrustedOriginLinks{
    Self: self,
    Deactivate: deactivate,
  }

  trustedOrigin := &okta.TrustedOrigin{
    Links: links,
  }

  return trustedOrigin
}

// Populates the Trusted Origin struct (used by the Okta SDK for API operaations) with the data resource provided by TF
func populateTrustedOrigin(trustedOrigin *okta.TrustedOrigin, d *schema.ResourceData) *okta.TrustedOrigin {
  trustedOrigin.Name = d.Get("name").(string)
  trustedOrigin.Origin = d.Get("origin").(string)

  var scopes []map[string]string

  for _, vals := range d.Get("scopes").([]interface{}) {
    scopes = append(scopes, map[string]string{"type": vals.(string)})
  }

  trustedOrigin.Scopes = scopes

  return trustedOrigin
}

func resourceTrustedOriginCreate(d *schema.ResourceData, m interface{}) error {
  log.Printf("[INFO] Create Trusted Origin %v", d.Get("name").(string))

  if !d.Get("active").(bool) {
    return fmt.Errorf("[ERROR] Okta will not allow a Trusted Origin to be created as INACTIVE. Can set to false for existing Trusted Origins only.")
  }

  client := m.(*Config).oktaClient
  trustedOrigin := assembleTrustedOrigin()
  populateTrustedOrigin(trustedOrigin, d)

  returnedTrustedOrigin, _, err := client.TrustedOrigins.CreateTrustedOrigin(trustedOrigin)

  if err != nil {
    return fmt.Errorf("[ERROR] %v.", err)
  }

  d.SetId(returnedTrustedOrigin.ID)

  return resourceTrustedOriginRead(d, m)
}

func resourceTrustedOriginRead(d *schema.ResourceData, m interface{}) error {
  log.Printf("[INFO] Read Trusted Origin %v", d.Get("name").(string))

  var trustedOrigin *okta.TrustedOrigin

  client := m.(*Config).oktaClient

  exists, err := trustedOriginExists(d, m)
  if err != nil {
    return err
  }

  if exists == true {
    trustedOrigin, _, err = client.TrustedOrigins.GetTrustedOrigin(d.Id())
  } else {
    d.SetId("")
    return nil
  }

  d.Set("active", trustedOrigin.Status == "ACTIVE")
  d.Set("origin", trustedOrigin.Origin)
  d.Set("name", trustedOrigin.Name)

  return nil
}

func resourceTrustedOriginUpdate(d *schema.ResourceData, m interface{}) error {
  log.Printf("[INFO] Update Trusted Origin %v", d.Get("name").(string))

  var trustedOrigin = assembleTrustedOrigin()
  populateTrustedOrigin(trustedOrigin, d)

  client := m.(*Config).oktaClient

  exists, err := trustedOriginExists(d, m)
  if err != nil {
    return err
  }

  if exists == true {
    _, _, err = client.TrustedOrigins.UpdateTrustedOrigin(d.Id(), trustedOrigin)
    if err != nil {
      return fmt.Errorf("[ERROR] Error Updating Trusted Origin with Okta: %v", err)
    }
  } else {
    d.SetId("")
    return nil
  }

  return resourceTrustedOriginRead(d, m)
}

func resourceTrustedOriginDelete(d *schema.ResourceData, m interface{}) error {
  log.Printf("[INFO] Delete Trusted Origin %v", d.Get("name").(string))

  client := m.(*Config).oktaClient

  exists, err := trustedOriginExists(d, m)
  if err != nil {
    return err
  }
  if exists == true {
    _, err = client.TrustedOrigins.DeleteTrustedOrigin(d.Id())
    if err != nil {
      return fmt.Errorf("[ERROR] Error Deleting Trusted Origin from Okta: %v", err)
    }
  }

  return nil
}

// check if Trusted Origin exists in Okta
func trustedOriginExists(d *schema.ResourceData, m interface{}) (bool, error) {
  client := m.(*Config).oktaClient
  _, _, err := client.TrustedOrigins.GetTrustedOrigin(d.Id())

  if client.OktaErrorCode == "E0000007" {
    return false, nil
  }
  if err != nil {
    return false, fmt.Errorf("[ERROR] Error Getting Trusted Origin in Okta: %v", err)
  }
  return true, nil
}
