inputs = {
  vm_name           = "demo_linux_vm"
  vm_size          = "Standard_B1s"
  admin_username   = "azureuser"
  ssh_public_key   = "~/.ssh/night12pm_github_key.pub"
  my_home_ip       = "106.51.163.43"
  public_ip        = false
  nsg_id           = "demo_linux_nsg"
  location         = "South India"
  prefix           = "demo_linux"
  resource_group_name  = "demo_linux_rg"
  subnet_id            = "demo_linux_subnet"
}
