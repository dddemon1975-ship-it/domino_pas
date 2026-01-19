components {
  id: "dot"
  component: "/main/dot.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"dot\"\nmask: \"default\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_CAPSULE\n    position {\n    }\n    rotation {\n      x: 0.70710677\n      w: 0.70710677\n    }\n    index: 0\n    count: 2\n    id: \"dot\"\n  }\n  data: 10.0\n  data: 100.0\n}\n"
}
