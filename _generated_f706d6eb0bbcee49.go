components {
  id: "button_reklama_rel"
  component: "/main/button_reklama_rel.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"default\"\nmask: \"dot\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_BOX\n    position {\n    }\n    rotation {\n    }\n    index: 0\n    count: 3\n    id: \"button_reklama_rel\"\n  }\n  data: 400.0\n  data: 400.0\n  data: 10.0\n}\n"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"button\"\nmaterial: \"/builtins/materials/sprite.material\"\nsize {\n  x: 800.0\n  y: 800.0\n}\ntextures {\n  sampler: \"texture_sampler\"\n  texture: \"/main/alfa.atlas\"\n}\n"
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"reklama_betta\"\nmaterial: \"/builtins/materials/sprite.material\"\nsize {\n  x: 600.0\n  y: 600.0\n}\ntextures {\n  sampler: \"texture_sampler\"\n  texture: \"/main/betta_bb2.atlas\"\n}\n"
  position {
    z: 1.0
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "default_animation: \"Kin-9\"\nmaterial: \"/builtins/materials/sprite.material\"\ntextures {\n  sampler: \"texture_sampler\"\n  texture: \"/main/betta.atlas\"\n}\n"
  position {
    x: 15.0
    y: -100.0
    z: 5.0
  }
  scale {
    x: 0.6
    y: 0.6
  }
}
