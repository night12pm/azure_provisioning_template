output "vm_private_ip" {
  value = azurerm_linux_virtual_machine.vm.private_ip_address
}

output "vm_public_ip" {
  value = azurerm_public_ip.vm_public_ip.ip_address
}


output "vm_disk_size" {
  value = azurerm_linux_virtual_machine.vm.os_disk.0.disk_size_gb
}