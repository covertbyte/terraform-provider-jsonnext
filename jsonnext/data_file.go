package jsonnext

import (
	"github.com/google/go-jsonnet"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
)

func dataFile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Jsonnet source file",
			},

			"output": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON output",
			},
		},
		Read: dataFileRead,
	}
}

func dataFileRead(d *schema.ResourceData, meta interface{}) error {
	vm := jsonnet.MakeVM()
	vm.Importer(&meta.(*config).importer)

	source := d.Get("source").(string)

	snippet, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	output, err := vm.EvaluateSnippet(source, string(snippet))
	if err != nil {
		return err
	}

	if err := d.Set("output", output); err != nil {
		return err
	}

	d.SetId(hash(output))

	return nil
}
