# IEEE-754 convert 

import struct

def float_to_hex(f):
    return hex(struct.unpack('<I', struct.pack('<f', f))[0])

print(float_to_hex(1.0)) # 0x3f800000
print(float_to_hex(1.5)) # 0x3fc00000
