package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/manifoldco/promptui"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sqladmin/v1"
)

var project string

func setProject(env string) string {
	switch env {
	case "dev":
		project = "beecash-staging"
	case "staging":
		project = "beecash-staging"
	case "prod":
		project = "tokko-production"
	default:
		fmt.Print("Invalid Environment\n Please enter a valid environment: dev, staging, prod\n")
		os.Exit(1)
	}
	return project
}

func listInstances(project string) []string {
	var list []string
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, sqladmin.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	sqladminService, err := sqladmin.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := sqladminService.Instances.List(project)
	if err := req.Pages(ctx, func(page *sqladmin.InstancesListResponse) error {
		for _, databaseInstance := range page.Items {
			list = append(list, databaseInstance.ConnectionName)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return list
}

func getInstance(env string) string {
	project := setProject(env)
	instancelist := listInstances(project)
	prompt := promptui.Select{
		Label: "Select Project" + project,
		Items: instancelist,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}

func connectInstance(env string, port int) {
	sqlConnectionName := getInstance(env)
	fmt.Println("Connecting Instance")
	cmd := exec.Command("cloud_sql_proxy", "-enable_iam_login", "-instances="+sqlConnectionName+"=tcp:"+strconv.Itoa(port))
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cloudsql proxy process is running in background, process_id: %d\n", cmd.Process.Pid)
}
