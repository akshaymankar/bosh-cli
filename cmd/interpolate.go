package cmd

import (
	"github.com/akshaymankar/int-yaml/template"
	appui "github.com/akshaymankar/int-yaml/ui"
	"github.com/cppforlife/go-patch/patch"
)

type InterpolateCmd struct {
	ui appui.UI
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

func NewInterpolateCmd(ui appui.UI) InterpolateCmd {
	return InterpolateCmd{ui: ui}
}

func (c InterpolateCmd) Run(opts InterpolateOpts) error {
	tpl := template.NewTemplate(opts.Args.Manifest.Bytes)

	vars := opts.VarFlags.AsVariables()
	op := opts.OpsFlags.AsOp()
	evalOpts := template.EvaluateOpts{
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
