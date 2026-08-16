components {
  id: "script"
  component: "/main/desk.script"
  properties {
    id: "coffee"
    value: "0.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "paper"
    value: "3.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "cats"
    value: "2.0"
    type: PROPERTY_TYPE_NUMBER
  }
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
