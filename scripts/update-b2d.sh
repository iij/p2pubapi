#! /bin/sh

b2d_iso=$(curl -s https://api.github.com/repos/boot2docker/boot2docker/releases | grep browser_download_url | grep -v rc | head -n 1 | cut -f4 -d\")
dest=/mnt/vda2

curl -LO $b2d_iso

mount -o loop,ro boot2docker.iso /mnt
cp /mnt/boot/vmlinuz64 $dest/boot/
cp /mnt/boot/initrd.img $dest/boot/
umount /mnt
