package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Everything about teams",
}

var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all teams",
	RunE: func(cmd *cobra.Command, args []string) error {
		teams, err := NewTeresa().GetTeams()
		if err != nil {
			return nil
		}
		fmt.Println("Teams:")
		for _, t := range teams {
			if t.IAmMember {
				fmt.Printf("  - %s (member)\n", *t.Name)
			} else {
				fmt.Printf("  - %s\n", *t.Name)
			}
			if t.Email != "" {
				fmt.Printf("    contact: %s\n", t.Email)
			}
			if t.URL != "" {
				fmt.Printf("    url: %s\n", t.URL)
			}
		}
		return nil
	},
}

var teamCreateCmd = &cobra.Command{
	Use:     "create <team-name>",
	Short:   "Create a team",
	Long:    "Create a team that can have many applications",
	Example: "$ teresa team create foo --email foo@foodomain.com --url http://site.foodomain.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return newUsageError("You should provide the name of the team in order to continue")
		}
		name := args[0]
		email, _ := cmd.Flags().GetString("email")
		site, _ := cmd.Flags().GetString("site")
		tc := NewTeresa()
		_, err := tc.CreateTeam(name, email, site)
		if err != nil {
			return err
		}
		fmt.Println("Team created with success")
		return nil
	},
}

//
//
// // delete team
// var deleteTeamCmd = &cobra.Command{
// 	Use:   "team",
// 	Short: "Delete a team",
// 	Long:  `Delete a team`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if teamIDFlag == 0 {
// 			Fatalf(cmd, "team ID is required")
// 		}
// 		if err := NewTeresa().DeleteTeam(teamIDFlag); err != nil {
// 			log.Fatalf("Failed to delete team: %s", err)
// 		}
// 		log.Infof("Team deleted.")
// 	},
// }
//
//
var teamAddUserCmd = &cobra.Command{
	Use:   "add-user",
	Short: "Add a member to a team",
	Long: `Add a member to a team.

You can add a new user as a member of a team with:

  $ teresa team add-user --email john.doe@foodomain.com --team foo

You need to create a user before use this command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		team, _ := cmd.Flags().GetString("team")
		user, _ := cmd.Flags().GetString("user")
		if team == "" || user == "" {
			return newUsageError("Team and User is required in order to continue")
		}
		tc := NewTeresa()
		_, err := tc.AddUserToTeam(team, user)
		if err != nil {
			if isUnprocessableEntity(err) {
				return newCmdError("Team or user doesn't exist, or the user is already member of the team")
			}
			return err
		}
		fmt.Printf("User %s is now member of the team %s\n", color.CyanString(user), color.CyanString(team))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(teamCmd)
	// Commands
	teamCmd.AddCommand(teamListCmd)
	teamCmd.AddCommand(teamCreateCmd)
	teamCmd.AddCommand(teamAddUserCmd)

	teamCreateCmd.Flags().String("email", "", "team email, if any")
	teamCreateCmd.Flags().String("url", "", "team site's URL, if any")

	teamAddUserCmd.Flags().String("user", "", "user email")
	teamAddUserCmd.Flags().String("team", "", "team name")

}
