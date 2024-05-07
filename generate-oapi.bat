@echo off
for %%I in (.) do set CurrDirName=%%~nxI
TITLE %CurrDirName%
set folder=internal\api\controllers\%1
set package=%1
set controller=%2
set spec=%3

if not exist %folder% md %folder%

oapi-codegen -generate gin,spec -include-tags=%controller% -package %package% -o %folder%\%controller%_controller.go %spec%


rem Es: generate-oapi.bat test-api check docs\tlp-proximity-channel-eb-service-api.yaml
rem openapi-generator-cli generate -g dynamic-html -g go-gin-server -i docs\petstore-expanded.yaml -o server/ --additional-properties=outputAsLibrary=true,sourceFolder=openapipetstore-expanded.yaml > petstore.gen.go