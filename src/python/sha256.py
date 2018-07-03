# -*- coding: utf-8 -*-

import hashlib

#SHA256
myhash = hashlib.sha256()  #sha512将是更安全的哈希算法
myhash.update("1000001".encode("utf-8"))

print(myhash.hexdigest())  #哈希密码的二进制串
