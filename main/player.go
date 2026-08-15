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
  id: "icon"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  "size {\n"
  "  x: 24.0\n"
  "  y: 24.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  ""
  position {
    x: -16.0
    y: 48.0
    z: 0.1
  }
}
embedded_components {
  id: "count"
  type: "label"
  data: "size {\n"
  "  x: 64.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"0\"\n"
  "font: \"/main/font_small.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  "pivot: PIVOT_W\n"
  ""
  position {
    x: 2.0
    y: 48.0
    z: 0.1
  }
}
embedded_components {
  id: "snd_coffee"
  type: "sound"
  data: "sound: \"/assets/audio/coffee-pouring.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_paper"
  type: "sound"
  data: "sound: \"/assets/audio/paper-rustle.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
embedded_components {
  id: "snd_cat"
  type: "sound"
  data: "sound: \"/assets/audio/cat-meow.wav\"\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
}
