components {
  id: "script"
  component: "/main/request_icon.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"icon_cat 1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/icons/Icons.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
embedded_components {
  id: "alarm"
  type: "sprite"
  data: "default_animation: \"exclamation\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/icons/Icons.atlas\"\n"
  "}\n"
  ""
  position {
    x: 55.0
    y: 15.0
    z: 1.0
  }
}
