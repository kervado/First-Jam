components {
  id: "script"
  component: "/main/solid.script"
  properties {
    id: "kind"
    value: "floor"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "footprint_offset_y"
    value: "-150.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 1882.0\n"
  "  y: 779.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "normal"
  type: "sprite"
  data: "default_animation: \"paul\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 1920.0\n"
  "  y: 1080.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Back.atlas\"\n"
  "}\n"
  ""
  position {
    y: -2.0
    z: 1.0E-6
  }
}
embedded_components {
  id: "psychotic"
  type: "sprite"
  data: "default_animation: \"paul_p\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 1920.0\n"
  "  y: 1080.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Back.atlas\"\n"
  "}\n"
  ""
  position {
    y: -2.0
  }
}
