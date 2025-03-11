include {
  path = find_in_parent_folders("root.hcl")
}

locals {
  inputs_config = try(read_terragrunt_config("inputs.hcl"), null)
}

inputs = local.inputs_config.inputs