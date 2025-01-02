packer {
  required_plugins {
    scaffolding = {
      version = "v0.2.4"
      source  = "github.com/bryborge/demo-pkr-plgin"
    }
  }
}

source "scaffolding-my-builder" "foo-example" {
  mock = local.foo
}

build {
  sources = [
    "source.scaffolding-my-builder.foo-example",
  ]

  source "source.scaffolding-my-builder.bar-example" {
    name = "bar"
  }

#   provisioner "scaffolding-my-provisioner" {
#     only = ["scaffolding-my-builder.foo-example"]
#     mock = "foo: ${local.foo}"
#   }
}
