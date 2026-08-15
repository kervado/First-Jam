components {
  id: "script"
  component: "/main/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"square\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 71.0\n"
  "  y: 33.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "snd_coffee"
  type: "sound"
  data: "sound: \"/assets/audio/coffee-pouring.wav\"\n"
  ""
}
embedded_components {
  id: "snd_paper"
  type: "sound"
  data: "sound: \"/assets/audio/paper-rustle.wav\"\n"
  ""
}
embedded_components {
  id: "snd_cat"
  type: "sound"
  data: "sound: \"/assets/audio/cat-meow.wav\"\n"
  ""
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/main/Player.spinescene\"\n"
  "default_animation: \"idle\"\n"
  "skin: \"base\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
}
