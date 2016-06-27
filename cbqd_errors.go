package cbqd

import "errors"

var (
	INVALID_DATABASE_CREDENTIALS   = errors.New("could not acquire the DATA STORE credentials.")
	INVALID_S3_STORAGE_CREDENTIALS = errors.New("could not acquire the S3 STORAGE credentials")
	S3_BUCKET_ERROR_1              = errors.New("the bucket for data upload is not set.")
	S3_BUCKET_ERROR_2              = errors.New("the specified bucket does not exist in the cloud store.")
	S3_BUCKET_ERROR_3              = errors.New("the S3 host IP/FQDN is not set.")
	BACKUP_UPLOAD_ERROR            = errors.New("backup could not be completed.check the logs for more info.")
	BACKUP_FOLDER_ERROR            = errors.New("could not open the temporary work directory.")
	BACKUP_ZIPROOT                 = errors.New("could not create the output file for compression.")
	GPG_KEY_ERROR                  = errors.New("could not acquire the gpg public key for encrytion.")
	OS_ENVIRONMENT_UNSET           = errors.New("no credentials have been set for CBQD.check the environment variables.")
	VAULT_CREDENTIAL_ERROR         = errors.New("vault failed to provide the access credentials.")
	DB_TYPE_ERROR                  = errors.New("unsupported datbase type.")
	DB_DUMP_ERROR                  = errors.New("could not execute the database backup command.")
	DB_DUMP_ERROR_EXEC             = errors.New("could not execute the database backup command.executable file not found in $PATH")
)
