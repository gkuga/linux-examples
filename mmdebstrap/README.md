docker build -t mmdebstrap .
docker run --rm -v $(pwd):/workspace mmdebstrap \
  stable debian-rootfs.tar http://deb.debian.org/debian
