variable "namespace" {
  description = "Name of the namespace that will deploy the code"
  default     = "video"
}

variable "container_registry" {
  description = "Name of the container registry where images will be pulled"
}

variable "database_user" {
  description = "Username for the database"
}

variable "database_password" {
  description = "Password for the database"
}

variable "database_host" {
  description = "Host for the database"
  default     = "postgres.video.cluster.local"
}

variable "minio_host" {
  description = "Host for minio object storage service"
}

variable "minio_access_key" {
  description = "Access key for minio"
}

variable "minio_secret_key" {
  description = "Password for minio"
}

variable "salt" {
  description = "Salt used by user service"
}

variable "signing_key" {
  description = "Signing key used by user service"
}