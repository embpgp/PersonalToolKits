#!/bin/bash
qemu-system-x86_64 -kernel ./linux-4.14.191/arch/x86_64/boot/bzImage  -hda ./busybox-1.35.0/rootfs.img  -append "root=/dev/sda console=ttyS0" -nographic -s -S  -smp 1 -nographic -hdb ./sharedisk/ext4.img
