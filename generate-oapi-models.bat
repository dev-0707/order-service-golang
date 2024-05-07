@echo off
for %%I in (.) do set CurrDirName=%%~nxI
TITLE %CurrDirName%
set folder=%1
set model=%2
rem set excludedSchemas=%3
set spec=%3
rem oapi-codegen --config types.cfg.yaml -o internal\api\dto\%model%-dto.go docs\%spec%.yaml
rem oapi-codegen generate types -o internal\api\dto\%model%_dto.go -i docs\%spec%.yaml
rem https://github.com/deepmap/oapi-codegen
REM oapi-codegen -generate types --exclude-schemas=%excludedSchemas% -o internal\api\controllers\%folder%\%model%_dto.go -package %model% %spec%

md internal\api\controllers\%folder%
echo "oapi-codegen -generate types --exclude-schemas=%excludedSchemas% -o internal\api\controllers\%folder%\%model%_dto.go -package %model% %spec%"

rem oapi-codegen -generate types --exclude-schemas=%excludedSchemas% -o internal\api\controllers\%folder%\%model%_dto.go -package %model% %spec%
oapi-codegen -generate types -o internal\api\controllers\%folder%\%model%_dto.go -package %model% %spec%


rem ES set excludedSchemas=%
rem ES generate-oapi-models.bat check check %excludedSchemas% docs\tlp-proximity-channel-eb-service-api.yaml