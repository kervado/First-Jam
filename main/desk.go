components {
  id: "script"
  component: "/main/desk.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 232.0\n"
  "  y: 139.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "icon"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 34.0\n"
  "  y: 34.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
  position {
    y: 84.0
    z: 0.02
  }
}
embedded_components {
  id: "snd_request"
  type: "sound"
  data: "sound: \"/assets/audio/request.wav\"\n"
  ""
}
embedded_components {
  id: "snd_coffee"
  type: "sound"
  data: "sound: \"/assets/audio/gulp.wav\"\n"
  ""
}
embedded_components {
  id: "snd_paper"
  type: "sound"
  data: "sound: \"/assets/audio/laser-printer.wav\"\n"
  ""
}
embedded_components {
  id: "snd_cat"
  type: "sound"
  data: "sound: \"/assets/audio/cat-purring.wav\"\n"
  ""
}
embedded_components {
  id: "snd_grunt1"
  type: "sound"
  data: "sound: \"/assets/audio/grunt1.wav\"\n"
  ""
}
embedded_components {
  id: "snd_grunt2"
  type: "sound"
  data: "sound: \"/assets/audio/grunt2.wav\"\n"
  ""
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"Wooden desk with drawers\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 223.0\n"
  "  y: 136.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Office.atlas\"\n"
  "}\n"
  ""
  position {
    x: 2.0
    y: 31.0
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"Office chair cartoon drawing copy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 103.0\n"
  "  y: 114.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Office.atlas\"\n"
  "}\n"
  ""
  position {
    x: -4.0
    y: -39.0
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"box_small_2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 68.0\n"
  "  y: 56.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Office.atlas\"\n"
  "}\n"
  ""
  position {
    x: -95.0
    y: 54.0
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"box_small_1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 68.0\n"
  "  y: 56.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Office.atlas\"\n"
  "}\n"
  ""
  position {
    x: 66.0
    y: 68.0
  }
}
