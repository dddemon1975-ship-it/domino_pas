components {
  id: "pole"
  component: "/main/pole.script"
}
embedded_components {
  id: "bone_factory"
  type: "factory"
  data: "prototype: \"/main/bone.go\"\n"
}
embedded_components {
  id: "ramka_factory"
  type: "factory"
  data: "prototype: \"/main/ramka.go\"\n"
}
embedded_components {
  id: "bone_up_factory"
  type: "factory"
  data: "prototype: \"/main/bone_up.go\"\n"
}
embedded_components {
  id: "camera"
  type: "camera"
  data: "aspect_ratio: 1.0\nfov: 0.5\nnear_z: -10.0\nfar_z: 1000.0\northographic_projection: 1\northographic_zoom: 0.1\n"
}
