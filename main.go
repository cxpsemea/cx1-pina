package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/cxpsemea/Cx1ClientGo"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	myformatter := &easy.Formatter{}
	myformatter.TimestampFormat = "2006-01-02 15:04:05.000"
	myformatter.LogFormat = "[%lvl%][%time%] %msg%\n"
	logger.SetFormatter(myformatter)
	logger.SetOutput(os.Stdout)

	APIKey := flag.String("apikey", "", "CheckmarxOne API Key (if not using client id/secret)")
	ClientID := flag.String("client", "", "CheckmarxOne Client ID (if not using API Key)")
	ClientSecret := flag.String("secret", "", "CheckmarxOne Client Secret (if not using API Key)")
	Cx1URL := flag.String("cx1url", "", "CheckmarxOne platform URL, eg: eu.ast.checkmarx.net")
	IAMURL := flag.String("iamurl", "", "CheckmarxOne IAM URL, eg: eu.iam.checkmarx.net")
	Tenant := flag.String("tenant", "", "CheckmarxOne tenant name")
	ProjectName := flag.String("project", "", "Project to be created")
	ApplicationName := flag.String("application", "", "Application (name) to which project should be assigned")

	flag.Parse()

	if (*APIKey == "" && (*ClientID == "" || *ClientSecret == "")) || *Cx1URL == "" || *IAMURL == "" || *Tenant == "" || *ProjectName == "" || *ApplicationName == "" {
		logger.Info("The purpose of this tool is to create a project inside an application to address a temporary gap in the AST-CLI. The application will also be created if it does not exist.")
		logger.Fatal("Some parameters were not provided. For a list of parameters run: cx1-pina -h")
	}

	var cx1client *Cx1ClientGo.Cx1Client
	var err error
	httpClient := &http.Client{}

	if *APIKey != "" {
		cx1client, err = Cx1ClientGo.NewAPIKeyClient(httpClient, *Cx1URL, *IAMURL, *Tenant, *APIKey, logger)
	} else {
		cx1client, err = Cx1ClientGo.NewOAuthClient(httpClient, *Cx1URL, *IAMURL, *Tenant, *ClientID, *ClientSecret, logger)
	}

	if err != nil {
		logger.Fatalf("Failed to create CheckmarxOne client: %s", err)
	}

	project, app, err := cx1client.GetOrCreateProjectInApplicationByName(*ProjectName, *ApplicationName)
	if err != nil {
		logger.Fatalf("Could not create project %v in application %v: %s", *ProjectName, *ApplicationName, err)
	}

	logger.Infof("Project %v is assigned to application %v and ready to use", project.String(), app.String())
}
