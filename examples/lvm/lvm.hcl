param "device" {
    default = "/dev/sdb"
}

lvm.vg "vg-test" {
  name = "test"
  devices = [ "{{ param `device` }}" ]
}

lvm.lv "lv-test" {
  group = "test"
  name = "test"
  size = "1G"
  depends  = [ "lvm.vg.vg-test" ]
}

lvm.fs "mnt-me"  {
  device = "/dev/mapper/test-test"
  mount = "/mnt"
  fstype = "xfs"
  depends = [ "lvm.lv.lv-test" ]
}
