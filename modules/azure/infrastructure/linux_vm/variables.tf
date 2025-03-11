variable "vm_size" {
  type        = string
  default     = "Standard_B1s"
}

variable "public_ip" {
  type        = bool
  default     = false
}

variable "admin_username" {
  type        = string
  default     = "azureuser"
}

variable "ssh_public_key" {
  type        = string
}

variable "my_home_ip" {
  type        = string
}

variable "location" {
  type        = string
  default     = "South India"
}

variable "prefix" {
  type        = string
  default     = "demo_linux"
}

variable "resource_group_name" {
  type        = string
}

variable "subnet_id" {
  type        = string
}

variable "nsg_id" {
  type        = string
}



