#!/bin/bash
#qemu-system-x86_64 -kernel ./linux-4.14.191/arch/x86_64/boot/bzImage  -hda ./busybox-1.35.0/rootfs.img  -append "root=/dev/sda console=ttyS0" -nographic -s -smp 1 -nographic -hdb ./sharedisk/ext4.img  -nic tap,model=e1000
#sudo qemu-system-x86_64 -kernel ./linux-4.14.191/arch/x86_64/boot/bzImage  -hda ./busybox-1.35.0/rootfs.img  -append "root=/dev/sda console=ttyS0" -nographic -s -smp 1 -nographic -hdb ./sharedisk/ext4.img -net nic,macaddr=00:0c:29:ec:a9:3b -net bridge,id=net0,helper=/usr/lib/qemu/qemu-bridge-helper,br=br0
sudo qemu-system-x86_64 -kernel ./linux-4.14.191/arch/x86_64/boot/bzImage  -hda ./busybox-1.35.0/rootfs.img  -append "root=/dev/sda console=ttyS0" -nographic -s -smp 1 -nographic -hdb ./sharedisk/ext4.img -net nic,macaddr=4a:07:0f:b4:d3:98 -net tap,ifname=tap0,script=no,downscript=no