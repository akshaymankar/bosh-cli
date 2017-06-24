package cmd

import (
	"github.com/cppforlife/go-patch/patch"

	boshtpl "github.com/akshaymankar/int-yaml/template"
	boshui "github.com/akshaymankar/int-yaml/ui"
)

type InterpolateCmd struct {
	ui boshui.UI
}

type InterpolateArgs struct {
	Manifest FileBytesArg `positional-arg-name:"PATH" description:"Path to a template that will be interpolated"`
}

type InterpolateOpts struct {
	Args InterpolateArgs `positional-args:"true" required:"true"`

	VarFlags
	OpsFlags

	Path            patch.Pointer `long:"path" value-name:"OP-PATH" description:"Extract value out of template (e.g.: /private_key)"`
	VarErrors       bool          `long:"var-errs"                  description:"Expect all variables to be found, otherwise error"`
	VarErrorsUnused bool          `long:"var-errs-unused"           description:"Expect all variables to be used, otherwise error"`

	command
}

type command struct{}

func NewInterpolateCmd(ui boshui.UI) InterpolateCmd {
	return InterpolateCmd{ui: ui}
}

func (c InterpolateCmd) Run(opts InterpolateOpts) error {
	tpl := boshtpl.NewTemplate(opts.Args.Manifest.Bytes)

	vars := opts.VarFlags.AsVariables()
	op := opts.OpsFlags.AsOp()
	evalOpts := boshtpl.EvaluateOpts{
		ExpectAllKeys:     opts.VarErrors,
		ExpectAllVarsUsed: opts.VarErrorsUnused,
	}

	if opts.Path.IsSet() {
		evalOpts.PostVarSubstitutionOp = patch.FindOp{Path: opts.Path}

		// Printing YAML indented multiline strings (eg SSH key) is not useful
		evalOpts.UnescapedMultiline = true
	}

	bytes, err := tpl.Evaluate(vars, op, evalOpts)
	if err != nil {
		return err
	}
	c.ui.PrintBlock(string(bytes))

	return nil
}
