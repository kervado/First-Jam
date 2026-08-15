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
    z: 0.02
  }
}
embedded_components {
  id: "snd_request"
  type: "sound"
  data: "sound: \"/assets/audio/request.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_coffee"
  type: "sound"
  data: "sound: \"/assets/audio/gulp.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_paper"
  type: "sound"
  data: "sound: \"/assets/audio/laser-printer.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_cat"
  type: "sound"
  data: "sound: \"/assets/audio/cat-purring.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_grunt1"
  type: "sound"
  data: "sound: \"/assets/audio/grunt1.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_grunt2"
  type: "sound"
  data: "sound: \"/assets/audio/grunt2.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
