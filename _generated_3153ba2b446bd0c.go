embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"default\"\nmask: \"dot\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_BOX\n    position {\n    }\n    rotation {\n    }\n    index: 0\n    count: 3\n    id: \"button_change\"\n  }\n  data: 400.0\n  data: 400.0\n  data: 10.0\n}\n"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"button\"\nmaterial: \"/builtins/materials/sprite.material\"\nsize {\n  x: 800.0\n  y: 800.0\n}\ntextures {\n  sampler: \"texture_sampler\"\n  texture: \"/main/alfa.atlas\"\n}\n"
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"change_betta\"\nmaterial: \"/builtins/materials/sprite.material\"\ntextures {\n  sampler: \"texture_sampler\"\n  texture: \"/main/betta_bb1.atlas\"\n}\n"
  position {
    z: 1.0
  }
}
