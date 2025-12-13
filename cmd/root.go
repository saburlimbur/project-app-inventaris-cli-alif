package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Aplikasi inventory CLI",
	Long:  "Sistem inventory kantor berbasis CLI (Golang + Postgres)",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
