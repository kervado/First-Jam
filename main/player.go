components {
  id: "script"
  component: "/main/player.script"
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
  "  x: 48.0\n"
  "  y: 48.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
}
embedded_components {
  id: "cup1"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 16.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
  position {
    x: -20.0
    y: 42.0
    z: 0.1
  }
}
embedded_components {
  id: "cup2"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 16.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
  position {
    x: 0.0
    y: 42.0
    z: 0.1
  }
}
embedded_components {
  id: "cup3"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 16.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
  position {
    x: 20.0
    y: 42.0
    z: 0.1
  }
}
