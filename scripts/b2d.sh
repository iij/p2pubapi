#! /bin/sh
# Install Boot2Docker to P2PUB VM
#
# 1. attach two SystemStorage to VM (BootDevice, DataDevice)
# 2. execute $0
# 3. poweroff
# 4. detach two SystemStorage from VM
# 5. attach SystemStorage to VM (old DataDevice -> new BootDevice)
# 6. poweron

mntpt=$(mktemp -d)
dstdev=/dev/vdb
if [ -n "$shred" ] ; then
  shred --verbose -z -n 0 ${dstdev}1
fi

cat <<EOF | fdisk $dstdev
d
n
p
1

+1G
t
82
n
p
2


a
2
w
EOF

mkswap ${dstdev}1
swapon ${dstdev}1
mkfs.ext4 -E discard -i 8192 -L boot2docker-data ${dstdev}2
mount ${dstdev}2 ${mntpt}
grub2-install ${dstdev} --boot-directory=${mntpt}/boot
if [ ! -f boot2docker.iso ] ; then
  b2d_iso=$(curl -s https://api.github.com/repos/boot2docker/boot2docker/releases | grep browser_download_url | grep -v rc | head -n 1 | cut -f4 -d\")
  curl -LO ${b2d_iso}
fi
mount -o loop,ro boot2docker.iso /mnt
for i in vmlinuz64 initrd.img; do
  cp /mnt/boot/$i ${mntpt}/boot/
done
umount /mnt
cat <<EOF > ${mntpt}/boot/grub2/grub.cfg
set pager=1
set default="boot2docker"
set timeout=5
menuentry 'boot2docker' --class gnu-linux --class gnu --class os {
  linux16 /boot/vmlinuz64 loglevel=3 user=docker console=ttyS0 console=tty0 noembed nomodeset norestore waitusb=1:LABEL=boot2docker-data base
  initrd16 /boot/initrd.img
}
EOF
mkdir -p ${mntpt}/var/lib/boot2docker
tmpdir=$(mktemp -d)
cd $tmpdir
touch 'boot2docker, please format-me'
mkdir .ssh
cp ~/.ssh/authorized_keys .ssh/
cp .ssh/authorized_keys .ssh/authorized_keys2
chmod 700 .ssh
chmod 600 .ssh/authorized_keys*
tar cf ${mntpt}/var/lib/boot2docker/userdata.tar 'boot2docker, please format-me' .ssh
rm -rf $tmpdir

mkdir -p ${mntpt}/var/lib/boot2docker/ssh
cat <<EOF > ${mntpt}/var/lib/boot2docker/ssh/sshd_config
UseDNS no
Protocol 2
AllowUsers docker
ClientAliveInterval 300
ClientAliveCountMax 0
IgnoreRhosts yes
HostbasedAuthentication no
PermitRootLogin no
RSAAuthentication no
PubkeyAuthentication yes
PasswordAuthentication no
PermitEmptyPasswords no
ChallengeResponseAuthentication no
UsePrivilegeSeparation sandbox
EOF
umount ${mntpt}
rmdir ${mntpt}
