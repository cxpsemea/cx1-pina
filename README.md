# cx1-pina
CheckmarxOne - single-purpose cli tool to create a project inside an application

Usage of cx1-pina.exe:
  -apikey string
        CheckmarxOne API Key (if not using client id/secret)
  -application string
        Application (name) to which project should be assigned
  -client string
        CheckmarxOne Client ID (if not using API Key)
  -cx1url string
        CheckmarxOne platform URL, eg: eu.ast.checkmarx.net
  -iamurl string
        CheckmarxOne IAM URL, eg: eu.iam.checkmarx.net
  -project string
        Project to be created
  -secret string
        CheckmarxOne Client Secret (if not using API Key)
  -tenant string
        CheckmarxOne tenant name


Example:
\cx1-pina>cx1-pina -project TestPina -application TestPina -apikey %CX1_EU_KEY% -cx1url https://eu.ast.checkmarx.net -iamurl https://eu.iam.checkmarx.net -tenant cx_tenant
[INFO][2023-10-19 07:52:19.623] Project: [68..94] TestPina
[INFO][2023-10-19 07:52:19.625] Application: [52..ea] TestPina
[INFO][2023-10-19 07:52:19.625] Project [68..94] TestPina is assigned to application [52..ea] TestPina and ready to use