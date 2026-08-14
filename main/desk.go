components {
  id: "script"
  component: "/main/desk.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  "size {\n"
  "  x: 150.0\n"
  "  y: 100.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
}
embedded_components {
  id: "icon"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  "size {\n"
  "  x: 34.0\n"
  "  y: 34.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
  position {
    x: 0.0
    y: 84.0
    z: 0.1
  }
}
