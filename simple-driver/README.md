```
make
sudo insmod dummy.ko

echo hello | sudo tee /dev/dummy0
dmesg | tail
```
