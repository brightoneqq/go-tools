package cmd

import (
	"errors"
	"fmt"
	"github.com/brightoneqq/go-tools/gstring"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func MinimumArgs(n int, msg string) cobra.PositionalArgs {
	if msg == "" {
		return cobra.MinimumNArgs(1)
	}

	return func(cmd *cobra.Command, args []string) error {
		if len(args) < n {
			return &FlagError{Err: errors.New(msg)}
		}
		return nil
	}
}

func ExactArgs(n int, msg string) cobra.PositionalArgs {

	return func(cmd *cobra.Command, args []string) error {
		if len(args) > n {
			return &FlagError{Err: errors.New("too many arguments")}
		}

		if len(args) < n {
			return &FlagError{Err: errors.New(msg)}
		}

		return nil
	}
}

func NoArgsQuoteReminder(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	errMsg := fmt.Sprintf("unknown argument %q", args[0])
	if len(args) > 1 {
		errMsg = fmt.Sprintf("unknown arguments %q", args)
	}

	hasValueFlag := false
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Value.Type() != "bool" {
			hasValueFlag = true
		}
	})

	if hasValueFlag {
		errMsg += "; please quote all values that have spaces"
	}

	return &FlagError{Err: errors.New(errMsg)}
}

func CheckAllNotEmpty(args ...string) error {
	for _, arg := range args {
		if gstring.IsEmpty(arg) {
			return EmptyParamError
		}
	}
	return nil
}
