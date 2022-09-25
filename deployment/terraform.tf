terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.0"
    }
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_namespace" "video" {
  metadata {
    name = var.namespace
  }
}

resource "kubernetes_secret" "rabbitmq_dsn" {
  metadata {
    name      = "rabbitmq"
    namespace = var.namespace
  }
  data = {
    dsn = "amqp://rabbitmq.common.svc.cluster.local"
  }
}

resource "kubernetes_secret" "minio" {
  metadata {
    name      = "minio"
    namespace = var.namespace
  }
  data = {
    endpoint   = var.minio_host
    access_key = var.minio_access_key
    secret_key = var.minio_secret_key
  }
}

resource "kubernetes_secret" "database" {
  metadata {
    name      = "database"
    namespace = var.namespace
  }
  data = {
    user     = var.database_user
    password = var.database_password
    host     = var.database_host
  }
}

resource "kubernetes_secret" "user_svc" {
  metadata {
    name      = "user_svc"
    namespace = var.namespace
  }
  data = {
    salt        = var.database_user
    signing_key = var.database_password
  }
}

resource "kubernetes_deployment" "content" {
  metadata {
    name      = "content-service"
    namespace = var.namespace
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "content"
      }
    }
    template {
      metadata {
        labels = {
          app = "content"
        }
      }
      spec {
        container {
          name  = "content-service"
          image = "${var.container_registry}/content:03"
          env {
            name = "AMQP_DSN"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.rabbitmq_dsn.metadata.0.name
                key  = "dsn"
              }
            }
          }
          env {
            name = "DB_HOST"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "host"
              }
            }
          }
          env {
            name = "DB_USER"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "user"
              }
            }
          }
          env {
            name = "DB_PASSWORD"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "password"
              }
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment" "upload" {
  metadata {
    name      = "upload-service"
    namespace = var.namespace
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "upload"
      }
    }
    template {
      metadata {
        labels = {
          app = "upload"
        }
      }
      spec {
        container {
          name  = "upload-service"
          image = "${var.container_registry}/upload:03"
          env {
            name = "AMQP_DSN"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.rabbitmq_dsn.metadata.0.name
                key  = "dsn"
              }
            }
          }
          env {
            name = "DB_HOST"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "host"
              }
            }
          }
          env {
            name = "DB_USER"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "user"
              }
            }
          }
          env {
            name = "DB_PASSWORD"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "password"
              }
            }
          }
          env {
            name = "MINIO_ENDPOINT"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.minio.metadata.0.name
                key  = "endpoint"
              }
            }
          }
          env {
            name = "MINIO_ACCESSKEY"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.minio.metadata.0.name
                key  = "access_key"
              }
            }
          }
          env {
            name = "MINIO_SECRETKEY"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.minio.metadata.0.name
                key  = "secret_key"
              }
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment" "user" {
  metadata {
    name      = "user-service"
    namespace = var.namespace
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "user"
      }
    }
    template {
      metadata {
        labels = {
          app = "user"
        }
      }
      spec {
        container {
          name  = "user-service"
          image = "${var.container_registry}/user:01"
          env {
            name = "SALT"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.user_svc.metadata.0.name
                key  = "salt"
              }
            }
          }
          env {
            name = "SIGNING_KEY"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.user_svc.metadata.0.name
                key  = "SIGNING_KEY"
              }
            }
          }
          env {
            name = "DB_HOST"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "host"
              }
            }
          }
          env {
            name = "DB_USER"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "user"
              }
            }
          }
          env {
            name = "DB_PASSWORD"
            value_from {
              secret_key_ref {
                name = kubernetes_secret.database.metadata.0.name
                key  = "password"
              }
            }
          }
        }
      }
    }
  }
}

