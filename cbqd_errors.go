package cbqd

import "errors"

var (
	INVALID_DATABASE_CREDENTIALS   = errors.New("Could not acquire the DATA STORE credentials.")
	INVALID_S3_STORAGE_CREDENTIALS = errors.New("Could not acquire the S3 STORAGE credentials")
	S3_BUCKET_ERROR_1              = errors.New("The bucket for data upload is not set.")
	S3_BUCKET_ERROR_2              = errors.New("The specified bucket does not exist in the cloud store.")
	S3_BUCKET_ERROR_3              = errors.New("The S3 host IP/FQDN is not set.")
	BACKUP_UPLOAD_ERROR            = errors.New("Backup could not be completed.Check the logs for more info.")
	BACKUP_FOLDER_ERROR            = errors.New("Could not open the temporary work directory.")
	BACKUP_ZIPROOT                 = errors.New("Could not create the output file for compression.")
	OS_ENVIRONMENT_UNSET           = errors.New("No credentials have been set for the process.Check the ENV variables.")
	VAULT_CREDENTIAL_ERROR         = errors.New("Vault faild to provide the access credentials.")
)
