terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 4.5"
    }
  }
}

provider "docker" {}

resource "docker_image" "ubuntu" {
  name = "ubuntu:24.04"
}

resource "docker_container" "server" {
  name  = "luckyops-server"
  image = docker_image.ubuntu.image_id
}