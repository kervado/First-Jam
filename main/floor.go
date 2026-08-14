components {
  id: "script"
  component: "/main/solid.script"
  properties {
    id: "kind"
    value: "floor"
    type: PROPERTY_TYPE_HASH
  }
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
  "  x: 1920.0\n"
  "  y: 1080.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
}
